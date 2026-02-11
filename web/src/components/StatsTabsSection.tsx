import { useEffect, useState } from 'react';
import {
  Title,
  Text,
  Paper,
  TextInput,
  Button,
  Stack,
  Group,
  Card,
  Grid,
  Center,
  Loader,
  Alert,
  Tabs,
  Table,
  Badge,
  Spoiler
} from '@mantine/core';
import {
  IconAlertCircle,
  IconDownload,
  IconTrophy,
  IconUsers
} from '@tabler/icons-react';

interface PersonalStats {
  legacy?: number;
  version?: number;
  credit?: number;
  credit_all?: number;
  medal_in?: number;
  medal_get?: number;
  ball_get?: number;
  ball_chain?: number;
  slot_start?: number;
  slot_startfev?: number;
  slot_hit?: number;
  slot_getfev?: number;
  sqr_get?: number;
  sqr_step?: number;
  jack_get?: number;
  jack_startmax?: number;
  jack_totalmax?: number;
  ult_get?: number;
  ult_combomax?: number;
  ult_totalmax?: number;
  rmshbi_get?: number;
  buy_shbi?: number;
  bstp_step?: number;
  bstp_rwd?: number;
  buy_total?: number;
  sp_use?: number;
  hide_record?: number;
  cpm_max?: number;
  jack_totalmax_v2?: number;
  ult_totalmax_v2?: number;
  palball_get?: number;
  pallot_lot_t0?: number;
  pallot_lot_t1?: number;
  pallot_lot_t2?: number;
  pallot_lot_t3?: number;
  jacksp_get_all?: number;
  jacksp_get_t0?: number;
  jacksp_get_t1?: number;
  jacksp_get_t2?: number;
  jacksp_get_t3?: number;
  jacksp_startmax?: number;
  jacksp_totalmax?: number;
  task_cnt?: number;
  firstboot?: string;
  lastsave?: string;
  playtime?: number;
  l_achieve?: string[];
}

interface RankingEntry {
  user_id: string;
  value: number;
  created_at: string;
}

interface GlobalStats {
  achievements_count: RankingEntry[];
  jacksp_startmax: RankingEntry[];
  golden_palball_get: RankingEntry[];
  cpm_max: RankingEntry[];
  max_chain_rainbow: RankingEntry[];
  jack_totalmax_v2: RankingEntry[];
  ult_combomax: RankingEntry[];
  ult_totalmax_v2: RankingEntry[];
  blackbox_total: RankingEntry[];
  sp_use: RankingEntry[];
  total_medals: number;
}

const rankingDefinitions: { key: keyof GlobalStats; label: string }[] = [
  { key: 'achievements_count', label: '実績解除数' },
  { key: 'jacksp_startmax', label: 'ジャックポット開始値' },
  { key: 'golden_palball_get', label: 'ゴールデンパレット獲得数' },
  { key: 'cpm_max', label: '最大CPM' },
  { key: 'max_chain_rainbow', label: '最大レインボーチェイン' },
  { key: 'jack_totalmax_v2', label: '最大ジャックポット(v2)' },
  { key: 'ult_combomax', label: '最大アルティメットコンボ' },
  { key: 'ult_totalmax_v2', label: 'アルティメット合計(v2)' },
  { key: 'blackbox_total', label: 'ブラックボックス累計' },
  { key: 'sp_use', label: 'スキルポイント使用数' }
];

const globalStatsEndpoint = 'https://push.trap.games/api/v4/statistics';
const objectValueSpoilerKeys = new Set(['dc_ball_chain', 'dc_ball_get', 'dc_medal_get', 'dc_palball_get', 'dc_palball_jp']);

function formatRankingValue(value: number, type: string) {
  switch (type) {
    case 'achievements_count':
    case 'blackbox_total':
      return `${value}個`;
    case 'cpm_max':
    case 'total_medals':
    default:
      return value.toLocaleString();
  }
}

function renderValue(key: string, val: unknown) {
  if (key === 'l_achieve' && Array.isArray(val)) {
    return (
      <Spoiler maxHeight={60} showLabel="もっと見る" hideLabel="閉じる">
        <Text size="sm" c="dimmed" style={{ wordBreak: 'break-all' }}>
          {val.join(', ')}
        </Text>
      </Spoiler>
    );
  }

  if (objectValueSpoilerKeys.has(key) && val && typeof val === 'object') {
    return (
      <Spoiler maxHeight={60} showLabel="展開" hideLabel="閉じる">
        <pre style={{ margin: 0, whiteSpace: 'pre-wrap', wordBreak: 'break-all', fontSize: '0.8rem' }}>
          {JSON.stringify(val, null, 2)}
        </pre>
      </Spoiler>
    );
  }

  return (
    <Text size="sm" c="dimmed" style={{ wordBreak: 'break-all' }}>
      {Array.isArray(val) ? val.join(', ') : `${val ?? 'N/A'}`}
    </Text>
  );
}

function StatsTabsSection({ primaryButtonColor }: { primaryButtonColor: string }) {
  const [dataUrl, setDataUrl] = useState('');
  const [activeTab, setActiveTab] = useState<string | null>('personal');
  const [isLoadingPersonal, setIsLoadingPersonal] = useState(false);
  const [isLoadingGlobal, setIsLoadingGlobal] = useState(false);
  const [personalStats, setPersonalStats] = useState<PersonalStats | null>(null);
  const [globalStats, setGlobalStats] = useState<GlobalStats | null>(null);
  const [error, setError] = useState('');
  const [rawPayload, setRawPayload] = useState('');

  useEffect(() => {
    if (activeTab !== 'global' || globalStats) {
      return;
    }

    const abortController = new AbortController();

    const fetchGlobalStats = async () => {
      setIsLoadingGlobal(true);
      try {
        const response = await fetch(globalStatsEndpoint, { signal: abortController.signal });
        if (!response.ok) {
          throw new Error('統計情報の取得に失敗しました');
        }
        const data = await response.json();
        setGlobalStats(data);
      } catch (err) {
        if (err instanceof DOMException && err.name === 'AbortError') {
          return;
        }
        console.error('統計情報の取得エラー:', err);
      } finally {
        setIsLoadingGlobal(false);
      }
    };

    fetchGlobalStats();
    return () => abortController.abort();
  }, [activeTab, globalStats]);

  const handleLoadPersonalData = async () => {
    if (!dataUrl.trim()) {
      setError('URLを入力してください');
      return;
    }

    setIsLoadingPersonal(true);
    setError('');
    setPersonalStats(null);

    try {
      const response = await fetch(dataUrl);

      if (!response.ok) {
        if (response.status === 404) {
          throw new Error('データが見つかりません。URLを確認してください。');
        } else if (response.status === 401) {
          throw new Error('署名認証に失敗しました。URLが正しいか確認してください。');
        } else {
          throw new Error(`データの取得に失敗しました（ステータスコード: ${response.status}）`);
        }
      }

      const payload = await response.json();
      if (payload && payload.data) {
        setRawPayload(payload.data);
        const normalized = payload.data.replace(/-/g, '+').replace(/_/g, '/').padEnd(Math.ceil(payload.data.length / 4) * 4, '=');
        const decoded = atob(normalized);
        setPersonalStats(JSON.parse(decoded));
      } else {
        setRawPayload('');
        setPersonalStats(payload);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'エラーが発生しました');
      setPersonalStats(null);
    } finally {
      setIsLoadingPersonal(false);
    }
  };

  const renderRankingTable = (data: RankingEntry[], title: string, type: string) => (
    <Card shadow="sm" padding="lg" radius="md" withBorder>
      <Group justify="space-between" mb="xs">
        <Title order={4}>{title}</Title>
        <Badge color="gray" variant="light">TOP {data.length}</Badge>
      </Group>
      <div style={{ maxHeight: '240px', overflowY: 'auto' }}>
        <Table highlightOnHover>
          <Table.Caption>{title} ランキング（TOP {data.length}）</Table.Caption>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>順位</Table.Th>
              <Table.Th>値</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {data.map((entry, index) => (
              <Table.Tr key={entry.user_id || `${index}`}>
                <Table.Td>
                  <Badge color={index === 0 ? 'yellow' : index === 1 ? 'gray' : index === 2 ? 'orange' : 'blue'}>
                    {index + 1}位
                  </Badge>
                </Table.Td>
                <Table.Td>{formatRankingValue(entry.value, type)}</Table.Td>
              </Table.Tr>
            ))}
          </Table.Tbody>
        </Table>
      </div>
    </Card>
  );

  return (
    <Tabs value={activeTab} onChange={setActiveTab} keepMounted={false} variant="outline">
      <Tabs.List aria-label="統計情報の切り替え">
        <Tabs.Tab value="personal" leftSection={<IconUsers size="1rem" />}>
          個人統計
        </Tabs.Tab>
        <Tabs.Tab value="global" leftSection={<IconTrophy size="1rem" />}>
          グローバル統計
        </Tabs.Tab>
      </Tabs.List>

      <Tabs.Panel value="personal" pt="md">
        <Paper p="xl" shadow="sm" radius="md" mih={360}>
          <Title order={2} mb="md" ta="center">
            個人統計情報
          </Title>
          <Text size="sm" c="dimmed" mb="md" ta="center">
            クラウドセーブで取得した URL を入力してください（Base64 応答を自動復号します）
          </Text>

          <Stack gap="md">
            <TextInput
              placeholder="https://push.trap.games/api/v4/users/xxxx/data?sig=xxxx"
              label="LoadSaveDataURL"
              description="クラウドセーブ URL を入力すると、レスポンスを自動で復号して表示します。"
              value={dataUrl}
              onChange={(e) => setDataUrl(e.currentTarget.value)}
            />

            <Group justify="center">
              <Button
                onClick={handleLoadPersonalData}
                loading={isLoadingPersonal}
                leftSection={<IconDownload size="1rem" />}
                color={primaryButtonColor}
                className="mmp-primary-button"
                radius="md"
              >
                データをロード
              </Button>
            </Group>

            {error && (
              <Alert icon={<IconAlertCircle size="1rem" />} title="エラー" color="red">
                {error}
              </Alert>
            )}

            {isLoadingPersonal && (
              <Center role="status" aria-live="polite">
                <Loader />
              </Center>
            )}

            {rawPayload && (
              <Card shadow="sm" padding="md" radius="md" withBorder>
                <Text size="sm" fw={600} mb="xs">受信データ (Base64)</Text>
                <Spoiler maxHeight={60} showLabel="展開" hideLabel="閉じる">
                  <Text size="xs" c="dimmed" style={{ wordBreak: 'break-all' }}>{rawPayload}</Text>
                </Spoiler>
              </Card>
            )}

            {personalStats && (
              <Card shadow="sm" padding="lg" radius="md" withBorder>
                <Title order={4}>個人統計データ</Title>
                <Grid mt="md" gutter="sm">
                  {Object.entries(personalStats).map(([key, val]) => (
                    <Grid.Col span={{ base: 12, sm: 6 }} key={key}>
                      <Text size="sm" fw={600}>{key}</Text>
                      {renderValue(key, val)}
                    </Grid.Col>
                  ))}
                </Grid>
              </Card>
            )}
          </Stack>
        </Paper>
      </Tabs.Panel>

      <Tabs.Panel value="global" pt="md">
        <Paper p="xl" shadow="sm" radius="md" mih={360}>
          {globalStats && (
            <Card shadow="sm" padding="lg" radius="md" withBorder mb="md">
              <Group justify="space-between">
                <Text fw={700}>世界の総メダル数</Text>
                <Text fw={800} fz="xl">
                  {globalStats.total_medals?.toLocaleString() ?? 'N/A'} 枚
                </Text>
              </Group>
            </Card>
          )}
          <Title order={2} mb="md">グローバル統計</Title>

          {isLoadingGlobal && (
            <Center role="status" aria-live="polite">
              <Loader />
            </Center>
          )}

          {globalStats && (
            <Grid gutter="md">
              {rankingDefinitions.map((def) => {
                const entries = globalStats[def.key] as RankingEntry[] | undefined;
                if (!entries) return null;
                return (
                  <Grid.Col span={{ base: 12, md: 6 }} key={def.key}>
                    {renderRankingTable(entries, def.label, def.key)}
                  </Grid.Col>
                );
              })}
            </Grid>
          )}
        </Paper>
      </Tabs.Panel>
    </Tabs>
  );
}

export default StatsTabsSection;

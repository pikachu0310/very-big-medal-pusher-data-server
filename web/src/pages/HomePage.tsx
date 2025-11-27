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
  ActionIcon,
  Anchor,
  Divider,
  Tooltip,
  ThemeIcon,
  Highlight
} from '@mantine/core';
import {
  IconAlertCircle,
  IconDownload,
  IconTrophy,
  IconUsers,
  IconBrandDiscord,
  IconWorld,
  IconBook2,
  IconBrandGithub,
  IconExternalLink,
  IconPlugConnected,
  IconRocket,
  IconCopy
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

function HomePage() {
  const [dataUrl, setDataUrl] = useState('');
  const [isLoadingPersonal, setIsLoadingPersonal] = useState(false);
  const [isLoadingGlobal, setIsLoadingGlobal] = useState(false);
  const [personalStats, setPersonalStats] = useState<PersonalStats | null>(null);
  const [globalStats, setGlobalStats] = useState<GlobalStats | null>(null);
  const [error, setError] = useState('');
  const [pingState, setPingState] = useState<'idle' | 'ok' | 'ng' | 'loading'>('idle');
  const [copyMessage, setCopyMessage] = useState<string>('');
  const twitterHashUrl = 'https://x.com/search?q=%23%E3%81%A7%E3%81%8B%E3%83%97%20OR%20%23VR%E3%81%A7%E3%81%8B%E3%83%97&f=live';

  useEffect(() => {
    const fetchGlobalStats = async () => {
      setIsLoadingGlobal(true);
      try {
        const response = await fetch('https://push.trap.games/api/v4/statistics');
        if (!response.ok) {
          throw new Error('統計情報の取得に失敗しました');
        }
        const data = await response.json();
        setGlobalStats(data);
      } catch (err) {
        console.error('統計情報の取得エラー:', err);
      } finally {
        setIsLoadingGlobal(false);
      }
    };

    fetchGlobalStats();
  }, []);

  useEffect(() => {
    const existing = document.getElementById('twitter-wjs');
    if (existing) return;
    const script = document.createElement('script');
    script.id = 'twitter-wjs';
    script.src = 'https://platform.twitter.com/widgets.js';
    script.async = true;
    document.body.appendChild(script);
  }, []);

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

      const data = await response.json();

      setPersonalStats(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'エラーが発生しました');
      setPersonalStats(null);
    } finally {
      setIsLoadingPersonal(false);
    }
  };

  const handlePing = async () => {
    setPingState('loading');
    try {
      const res = await fetch('https://push.trap.games/api/ping');
      setPingState(res.ok ? 'ok' : 'ng');
    } catch (err) {
      console.error(err);
      setPingState('ng');
    }
  };

  const handleCopy = async (text: string) => {
    try {
      await navigator.clipboard.writeText(text);
      setCopyMessage('コピーしました！');
      setTimeout(() => setCopyMessage(''), 2000);
    } catch {
      setCopyMessage('コピーに失敗しました');
      setTimeout(() => setCopyMessage(''), 2000);
    }
  };

  const formatRankingValue = (value: number, type: string) => {
    switch (type) {
      case 'achievements_count':
        return `${value}個`;
      case 'blackbox_total':
        return `${value}個`;
      case 'cpm_max':
        return value.toLocaleString();
      case 'total_medals':
        return value.toLocaleString();
      default:
        return value.toLocaleString();
    }
  };

  const renderRankingTable = (data: RankingEntry[], title: string, type: string) => (
    <Card shadow="sm" padding="lg" radius="md" withBorder>
      <Title order={4} mb="md">{title}</Title>
      <Text size="sm" c="dimmed" mb="sm">総エントリー数: {data.length}件</Text>
      <div style={{ maxHeight: '420px', overflowY: 'auto' }}>
        <Table>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>順位</Table.Th>
              <Table.Th>値</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {data.map((entry, index) => (
              <Table.Tr key={`${index}`}>
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
    <Stack gap="xl" pt={0}>
      {/* ヒーロー */}
      <Card padding="xl" radius="md" shadow="sm" style={{ background: 'linear-gradient(135deg, #e7f5ff 0%, #d0ebff 100%)' }}>
        <Stack gap="md">
          <Highlight
            highlight={['Massive Medal Pusher']}
            fw={700}
            fz={14}
            c="#1e1e1e"
            highlightStyles={(theme) => ({
              backgroundColor: theme.colors.gray[1],
              color: theme.colors.blue[7],
            })}
          >
            Massive Medal Pusher / リンク集
          </Highlight>
          <Grid gutter="md">
            <Grid.Col span={{ base: 12, md: 6 }}>
              <Button
                size="xl"
                radius="md"
                fullWidth
                leftSection={<IconBrandDiscord size={22} />}
                component="a"
                href="https://discord.com/invite/CgnYyXecKm"
                target="_blank"
                rel="noreferrer"
                variant="filled"
                color="indigo"
                fw={700}
                c="white"
              >
                公式Discord でかプ同好会
              </Button>
              <Button
                mt="sm"
                size="xl"
                radius="md"
                fullWidth
                leftSection={<IconBook2 size={22} />}
                component="a"
                href="https://wikiwiki.jp/vr_bigpusher/"
                target="_blank"
                rel="noreferrer"
                variant="filled"
                color="cyan"
                fw={700}
                c="white"
              >
                公式Wiki クソでっけぇプッシャーゲーム
              </Button>
            </Grid.Col>
            <Grid.Col span={{ base: 12, md: 6 }}>
              <Button
                size="lg"
                radius="md"
                fullWidth
                leftSection={<IconWorld size={20} />}
                component="a"
                href="https://vrchat.com/home/group/grp_5900a25d-0bb9-48d4-bab1-f3bd5c9a5e73"
                target="_blank"
                rel="noreferrer"
                variant="outline"
                color="dark"
                fw={700}
              >
                VRChatグループ クソでっけぇプッシャーゲーム同好会
              </Button>
              <Button
                mt="sm"
                size="lg"
                radius="md"
                fullWidth
                leftSection={<IconWorld size={20} />}
                component="a"
                href="https://vrchat.com/home/group/grp_f38ec6a3-0de5-499e-a85f-1038013bdd04"
                target="_blank"
                rel="noreferrer"
                variant="outline"
                color="dark"
                fw={700}
              >
                でかプ交流会 ～ MMP Meeting
              </Button>
              <Button
                mt="sm"
                size="lg"
                radius="md"
                fullWidth
                leftSection={<IconRocket size={20} />}
                component="a"
                href="https://github.com/pikachu0310/very-big-medal-pusher-data-server"
                target="_blank"
                rel="noreferrer"
                variant="gradient"
                gradient={{ from: 'teal', to: 'blue' }}
                fw={700}
                c="white"
              >
                v4 クラウドセーブ稼働中（GitHub）
              </Button>
            </Grid.Col>
          </Grid>
        </Stack>
      </Card>

      <Tabs defaultValue="personal" variant="outline">
        <Tabs.List>
          <Tabs.Tab value="personal" leftSection={<IconUsers size="1rem" />}>
            個人統計
          </Tabs.Tab>
          <Tabs.Tab value="global" leftSection={<IconTrophy size="1rem" />}>
            グローバル統計
          </Tabs.Tab>
        </Tabs.List>

        <Tabs.Panel value="personal" pt="md">
          <Paper p="xl" shadow="sm" radius="md">
            <Title order={2} mb="md" ta="center">
              個人統計情報
            </Title>
            <Text size="sm" c="dimmed" mb="md" ta="center">
              クラウドセーブで取得した URL を入力してください
            </Text>

            <Stack gap="md">
              <TextInput
                placeholder="https://push.trap.games/api/v4/users/xxxx/data?sig=xxxx"
                label="LoadSaveDataURL"
                value={dataUrl}
                onChange={(e) => setDataUrl(e.currentTarget.value)}
              />

              <Group justify="center">
                <Button
                  onClick={handleLoadPersonalData}
                  loading={isLoadingPersonal}
                  leftSection={<IconDownload size="1rem" />}
                  color="blue"
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
                <Center>
                  <Loader />
                </Center>
              )}

              {personalStats && (
                <Card shadow="sm" padding="lg" radius="md" withBorder>
                  <Title order={4}>個人統計データ</Title>
                  <Grid mt="md">
                    <Grid.Col span={6}>バージョン: {personalStats.version ?? 'N/A'}</Grid.Col>
                    <Grid.Col span={6}>クレジット総額: {personalStats.credit_all?.toLocaleString() ?? 'N/A'}</Grid.Col>
                    <Grid.Col span={6}>プレイ時間: {personalStats.playtime?.toLocaleString() ?? 'N/A'}</Grid.Col>
                    <Grid.Col span={6}>実績数: {personalStats.l_achieve?.length ?? 'N/A'}</Grid.Col>
                    <Grid.Col span={6}>ジャックポット獲得: {personalStats.jack_get?.toLocaleString() ?? 'N/A'}</Grid.Col>
                    <Grid.Col span={6}>すごろく進行: {personalStats.sqr_step?.toLocaleString() ?? 'N/A'}</Grid.Col>
                  </Grid>
                </Card>
              )}
            </Stack>
          </Paper>
        </Tabs.Panel>

        <Tabs.Panel value="global" pt="md">
          <Paper p="xl" shadow="sm" radius="md">
            <Group justify="space-between" mb="md" align="center">
              <Title order={2}>グローバル統計</Title>
              <Group gap="sm">
                <Button
                  size="xs"
                  variant="light"
                  leftSection={<IconExternalLink size={14} />}
                  component="a"
                  href="https://push.trap.games/api/v4/statistics"
                  target="_blank"
                  rel="noreferrer"
                  color="blue"
                >
                  APIを見る
                </Button>
              </Group>
            </Group>

            {isLoadingGlobal && (
              <Center>
                <Loader />
              </Center>
            )}

            {globalStats && (
              <Grid gutter="md">
                {rankingDefinitions.map((def) => {
                  const entries = (globalStats as any)[def.key] as RankingEntry[] | undefined;
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

      {/* Twitter / X セクション */}
      <Card withBorder radius="md" padding="lg" shadow="sm">
        <Group justify="space-between" mb="xs">
          <Title order={3}>#でかプ / #VRでかプ リアルタイム</Title>
          <Button
            size="xs"
            variant="light"
            component="a"
            href={twitterHashUrl}
            target="_blank"
            rel="noreferrer"
            leftSection={<IconExternalLink size={14} />}
            color="blue"
          >
            Xで開く
          </Button>
        </Group>
        <Text size="sm" c="dimmed" mb="md">
          公式ハッシュタグの最新投稿をチェックできます。（読み込めない場合は上のボタンから直接開いてください）
        </Text>
        <div style={{ border: '1px solid #e9ecef', borderRadius: 12, overflow: 'hidden', minHeight: 420, padding: '0.5rem' }}>
          <a
            className="twitter-timeline"
            data-theme="light"
            data-height="520"
            data-chrome="nofooter noborders transparent"
            href={twitterHashUrl}
          >
            Tweets about #でかプ
          </a>
        </div>
      </Card>

      {/* 開発者ツール */}
      <Grid gutter="md">
        <Grid.Col span={{ base: 12, md: 6 }}>
          <Card withBorder radius="md" padding="lg" shadow="sm">
            <Group justify="space-between" mb="sm">
              <Title order={3}>開発者向けリンク</Title>
              <Tooltip label="コピーできるよ" position="left">
                <Badge color={copyMessage ? 'teal' : 'gray'}>{copyMessage || 'copy ready'}</Badge>
              </Tooltip>
            </Group>
            <Stack gap="sm">
              <Group justify="space-between">
                <Text>本番 API</Text>
                <Group gap="xs">
                  <Anchor href="https://push.trap.games/api" target="_blank" rel="noreferrer" c="blue">
                    開く
                  </Anchor>
                  <ActionIcon variant="light" onClick={() => handleCopy('https://push.trap.games/api')}>
                    <IconCopy size={16} />
                  </ActionIcon>
                </Group>
              </Group>
              <Group justify="space-between">
                <Text>テスト API</Text>
                <Group gap="xs">
                  <Anchor href="https://push-test.trap.games/api" target="_blank" rel="noreferrer" c="blue">
                    開く
                  </Anchor>
                  <ActionIcon variant="light" onClick={() => handleCopy('https://push-test.trap.games/api')}>
                    <IconCopy size={16} />
                  </ActionIcon>
                </Group>
              </Group>
              <Group justify="space-between">
                <Text>ローカル API</Text>
                <Group gap="xs">
                  <Anchor href="http://localhost:8080/api" target="_blank" rel="noreferrer" c="blue">
                    開く
                  </Anchor>
                  <ActionIcon variant="light" onClick={() => handleCopy('http://localhost:8080/api')}>
                    <IconCopy size={16} />
                  </ActionIcon>
                </Group>
              </Group>
              <Divider />
              <Group gap="sm">
                <Button
                  variant="gradient"
                  gradient={{ from: 'indigo', to: 'cyan' }}
                  component="a"
                  href="/swagger/index.html"
                  target="_blank"
                  radius="md"
                  leftSection={<IconExternalLink size={16} />}
                  c="white"
                >
                  Swagger UI
                </Button>
                <Button
                  variant="subtle"
                  component="a"
                  href="/api/openapi.yaml"
                  target="_blank"
                  radius="md"
                  leftSection={<IconDownload size={16} />}
                >
                  OpenAPI をダウンロード
                </Button>
              </Group>
              <Button
                variant="light"
                component="a"
                href="https://push.trap.show/?server=mariadb.ns-system.svc.cluster.local&username=nsapp_c27d6f571f88ffff360fe2&db=nsapp_c27d6f571f88ffff360fe2"
                target="_blank"
                rel="noreferrer"
                radius="md"
                color="blue"
              >
                DBにアクセス
              </Button>
            </Stack>
          </Card>
        </Grid.Col>
        <Grid.Col span={{ base: 12, md: 6 }}>
          <Card withBorder radius="md" padding="lg" shadow="sm">
            <Group justify="space-between" align="center">
              <Title order={3}>サーバーヘルス</Title>
              <Button
                size="sm"
                variant="light"
                leftSection={<IconPlugConnected size={14} />}
                loading={pingState === 'loading'}
                onClick={handlePing}
              >
                ping
              </Button>
            </Group>
            <Text size="sm" c="dimmed" mb="sm">
              https://push.trap.games/api/ping
            </Text>
            <Group gap="sm">
              <Badge color={pingState === 'ok' ? 'teal' : pingState === 'ng' ? 'red' : 'gray'} radius="sm">
                {pingState === 'idle' && '未実行'}
                {pingState === 'loading' && '確認中...'}
                {pingState === 'ok' && '稼働中'}
                {pingState === 'ng' && '疎通 NG'}
              </Badge>
            </Group>
            <Divider my="md" />
            <Group gap="xs">
              <ThemeIcon size={34} radius="lg" variant="light" color="blue">
                <IconBrandGithub size={18} />
              </ThemeIcon>
              <Stack gap={4}>
                <Text fw={600} fz="sm">GitHub</Text>
                <Anchor href="https://github.com/pikachu0310/very-big-medal-pusher-data-server" target="_blank" rel="noreferrer" c="blue">
                  very-big-medal-pusher-data-server
                </Anchor>
              </Stack>
            </Group>
          </Card>
        </Grid.Col>
      </Grid>
    </Stack>
  );
}

export default HomePage;

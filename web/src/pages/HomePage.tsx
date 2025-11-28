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
  Anchor,
  ThemeIcon,
  Highlight,
  Spoiler,
  Box
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
  IconLock
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
  const twitterHashUrl = 'https://x.com/search?q=%23%E3%81%A7%E3%81%8B%E3%83%97%20OR%20%23VR%E3%81%A7%E3%81%8B%E3%83%97&f=live';
  const [rawPayload, setRawPayload] = useState<string>('');
  const isObjectKey = new Set(['dc_ball_chain', 'dc_ball_get', 'dc_medal_get', 'dc_palball_get', 'dc_palball_jp']);

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

      // v4 は { data: base64, sig: ... } で返るので data を decode
      if (payload && payload.data) {
        setRawPayload(payload.data);
        const decoded = decodeBase64(payload.data);
        const parsed = JSON.parse(decoded);
        setPersonalStats(parsed);
      } else {
        setRawPayload('');
        // 互換: 直接 SaveData が返る場合
        setPersonalStats(payload);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'エラーが発生しました');
      setPersonalStats(null);
    } finally {
      setIsLoadingPersonal(false);
    }
  };

  const decodeBase64 = (b64: string) => {
    const normalized = b64.replace(/-/g, '+').replace(/_/g, '/').padEnd(Math.ceil(b64.length / 4) * 4, '=');
    return atob(normalized);
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
      <Group justify="space-between" mb="xs">
        <Title order={4}>{title}</Title>
        <Badge color="gray" variant="light">TOP {data.length}</Badge>
      </Group>
      <div style={{ maxHeight: '240px', overflowY: 'auto' }}>
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

  const renderValue = (key: string, val: any) => {
    if (key === 'l_achieve' && Array.isArray(val)) {
      return (
        <Spoiler maxHeight={60} showLabel="もっと見る" hideLabel="閉じる">
          <Text size="sm" c="dimmed" style={{ wordBreak: 'break-all' }}>
            {val.join(', ')}
          </Text>
        </Spoiler>
      );
    }
    if (isObjectKey.has(key) && val && typeof val === 'object') {
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
  };

  const HeroButton = ({
    children,
    href,
    icon,
    size = 'lg',
    variant = 'filled',
    gradient,
    color,
    heightMultiplier = 1,
  }: {
    children: React.ReactNode;
    href: string;
    icon: React.ReactNode;
    size?: 'lg' | 'xl';
    variant?: 'filled' | 'outline' | 'light' | 'gradient';
    gradient?: { from: string; to: string };
    color?: string;
    heightMultiplier?: number;
  }) => (
    <Button
      component="a"
      href={href}
      target="_blank"
      rel="noreferrer"
      leftSection={icon}
      size={size}
      radius="md"
      fullWidth
      variant={variant}
      gradient={gradient}
      color={color}
      fw={700}
      c={variant === 'outline' ? color || 'dark' : 'white'}
      style={{
        minHeight: `calc(${size === 'xl' ? 52 : 44}px * ${heightMultiplier})`,
        whiteSpace: 'normal',
        lineHeight: 1.3,
        paddingInline: '1.4rem',
        textAlign: 'center',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        wordBreak: 'break-word'
      }}
    >
      {children}
    </Button>
  );

  return (
    <Stack gap="xl" pt={0}>
      {/* ロゴ */}
      <Center mt={0} pt={0} mb={0}>
        <img
          src="/MPG_logo.png"
          alt="MPG Logo"
          style={{
            maxWidth: '700px',
            width: '100%',
            height: 'auto',
            marginBottom: '0.5rem',
            marginTop: '0'
          }}
        />
      </Center>

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
              <HeroButton
                href="https://discord.com/invite/CgnYyXecKm"
                icon={<IconBrandDiscord size={33} />}
                size="xl"
                variant="gradient"
                gradient={{ from: 'grape', to: 'indigo' }}
                heightMultiplier={2}
              >
                公式Discord でかプ同好会
              </HeroButton>
              <Box mt="sm">
                <HeroButton
                  href="https://wikiwiki.jp/vr_bigpusher/"
                  icon={<IconBook2 size={22} />}
                  size="xl"
                  variant="filled"
                  color="blue"
                >
                  公式Wiki
                </HeroButton>
              </Box>
            </Grid.Col>
            <Grid.Col span={{ base: 12, md: 6 }}>
              <HeroButton
                href="https://vrchat.com/home/group/grp_5900a25d-0bb9-48d4-bab1-f3bd5c9a5e73"
                icon={<IconWorld size={20} />}
                size="lg"
                variant="filled"
                color="indigo"
              >
                公式グループ(VRChat)
              </HeroButton>
              <Box mt="sm">
                <HeroButton
                  href="https://vrchat.com/home/launch?worldId=wrld_1af53798-92a3-4c3f-99ae-a7c42ec6084d"
                  icon={<IconWorld size={20} />}
                  size="lg"
                  variant="filled"
                  color="indigo"
                >
                  VRChatワールドリンク
                </HeroButton>
              </Box>
              <Box mt="sm">
                <HeroButton
                  href={twitterHashUrl}
                  icon={<IconExternalLink size={20} />}
                  size="lg"
                  variant="light"
                  color="blue"
                  heightMultiplier={1}
                >
                  #でかプ / #VRでかプ (X投稿)
                </HeroButton>
              </Box>
            </Grid.Col>
          </Grid>
        </Stack>
      </Card>

      {/* 開発者向けリンク集 */}
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
            Massive Medal Pusher / 開発者向けリンク集
          </Highlight>
          <Grid gutter="sm">
            <Grid.Col span={{ base: 12, sm: 4 }}>
              <HeroButton
                href="/swagger/index.html"
                icon={<IconExternalLink size={18} />}
                size="lg"
                variant="gradient"
                gradient={{ from: 'orange', to: 'red' }}
                heightMultiplier={1.1}
              >
                SwaggerUI (API一覧)
              </HeroButton>
            </Grid.Col>
            <Grid.Col span={{ base: 12, sm: 4 }}>
              <HeroButton
                href="https://push.trap.show/?server=mariadb.ns-system.svc.cluster.local&username=nsapp_c27d6f571f88ffff360fe2&db=nsapp_c27d6f571f88ffff360fe2"
                icon={<IconLock size={18} />}
                size="lg"
                variant="outline"
                color="black"
                heightMultiplier={1.1}
              >
                データベース
              </HeroButton>
            </Grid.Col>
            <Grid.Col span={{ base: 12, sm: 4 }}>
              <HeroButton
                href="https://github.com/pikachu0310/very-big-medal-pusher-data-server"
                icon={<IconBrandGithub size={18} />}
                size="lg"
                variant="light"
                color="gray"
                heightMultiplier={1.1}
              >
                Data Server GitHub
              </HeroButton>
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
              クラウドセーブで取得した URL を入力してください（Base64 応答を自動復号します）
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
          <Paper p="xl" shadow="sm" radius="md">
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

      {/* 開発者ツール */}
      <Grid gutter="md">
        <Grid.Col span={{ base: 12, md: 6 }}>
          <Card withBorder radius="md" padding="lg" shadow="sm">
            <Group justify="space-between" align="center">
              <Title order={3}>サーバーヘルス</Title>
              <Button
                size="sm"
                variant="light"
                color="gray"
                c="dark"
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
          </Card>
        </Grid.Col>
        <Grid.Col span={{ base: 12, md: 6 }}>
          <Card withBorder radius="md" padding="lg" shadow="sm">
            <Group gap="xs">
              <ThemeIcon size={34} radius="lg" variant="light" color="blue">
                <IconBrandGithub size={18} />
              </ThemeIcon>
              <Stack gap={4}>
                <Text fw={600} fz="sm">GitHub</Text>
                <Anchor href="https://github.com/pikachu0310/very-big-medal-pusher-data-server" target="_blank" rel="noreferrer" c="blue">
                  very-big-medal-pusher-data-server
                </Anchor>
                <Anchor href="https://github.com/pikachu0310/VRCWorld-MassiveMedalPusher" target="_blank" rel="noreferrer" c="blue">
                  <Group gap={6}>
                    <IconLock size={14} />
                    <Text component="span" size="sm" c="blue">VRCWorld-MassiveMedalPusher</Text>
                  </Group>
                </Anchor>
                <Anchor href="https://github.com/pikariku/VRCWorld-VeryBigMedalPusher" target="_blank" rel="noreferrer" c="blue">
                  <Group gap={6}>
                    <IconLock size={14} />
                    <Text component="span" size="sm" c="blue">VRCWorld-VeryBigMedalPusher</Text>
                  </Group>
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

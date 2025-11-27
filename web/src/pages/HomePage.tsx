import { useState, useEffect } from 'react';
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
  Highlight,
  Accordion
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
  IconSparkles,
  IconMoodSmile,
  IconBulb,
  IconWand
} from '@tabler/icons-react';

// 個人統計情報の型定義（SaveDataV2の全フィールド）
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

// ランキングエントリの型定義
interface RankingEntry {
  user_id: string;
  value: number;
  created_at: string;
}

// グローバル統計情報の型定義
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

function HomePage() {
  const [dataUrl, setDataUrl] = useState('');
  const [isLoadingPersonal, setIsLoadingPersonal] = useState(false);
  const [isLoadingGlobal, setIsLoadingGlobal] = useState(false);
  const [personalStats, setPersonalStats] = useState<PersonalStats | null>(null);
  const [globalStats, setGlobalStats] = useState<GlobalStats | null>(null);
  const [error, setError] = useState('');
  const [pingState, setPingState] = useState<'idle' | 'ok' | 'ng' | 'loading'>('idle');
  const [copyMessage, setCopyMessage] = useState<string>('');
  const [randomMission, setRandomMission] = useState<string>('');
  const [randomTip, setRandomTip] = useState<string>('');
  const twitterSearchUrl =
    'https://x.com/search?q=%28%23%E3%81%A7%E3%81%8B%E3%83%97%20OR%20%23VR%E3%81%A7%E3%81%8B%E3%83%97%29&f=live';

  // グローバル統計情報を取得
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

  // Twitter ウィジェット読み込み
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
      // 入力されたURLをそのまま使用してAPIリクエスト
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
      
      // APIから取得したデータを全てセット
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

  const missions = [
    'ジャックポットを1回以上引こう！',
    'レインボーボールでコンボを狙え！',
    'パレットボールを10回拾う',
    'シャルベチャンスでフィーバー突入',
    'スキルポイントを1回リセットしてみよう'
  ];

  const tips = [
    'クラウドセーブはこまめに！ロードで復元が安心。',
    'でかプ交流会に参加するとギミック講習が聞けるかも？',
    'ランキングは v4/statistics から高速に取得されます。',
    'ゴールデンパレットボール(100番)を狙うと実績が埋まるぞ。',
    'ジャックポット中はチェインを切らさない立ち回りが大事！'
  ];

  const pickMission = () => setRandomMission(missions[Math.floor(Math.random() * missions.length)]);
  const pickTip = () => setRandomTip(tips[Math.floor(Math.random() * tips.length)]);

  const renderRankingTable = (data: RankingEntry[], title: string, type: string) => (
    <Card shadow="sm" padding="lg" radius="md" withBorder>
      <Title order={4} mb="md">{title}</Title>
      <Text size="sm" c="dimmed" mb="sm">総エントリー数: {data.length}件</Text>
      <div style={{ maxHeight: '400px', overflowY: 'auto' }}>
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
      {/* ロゴを中央上部に */}
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
      <Card padding="xl" radius="md" shadow="sm" style={{ background: 'linear-gradient(135deg, #f1f3f5 0%, #e7f5ff 100%)' }}>
        <Stack gap="md">
          <Group justify="space-between" align="flex-start">
            <Stack gap={8}>
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
                Massive Medal Pusher / ページのタイトル
              </Highlight>
              <Title order={1} style={{ letterSpacing: '-0.02em' }}>
                クソでっけぇプッシャーゲーム
              </Title>
              <Text size="lg" c="#1e1e1e">
                セーブ共有・ランキング・最新情報をまとめた公式ポータル。プレイヤー向けリンクと遊び方ヒントを集約しました。
              </Text>
              <Group gap="sm">
                <Button
                  size="md"
                  leftSection={<IconBrandDiscord size={18} />}
                  component="a"
                  href="https://discord.com/invite/CgnYyXecKm"
                  target="_blank"
                  rel="noreferrer"
                  radius="md"
                >
                  公式Discordへ参加
                </Button>
                <Button
                  size="md"
                  variant="outline"
                  leftSection={<IconBook2 size={18} />}
                  component="a"
                  href="https://wikiwiki.jp/vr_bigpusher/"
                  target="_blank"
                  rel="noreferrer"
                  radius="md"
                >
                  公式Wikiを開く
                </Button>
              </Group>
              <Group gap="sm">
                <Button
                  variant="light"
                  size="sm"
                  leftSection={<IconWorld size={16} />}
                  component="a"
                  href="https://vrchat.com/home/group/grp_5900a25d-0bb9-48d4-bab1-f3bd5c9a5e73"
                  target="_blank"
                  rel="noreferrer"
                  radius="md"
                >
                  VRChatグループ
                </Button>
                <Button
                  variant="light"
                  size="sm"
                  leftSection={<IconWorld size={16} />}
                  component="a"
                  href="https://vrchat.com/home/group/grp_f38ec6a3-0de5-499e-a85f-1038013bdd04"
                  target="_blank"
                  rel="noreferrer"
                  radius="md"
                >
                  でかプ交流会
                </Button>
              </Group>
            </Stack>
            <img
              src="/MPG_logo.png"
              alt="MPG Logo"
              style={{
                maxWidth: '260px',
                width: '35vw',
                height: 'auto',
                objectFit: 'contain',
                filter: 'drop-shadow(0 8px 18px rgba(0,0,0,0.15))'
              }}
            />
          </Group>
          <Divider />
          <Group gap="sm" wrap="wrap">
            <Badge variant="dot" color="teal" size="lg" radius="sm" leftSection={<IconRocket size={14} />}>
              v4 クラウドセーブ稼働中
            </Badge>
            <Badge variant="outline" color="blue" size="lg" radius="sm">
              API: https://push.trap.games/api
            </Badge>
            <Badge variant="outline" color="gray" size="lg" radius="sm">
              Swagger: /swagger/index.html
            </Badge>
          </Group>
        </Stack>
      </Card>

      {/* 楽しむセクション */}
      <Grid gutter="md">
        <Grid.Col span={{ base: 12, md: 8 }}>
          <Card withBorder radius="md" padding="lg" shadow="sm">
            <Group justify="space-between" mb="sm">
              <Title order={3}>ミッション & ランダムTip</Title>
              <Group gap="xs">
                <Button size="xs" variant="light" leftSection={<IconSparkles size={14} />} onClick={pickMission}>
                  ミッションを引く
                </Button>
                <Button size="xs" variant="light" leftSection={<IconBulb size={14} />} onClick={pickTip}>
                  Tipを引く
                </Button>
              </Group>
            </Group>
            <Stack gap="sm">
              <Paper shadow="xs" radius="md" p="md" withBorder style={{ background: '#f8f9fa' }}>
                <Group gap="sm">
                  <ThemeIcon variant="light" color="violet" radius="xl">
                    <IconMoodSmile size={18} />
                  </ThemeIcon>
                  <Text fw={600}>今日のミッション</Text>
                </Group>
                <Text mt="xs">{randomMission || 'ボタンを押して今日のミッションを取得しよう！'}</Text>
              </Paper>
              <Paper shadow="xs" radius="md" p="md" withBorder style={{ background: '#f8f9fa' }}>
                <Group gap="sm">
                  <ThemeIcon variant="light" color="yellow" radius="xl">
                    <IconWand size={18} />
                  </ThemeIcon>
                  <Text fw={600}>ランダムTip</Text>
                </Group>
                <Text mt="xs">{randomTip || 'ボタンを押してヒントを取得しよう！'}</Text>
              </Paper>
            </Stack>
          </Card>
        </Grid.Col>
        <Grid.Col span={{ base: 12, md: 4 }}>
          <Card withBorder radius="md" padding="lg" shadow="sm">
            <Title order={3} mb="sm">イベント・最新情報</Title>
            <Stack gap="xs">
              <Badge color="pink" variant="light">コミュニティ</Badge>
              <Text>・毎週末「でかプ交流会」実施中（詳細はDiscordで告知）。</Text>
              <Text>・グループに参加してプッシュ通知を受け取ろう。</Text>
              <Badge color="indigo" variant="light" mt="sm">ワールドTips</Badge>
              <Text>・ジャックポットの種は拾い忘れ注意！</Text>
              <Text>・スキルポイントの割り振りで立ち回りが大きく変わるよ。</Text>
            </Stack>
          </Card>
        </Grid.Col>
      </Grid>

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
          {/* 個人統計情報セクション */}
          <Paper p="xl" shadow="sm" radius="md">
            <Title order={2} mb="md" ta="center">
              個人統計情報
            </Title>
            <Text size="sm" c="dimmed" mb="md" ta="center">
              あなたのLoadSaveDataURLを入力してください
            </Text>
            
            <Group>
              <TextInput
                placeholder="https://push.trap.games/api/v3/users/username/data?sig=..."
                value={dataUrl}
                onChange={(e) => setDataUrl(e.target.value)}
                style={{ flex: 1 }}
                size="md"
              />
              <Button 
                onClick={handleLoadPersonalData}
                loading={isLoadingPersonal}
                leftSection={<IconDownload size="1rem" />}
                size="md"
              >
                データをロード
              </Button>
            </Group>

            {error && (
              <Alert 
                icon={<IconAlertCircle size="1rem" />} 
                title="注意" 
                color="orange" 
                mt="md"
              >
                {error}
              </Alert>
            )}

            {isLoadingPersonal && (
              <Center mt="md">
                <Loader size="lg" />
              </Center>
            )}

            {personalStats && (
              <Stack gap="lg" mt="xl">
                {/* 基本情報 */}
                <div>
                  <Title order={3} mb="md">基本情報</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">プレイ時間</Text>
                        <Text size="lg" fw={700} c="blue">
                          {personalStats.playtime ? `${Math.floor(personalStats.playtime / 3600)}時間 ${Math.floor((personalStats.playtime % 3600) / 60)}分` : 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">バージョン</Text>
                        <Text size="lg" fw={700}>{personalStats.version || 'N/A'}</Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">実績数</Text>
                        <Text size="lg" fw={700} c="yellow">
                          {personalStats.l_achieve?.length || 0}個
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>

                {/* メダル・クレジット関連 */}
                <div>
                  <Title order={3} mb="md">メダル・クレジット</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">クレジット</Text>
                        <Text size="lg" fw={700} c="green">
                          {personalStats.credit?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">累計クレジット</Text>
                        <Text size="lg" fw={700} c="green">
                          {personalStats.credit_all?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">投入メダル</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.medal_in?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">獲得メダル</Text>
                        <Text size="lg" fw={700} c="teal">
                          {personalStats.medal_get?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>

                {/* ボール関連 */}
                <div>
                  <Title order={3} mb="md">ボール</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">獲得ボール</Text>
                        <Text size="lg" fw={700} c="orange">
                          {personalStats.ball_get?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">最大ボールチェーン</Text>
                        <Text size="lg" fw={700} c="orange">
                          {personalStats.ball_chain || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>

                {/* スロット関連 */}
                <div>
                  <Title order={3} mb="md">スロット</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">スロット開始回数</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.slot_start?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">スロットヒット回数</Text>
                        <Text size="lg" fw={700} c="pink">
                          {personalStats.slot_hit?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">フィーバー開始回数</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.slot_startfev?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">フィーバー獲得回数</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.slot_getfev?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>

                {/* すごろく関連 */}
                <div>
                  <Title order={3} mb="md">すごろく</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">すごろく獲得</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.sqr_get?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">すごろく歩数</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.sqr_step?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>

                {/* ジャックポット関連 */}
                <div>
                  <Title order={3} mb="md">ジャックポット</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">ジャックポット獲得</Text>
                        <Text size="lg" fw={700} c="yellow">
                          {personalStats.jack_get?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">ジャックポットスタート最大</Text>
                        <Text size="lg" fw={700} c="yellow">
                          {personalStats.jack_startmax?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">ジャックポット合計最大 v2</Text>
                        <Text size="lg" fw={700} c="yellow">
                          {personalStats.jack_totalmax_v2?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>

                {/* ジャックポットSP関連 */}
                <div>
                  <Title order={3} mb="md">ジャックポットSP</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">ジャックポットSP合計獲得</Text>
                        <Text size="lg" fw={700} c="grape">
                          {personalStats.jacksp_get_all?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">ジャックポットSPスタート最大</Text>
                        <Text size="lg" fw={700} c="grape">
                          {personalStats.jacksp_startmax?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">ジャックポットSP合計最大</Text>
                        <Text size="lg" fw={700} c="grape">
                          {personalStats.jacksp_totalmax?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>

                {/* ウルティメイト関連 */}
                <div>
                  <Title order={3} mb="md">ウルティメイト</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">ウルティメイト獲得</Text>
                        <Text size="lg" fw={700} c="red">
                          {personalStats.ult_get?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">ウルティメイトコンボ最大</Text>
                        <Text size="lg" fw={700} c="red">
                          {personalStats.ult_combomax?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">ウルティメイト合計最大 v2</Text>
                        <Text size="lg" fw={700} c="red">
                          {personalStats.ult_totalmax_v2?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>

                {/* パルボール関連 */}
                <div>
                  <Title order={3} mb="md">パルボール</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">パルボール獲得</Text>
                        <Text size="lg" fw={700} c="violet">
                          {personalStats.palball_get?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">パルロット T0</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.pallot_lot_t0?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">パルロット T1</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.pallot_lot_t1?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">パルロット T2</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.pallot_lot_t2?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">パルロット T3</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.pallot_lot_t3?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>

                {/* その他の統計 */}
                <div>
                  <Title order={3} mb="md">その他</Title>
                  <Grid>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">最高CPM</Text>
                        <Text size="lg" fw={700} c="cyan">
                          {personalStats.cpm_max?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">SP使用回数</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.sp_use?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">シルベ購入</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.buy_shbi?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">虹メダルシルベ獲得</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.rmshbi_get?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">購入合計</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.buy_total?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">バトルパス歩数</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.bstp_step?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">バトルパス報酬</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.bstp_rwd?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                    <Grid.Col span={{ base: 12, md: 6, lg: 4 }}>
                      <Card shadow="sm" padding="md" radius="md" withBorder>
                        <Text size="sm" c="dimmed">タスクカウント</Text>
                        <Text size="lg" fw={700}>
                          {personalStats.task_cnt?.toLocaleString() || 'N/A'}
                        </Text>
                      </Card>
                    </Grid.Col>
                  </Grid>
                </div>
              </Stack>
            )}
          </Paper>
        </Tabs.Panel>

        <Tabs.Panel value="global" pt="md">
          {/* グローバル統計情報セクション */}
          <Paper p="xl" shadow="sm" radius="md">
            <Title order={2} mb="md" ta="center">
              グローバル統計情報
            </Title>
            
            {isLoadingGlobal && (
              <Center>
                <Loader size="lg" />
              </Center>
            )}

            {globalStats && (
              <Stack gap="xl">
                {/* 総メダル数 */}
                <Card shadow="sm" padding="lg" radius="md" withBorder>
                  <Title order={3} mb="md" ta="center">総メダル数</Title>
                  <Text size="2xl" fw={700} c="blue" ta="center">
                    {globalStats.total_medals.toLocaleString()}
                  </Text>
                </Card>

                {/* ランキング表示 */}
                <Grid>
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.achievements_count, '実績数ランキング', 'achievements_count')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.cpm_max, 'CPM最大ランキング', 'cpm_max')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.jacksp_startmax, 'ジャックポットスタート最大ランキング', 'jacksp_startmax')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.golden_palball_get, 'ゴールデンパルボール獲得ランキング', 'golden_palball_get')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.max_chain_rainbow, 'レインボーチェーン最大ランキング', 'max_chain_rainbow')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.jack_totalmax_v2, 'ジャックポット合計最大ランキング', 'jack_totalmax_v2')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.ult_combomax, 'ウルティメイトコンボ最大ランキング', 'ult_combomax')}
                  </Grid.Col>
                
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.ult_totalmax_v2, 'ウルティメイト合計最大ランキング', 'ult_totalmax_v2')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.blackbox_total, 'ブラックボックス総獲得ランキング', 'blackbox_total')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.sp_use, 'スペシャル使用回数ランキング', 'sp_use')}
                  </Grid.Col>
                </Grid>
              </Stack>
            )}
          </Paper>
        </Tabs.Panel>
      </Tabs>

      {/* Twitter セクション */}
      <Card withBorder radius="md" padding="lg" shadow="sm">
        <Group justify="space-between" mb="xs">
          <Title order={3}>#でかプ / #VRでかプ リアルタイム</Title>
          <Button
            size="xs"
            variant="light"
            component="a"
            href={twitterSearchUrl}
            target="_blank"
            rel="noreferrer"
            leftSection={<IconExternalLink size={14} />}
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
            href={twitterSearchUrl}
          >
            Tweets by #でかプ
          </a>
        </div>
      </Card>

      {/* リソース & FAQ */}
      <Grid gutter="md">
        <Grid.Col span={{ base: 12, md: 7 }}>
          <Card withBorder radius="md" padding="lg" shadow="sm">
            <Title order={3} mb="sm">ガイド & ヒント</Title>
            <Stack gap="sm">
              <Text>
                <strong>クラウドセーブ手順:</strong> ワールド内パネルでセーブ → URL をコピーして本ページの「個人統計」に貼り付け → 最新統計を表示。
              </Text>
              <Text>
                <strong>ロード確認:</strong> ワールド内「ロード確認モーダル」でセーブ差分を確認後にロード。改ざん防止のため HMAC 署名を必須化。
              </Text>
              <Text>
                <strong>API で直接触る:</strong> Swagger UI から /v4/data, /v4/users/{'{user_id}'}/data を試せます。セーブは Base64URL、ロードは Base64 + 署名です。
              </Text>
              <Text>
                <strong>ランキング反映:</strong> 最新セーブが v3_user_latest_save_data に集約され、v4/statistics が高速に返却されます。
              </Text>
            </Stack>
          </Card>
        </Grid.Col>
        <Grid.Col span={{ base: 12, md: 5 }}>
          <Card withBorder radius="md" padding="lg" shadow="sm">
            <Title order={3} mb="sm">コミュニティ・その他</Title>
            <Stack gap="xs">
              <Anchor href="https://discord.com/invite/CgnYyXecKm" target="_blank" rel="noreferrer" c="blue">
                公式Discord (でかプ同好会 Discord支部)
              </Anchor>
              <Anchor href="https://wikiwiki.jp/vr_bigpusher/" target="_blank" rel="noreferrer" c="blue">
                公式Wiki
              </Anchor>
              <Anchor href="https://vrchat.com/home/group/grp_5900a25d-0bb9-48d4-bab1-f3bd5c9a5e73" target="_blank" rel="noreferrer" c="blue">
                VRChatグループ
              </Anchor>
              <Anchor href="https://vrchat.com/home/group/grp_f38ec6a3-0de5-499e-a85f-1038013bdd04" target="_blank" rel="noreferrer" c="blue">
                でかプ交流会 ～ MMP Meeting
              </Anchor>
              <Anchor href="https://github.com/pikachu0310/very-big-medal-pusher-data-server" target="_blank" rel="noreferrer" c="dimmed">
                GitHub (サーバー)
              </Anchor>
            </Stack>
          </Card>
        </Grid.Col>
      </Grid>

      {/* 開発者向け（折りたたみ） */}
      <Accordion variant="contained" radius="md">
        <Accordion.Item value="dev">
          <Accordion.Control>開発者向けツール</Accordion.Control>
          <Accordion.Panel>
            <Grid gutter="md">
              <Grid.Col span={{ base: 12, md: 6 }}>
                <Card withBorder radius="md" padding="lg" shadow="sm">
                  <Group justify="space-between" mb="sm">
                    <Title order={4}>クイックリンク</Title>
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
                          <IconDownload size={16} />
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
                          <IconDownload size={16} />
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
                          <IconDownload size={16} />
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
                      >
                        Swagger UI
                      </Button>
                      <Button
                        variant="subtle"
                        component="a"
                        href="/api/openapi.yaml"
                        target="_blank"
                        radius="md"
                      >
                        OpenAPI
                      </Button>
                    </Group>
                  </Stack>
                </Card>
              </Grid.Col>
              <Grid.Col span={{ base: 12, md: 6 }}>
                <Card withBorder radius="md" padding="lg" shadow="sm">
                  <Group justify="space-between" align="center">
                    <Title order={4}>サーバーヘルス</Title>
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
          </Accordion.Panel>
        </Accordion.Item>
      </Accordion>
    </Stack>
  );
}

export default HomePage;

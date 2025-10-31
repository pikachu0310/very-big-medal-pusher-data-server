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
  Badge
} from '@mantine/core';
import { IconAlertCircle, IconDownload, IconTrophy, IconUsers } from '@tabler/icons-react';

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
      {/* ロゴセクション */}
      <Center mt={0} pt={0} mb={0}>
        <img 
          src="/MPG_logo.png" 
          alt="MPG Logo" 
          className="logo-container"
          style={{ 
            maxWidth: '700px', 
            width: '100%',
            height: 'auto',
            marginBottom: '0.5rem',
            marginTop: '0'
          }} 
        />
      </Center>

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
    </Stack>
  );
}

export default HomePage;

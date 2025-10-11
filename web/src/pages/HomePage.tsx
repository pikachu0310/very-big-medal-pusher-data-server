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

// 個人統計情報の型定義
interface PersonalStats {
  playtime?: number;
  medal_get?: number;
  cpm_max?: number;
  credit?: number;
  ball_get?: number;
  slot_hit?: number;
  ult_get?: number;
  jack_get?: number;
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
      
      // APIから取得したデータをセット
      setPersonalStats({
        playtime: data.playtime,
        medal_get: data.medal_get,
        cpm_max: data.cpm_max,
        credit: data.credit,
        ball_get: data.ball_get,
        slot_hit: data.slot_hit,
        ult_get: data.ult_get,
        jack_get: data.jack_get
      });
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
                  <Badge color={index < 3 ? ['gold', 'silver', 'bronze'][index] : 'gray'}>
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
    <Stack gap="xl">
      {/* ロゴセクション */}
      <Center>
        <img 
          src="/MPG_logo.png" 
          alt="MPG Logo" 
          className="logo-container"
          style={{ 
            maxWidth: '500px', 
            height: 'auto',
            marginBottom: '2rem'
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
              <Grid mt="xl">
                <Grid.Col span={{ base: 12, md: 6 }}>
                  <Card shadow="sm" padding="lg" radius="md" withBorder>
                    <Text size="lg" fw={500} mb="sm" c="dimmed">
                      プレイ時間
                    </Text>
                    <Text size="xl" fw={700} c="blue">
                      {personalStats.playtime ? `${Math.floor(personalStats.playtime / 3600)}時間` : 'N/A'}
                    </Text>
                  </Card>
                </Grid.Col>
                
                <Grid.Col span={{ base: 12, md: 6 }}>
                  <Card shadow="sm" padding="lg" radius="md" withBorder>
                    <Text size="lg" fw={500} mb="sm" c="dimmed">
                      獲得メダル数
                    </Text>
                    <Text size="xl" fw={700} c="green">
                      {personalStats.medal_get?.toLocaleString() || 'N/A'}
                    </Text>
                  </Card>
                </Grid.Col>
                
                <Grid.Col span={{ base: 12, md: 6 }}>
                  <Card shadow="sm" padding="lg" radius="md" withBorder>
                    <Text size="lg" fw={500} mb="sm" c="dimmed">
                      最高CPM
                    </Text>
                    <Text size="xl" fw={700} c="orange">
                      {personalStats.cpm_max?.toLocaleString() || 'N/A'}
                    </Text>
                  </Card>
                </Grid.Col>
                
                <Grid.Col span={{ base: 12, md: 6 }}>
                  <Card shadow="sm" padding="lg" radius="md" withBorder>
                    <Text size="lg" fw={500} mb="sm" c="dimmed">
                      クレジット
                    </Text>
                    <Text size="xl" fw={700} c="purple">
                      {personalStats.credit?.toLocaleString() || 'N/A'}
                    </Text>
                  </Card>
                </Grid.Col>
              </Grid>
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
                    {renderRankingTable(globalStats.cpm_max, 'CPMランキング', 'cpm_max')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.jacksp_startmax, 'ジャックポットスタート最大ランキング', 'jacksp_startmax')}
                  </Grid.Col>
                  
                  <Grid.Col span={{ base: 12, md: 6 }}>
                    {renderRankingTable(globalStats.ult_combomax, 'ウルティメイトコンボ最大ランキング', 'ult_combomax')}
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

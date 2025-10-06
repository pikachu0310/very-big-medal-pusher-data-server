import { Title, Text, Paper, Stack } from '@mantine/core';
import { Layout } from '../components/Layout';

function StatsPage() {
  return (
    <Layout title="統計" description="Very Big Medal Pusherの統計データを表示します">
      <Paper p="xl" shadow="sm" radius="md">
        <Title order={2} mb="md">
          統計データ
        </Title>
        <Text>
          統計データの表示機能は今後実装予定です。
        </Text>
      </Paper>
    </Layout>
  );
}

export default StatsPage;

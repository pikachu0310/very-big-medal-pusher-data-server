import { Title, Text, Paper } from '@mantine/core';
import { Layout } from '../components/Layout';

function HomePage() {
  return (
    <Layout title="Very Big Medal Pusher" description="データサーバーへようこそ">
      <Paper p="xl" shadow="sm" radius="md">
        <Title order={2} mb="md">
          機能
        </Title>
        <Text>
          このアプリケーションは、Very Big Medal Pusherのデータを管理・表示するためのサーバーです。
        </Text>
      </Paper>
    </Layout>
  );
}

export default HomePage;

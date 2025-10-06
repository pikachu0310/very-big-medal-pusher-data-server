import { Title, Text, Paper, Stack } from '@mantine/core';
import { Layout } from '../components/Layout';

function SettingsPage() {
  return (
    <Layout title="設定" description="アプリケーションの設定を行います">
      <Paper p="xl" shadow="sm" radius="md">
        <Title order={2} mb="md">
          アプリケーション設定
        </Title>
        <Text>
          設定機能は今後実装予定です。
        </Text>
      </Paper>
    </Layout>
  );
}

export default SettingsPage;

import { Stack, Title, Text, Paper, List, Anchor } from '@mantine/core';
import { Link } from 'react-router';

function PrivacyPolicyPage() {
  return (
    <Stack gap="xl">
      <Paper p="xl" shadow="sm" radius="md">
        <Anchor component={Link} to="/" size="sm" c="dimmed" underline="hover" mb="md" style={{ display: 'inline-block' }}>
          ← ホームに戻る
        </Anchor>
        <Title order={1} mb="xl" ta="center">
          プライバシーポリシー
        </Title>

        <Stack gap="lg">
          <div>
            <Title order={2} mb="md">1. 収集する情報</Title>
            <Text>
              本サービスは、VRChatワールド「クソでっけぇプッシャーゲーム」のプレイデータを収集・保存します。
              収集されるデータには以下が含まれます：
            </Text>
            <List mt="sm">
              <List.Item>VRChatユーザー名</List.Item>
              <List.Item>ゲームプレイに関する統計情報（メダル数、プレイ時間、スコアなど）</List.Item>
              <List.Item>実績データ</List.Item>
              <List.Item>セーブデータの更新日時</List.Item>
            </List>
          </div>

          <div>
            <Title order={2} mb="md">2. 情報の利用目的</Title>
            <Text>
              収集した情報は以下の目的で利用されます：
            </Text>
            <List mt="sm">
              <List.Item>プレイヤーのゲームデータの保存と復元</List.Item>
              <List.Item>グローバルランキングの生成と表示</List.Item>
              <List.Item>ゲームの統計情報の提供</List.Item>
              <List.Item>サービスの改善と機能追加</List.Item>
            </List>
          </div>

          <div>
            <Title order={2} mb="md">3. 情報の共有と開示</Title>
            <Text>
              本サービスでは、以下の情報が公開されます：
            </Text>
            <List mt="sm">
              <List.Item>グローバルランキング（ユーザー名と統計値）</List.Item>
              <List.Item>実績取得率などの集計情報</List.Item>
            </List>
            <Text mt="sm">
              個人を特定できる情報を第三者に販売または貸与することはありません。
            </Text>
          </div>

          <div>
            <Title order={2} mb="md">4. データの保存期間</Title>
            <Text>
              収集したデータは、サービスの提供に必要な期間保存されます。
              ユーザーがデータの削除を希望する場合は、運営者にお問い合わせください。
            </Text>
          </div>

          <div>
            <Title order={2} mb="md">5. セキュリティ</Title>
            <Text>
              本サービスは、収集したデータを保護するために適切なセキュリティ対策を実施しています。
              データの送受信にはHMAC-SHA256署名による認証を使用しています。
            </Text>
          </div>

          <div>
            <Title order={2} mb="md">6. プライバシーポリシーの変更</Title>
            <Text>
              本プライバシーポリシーは、法令の変更やサービスの改善に伴い、予告なく変更される場合があります。
              変更後のプライバシーポリシーは、本ページに掲載された時点で効力を生じるものとします。
            </Text>
          </div>

          <div>
            <Title order={2} mb="md">7. お問い合わせ</Title>
            <Text>
              本プライバシーポリシーに関するご質問やご意見、データの削除依頼などは、
              サービス運営者までお問い合わせください。
            </Text>
          </div>

          <Text size="sm" c="dimmed" mt="xl" ta="right">
            最終更新日：2025年10月11日
          </Text>
        </Stack>
      </Paper>
    </Stack>
  );
}

export default PrivacyPolicyPage;


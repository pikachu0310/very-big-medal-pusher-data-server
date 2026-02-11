import { useState } from 'react';
import {
  Card,
  Grid,
  Group,
  Title,
  Button,
  Text,
  Badge,
  ThemeIcon,
  Anchor,
  Stack
} from '@mantine/core';
import {
  IconPlugConnected,
  IconBrandGithub,
  IconLock
} from '@tabler/icons-react';

function DeveloperToolsSection() {
  const [pingState, setPingState] = useState<'idle' | 'ok' | 'ng' | 'loading'>('idle');

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

  return (
    <Grid gutter="md">
      <Grid.Col span={{ base: 12, md: 6 }}>
        <Card withBorder radius="md" padding="lg" shadow="sm">
          <Group justify="space-between" align="center">
            <Title order={3}>サーバーヘルス</Title>
            <Button
              size="sm"
              variant="outline"
              color="black"
              className="mmp-outline-black-button"
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
            <Badge color={pingState === 'ok' ? 'teal' : pingState === 'ng' ? 'red' : 'gray'} radius="sm" aria-live="polite">
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
  );
}

export default DeveloperToolsSection;

import { Container, Title, Text, Stack, Paper } from '@mantine/core';

interface LayoutProps {
  children: React.ReactNode;
  title?: string;
  description?: string;
}

export function Layout({ children, title, description }: LayoutProps) {
  return (
    <Container size="lg" py="xl">
      <Stack gap="xl">
        {(title || description) && (
          <Paper p="xl" shadow="sm" radius="md">
            {title && (
              <Title order={1} mb="md">
                {title}
              </Title>
            )}
            {description && (
              <Text size="lg" c="dimmed">
                {description}
              </Text>
            )}
          </Paper>
        )}
        {children}
      </Stack>
    </Container>
  );
}

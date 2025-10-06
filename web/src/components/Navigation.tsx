import { NavLink, Stack, Title } from '@mantine/core';
import { IconHome, IconChartBar, IconSettings } from '@tabler/icons-react';
import { Link, useMatches } from 'react-router';

const navigationItems = [
  { label: 'ホーム', href: '/', icon: IconHome },
  { label: '統計', href: '/stats', icon: IconChartBar },
  { label: '設定', href: '/settings', icon: IconSettings },
];

export function Navigation() {
  const matches = useMatches();
  const currentPath = matches[matches.length - 1]?.pathname || '/';

  return (
    <>
      <Title order={3} mb="md">
        Very Big Medal Pusher
      </Title>
      <Stack gap="xs">
        {navigationItems.map((item) => (
          <NavLink
            key={item.href}
            component={Link}
            to={item.href}
            label={item.label}
            leftSection={<item.icon size="1rem" />}
            active={currentPath === item.href}
          />
        ))}
      </Stack>
    </>
  );
}

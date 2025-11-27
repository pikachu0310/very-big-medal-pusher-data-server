import { Container, Box, Text, Anchor } from '@mantine/core';
import { Link, Route, Routes } from 'react-router';
import HomePage from './pages/HomePage';
import PrivacyPolicyPage from './pages/PrivacyPolicyPage';

function App() {
  return (
    <Box style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
      <Container size="lg" style={{ flex: 1, padding: '1rem 1rem 2rem 1rem' }}>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/privacy" element={<PrivacyPolicyPage />} />
        </Routes>
      </Container>
      
      {/* フッター */}
      <Box
        component="footer"
        style={{
          padding: '1rem',
          textAlign: 'center',
          borderTop: '1px solid #e9ecef',
          marginTop: 'auto'
        }}
      >
        <Text size="xs" c="dimmed" style={{ display: 'flex', justifyContent: 'center', gap: '0.5rem', alignItems: 'center', flexWrap: 'wrap' }}>
          <Anchor component={Link} to="/privacy" size="xs" c="dimmed" underline="hover">
            プライバシーポリシー
          </Anchor>
          <Text span c="dimmed">|</Text>
          <Anchor href="https://github.com/pikachu0310/very-big-medal-pusher-data-server" target="_blank" rel="noreferrer" size="xs" c="dimmed" underline="hover">
            GitHub
          </Anchor>
          <Text span c="dimmed">|</Text>
          <Text span c="dimmed">© 2025 Massive Medal Pusher</Text>
        </Text>
      </Box>
    </Box>
  );
}

export default App;

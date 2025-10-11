import { Container, Box, Text, Anchor } from '@mantine/core';
import { Link, Route, Routes } from 'react-router';
import HomePage from './pages/HomePage';
import PrivacyPolicyPage from './pages/PrivacyPolicyPage';

function App() {
  return (
    <Box style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
      <Container size="lg" style={{ flex: 1, padding: '2rem 1rem' }}>
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
        <Text size="xs" c="dimmed">
          <Anchor component={Link} to="/privacy" size="xs" c="dimmed" underline="hover">
            プライバシーポリシー
          </Anchor>
          {' | '}
          © 2025 Very Big Medal Pusher
        </Text>
      </Box>
    </Box>
  );
}

export default App;

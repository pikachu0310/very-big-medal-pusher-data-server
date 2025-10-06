import { Routes, Route } from 'react-router';
import HomePage from './HomePage';
import StatsPage from './StatsPage';
import SettingsPage from './SettingsPage';

export function AppRouter() {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/stats" element={<StatsPage />} />
      <Route path="/settings" element={<SettingsPage />} />
    </Routes>
  );
}

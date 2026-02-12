import { lazy, Suspense } from 'react';
import { Link, Route, Routes } from 'react-router';
import HomePage from './pages/HomePage';

const PrivacyPolicyPage = lazy(() => import('./pages/PrivacyPolicyPage'));

function App() {
  return (
    <div className="app-shell">
      <a href="#main-content" className="skip-link">
        メインコンテンツへスキップ
      </a>

      <main id="main-content" className="app-main-container">
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/privacy" element={<Suspense fallback={null}><PrivacyPolicyPage /></Suspense>} />
        </Routes>
      </main>

      <footer className="app-footer">
        <p className="app-footer-text">
          <Link className="app-footer-link" to="/privacy">
            プライバシーポリシー
          </Link>
          <span aria-hidden="true">|</span>
          <a className="app-footer-link" href="https://github.com/pikachu0310/very-big-medal-pusher-data-server" target="_blank" rel="noreferrer">
            GitHub
          </a>
          <span aria-hidden="true">|</span>
          <span>© 2025 Massive Medal Pusher</span>
        </p>
      </footer>
    </div>
  );
}

export default App;

import { lazy, Suspense } from 'react';
import { Link, Route, Routes } from 'react-router';
import HomePage from './pages/HomePage';
import ClassicHomePage from './pages/ClassicHomePage';

const PrivacyPolicyPage = lazy(() => import('./pages/PrivacyPolicyPage'));

function App() {
  return (
    <div className="app-shell">
      <a href="#main-content" className="skip-link">
        メインコンテンツへスキップ
      </a>

      <header className="app-global-notice" aria-label="イベント通知">
        <p>春イベント期間中 (2026年4月1日 - 2026年4月13日): 演出強化版は `/`、通常版は `/classic` です。</p>
      </header>

      <main id="main-content" className="app-main-container">
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/classic" element={<ClassicHomePage />} />
          <Route path="/privacy" element={<Suspense fallback={null}><PrivacyPolicyPage /></Suspense>} />
        </Routes>
      </main>

      <footer className="app-footer">
        <p className="app-footer-text">
          <Link className="app-footer-link" to="/privacy">
            プライバシーポリシー(やや真面目)
          </Link>
          <span aria-hidden="true">|</span>
          <Link className="app-footer-link" to="/classic">
            通常版ページ
          </Link>
          <span aria-hidden="true">|</span>
          <a className="app-footer-link" href="https://github.com/pikachu0310/very-big-medal-pusher-data-server" target="_blank" rel="noreferrer">
            GitHub
          </a>
          <span aria-hidden="true">|</span>
          <span>© 2026 Massive Medal Pusher April Fools Dept.</span>
        </p>
      </footer>
    </div>
  );
}

export default App;

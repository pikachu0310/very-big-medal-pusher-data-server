import { lazy, Suspense, useEffect } from 'react';
import { Link, Route, Routes, useLocation } from 'react-router';
import HomePage from './pages/HomePage';
import ClassicHomePage from './pages/ClassicHomePage';

const PrivacyPolicyPage = lazy(() => import('./pages/PrivacyPolicyPage'));
const aprilFoolsPath = '/april-fools';

function App() {
  const location = useLocation();
  const isAprilFoolsRoute = location.pathname.startsWith(aprilFoolsPath);
  const isClassicMode = !isAprilFoolsRoute;

  useEffect(() => {
    document.body.classList.toggle('classic-mode-body', isClassicMode);
    return () => document.body.classList.remove('classic-mode-body');
  }, [isClassicMode]);

  return (
    <div className="app-shell">
      <a href="#main-content" className="skip-link">
        メインコンテンツへスキップ
      </a>

      {!isAprilFoolsRoute && (
        <header className="app-global-notice" aria-label="イベント通知">
          <p>エイプリルフール特別版はこちら！</p>
          <Link className="app-april-cta-link" to={aprilFoolsPath}>
            ド派手版を見る
          </Link>
        </header>
      )}

      <main id="main-content" className={`app-main-container ${isClassicMode ? 'classic-main-container' : ''}`}>
        <Routes>
          <Route path="/" element={<ClassicHomePage />} />
          <Route path="/classic" element={<ClassicHomePage />} />
          <Route path={aprilFoolsPath} element={<HomePage />} />
          <Route path="/privacy" element={<Suspense fallback={null}><PrivacyPolicyPage /></Suspense>} />
        </Routes>
      </main>

      <footer className="app-footer">
        <p className="app-footer-text">
          <Link className="app-footer-link" to="/privacy">
            プライバシーポリシー
          </Link>
          {isAprilFoolsRoute && (
            <>
              <span aria-hidden="true">|</span>
              <Link className="app-footer-link" to="/">
                通常版ページ
              </Link>
            </>
          )}
          <span aria-hidden="true">|</span>
          <a className="app-footer-link" href="https://github.com/pikachu0310/very-big-medal-pusher-data-server" target="_blank" rel="noreferrer">
            GitHub
          </a>
          <span aria-hidden="true">|</span>
          <span>{isAprilFoolsRoute ? '© 2026 Massive Medal Pusher April Fools Dept.' : '© 2025 Massive Medal Pusher'}</span>
        </p>
      </footer>
    </div>
  );
}

export default App;

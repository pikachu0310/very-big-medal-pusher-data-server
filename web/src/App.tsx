import { lazy, Suspense, useEffect } from 'react';
import { Link, Route, Routes, useLocation } from 'react-router';
import HomePage from './pages/HomePage';
import ClassicHomePage from './pages/ClassicHomePage';

const PrivacyPolicyPage = lazy(() => import('./pages/PrivacyPolicyPage'));

function App() {
  const location = useLocation();
  const isClassicRoute = location.pathname.startsWith('/classic');

  useEffect(() => {
    document.body.classList.toggle('classic-mode-body', isClassicRoute);
    return () => document.body.classList.remove('classic-mode-body');
  }, [isClassicRoute]);

  return (
    <div className="app-shell">
      <a href="#main-content" className="skip-link">
        メインコンテンツへスキップ
      </a>

      {!isClassicRoute && (
        <header className="app-global-notice" aria-label="イベント通知">
          <p>Happy April Fool&apos;s Day!! にぎやか版は `/`、通常版は `/classic` で見られます。</p>
        </header>
      )}

      <main id="main-content" className={`app-main-container ${isClassicRoute ? 'classic-main-container' : ''}`}>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/classic" element={<ClassicHomePage />} />
          <Route path="/privacy" element={<Suspense fallback={null}><PrivacyPolicyPage /></Suspense>} />
        </Routes>
      </main>

      <footer className="app-footer">
        <p className="app-footer-text">
          <Link className="app-footer-link" to="/privacy">
            {isClassicRoute ? 'プライバシーポリシー' : 'プライバシーポリシー(やや真面目)'}
          </Link>
          {!isClassicRoute && (
            <>
              <span aria-hidden="true">|</span>
              <Link className="app-footer-link" to="/classic">
                通常版ページ
              </Link>
            </>
          )}
          <span aria-hidden="true">|</span>
          <a className="app-footer-link" href="https://github.com/pikachu0310/very-big-medal-pusher-data-server" target="_blank" rel="noreferrer">
            GitHub
          </a>
          <span aria-hidden="true">|</span>
          <span>{isClassicRoute ? '© 2025 Massive Medal Pusher' : '© 2026 Massive Medal Pusher April Fools Dept.'}</span>
        </p>
      </footer>
    </div>
  );
}

export default App;

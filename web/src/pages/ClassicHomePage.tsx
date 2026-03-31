import { lazy, Suspense, useEffect, useRef, useState } from 'react';
import { IconBook2, IconBrandDiscord, IconBrandGithub, IconWorld } from '@tabler/icons-react';
import { MmpOutlineButton, MmpPrimaryButton } from '../components/MmpButton';

const DeferredSections = lazy(() => import('../components/DeferredSections'));

function DeferredPlaceholders() {
  return (
    <>
      <section className="deferred-placeholder deferred-placeholder-lg" aria-live="polite" role="status">
        <span className="spinner" aria-hidden="true" />
      </section>
      <section className="deferred-placeholder" aria-live="polite" role="status">
        <span className="spinner spinner-sm" aria-hidden="true" />
      </section>
    </>
  );
}

function ClassicHomePage() {
  const [showDeferredSections, setShowDeferredSections] = useState(false);
  const deferredSectionAnchorRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    if (showDeferredSections) {
      return;
    }

    let fallbackTimeoutId: number | undefined;
    const revealDeferredSections = () => setShowDeferredSections(true);
    const anchorElement = deferredSectionAnchorRef.current;

    if (anchorElement && typeof window.IntersectionObserver === 'function') {
      const observer = new window.IntersectionObserver(
        (entries) => {
          if (entries.some((entry) => entry.isIntersecting)) {
            revealDeferredSections();
            observer.disconnect();
          }
        },
        { rootMargin: '320px 0px' }
      );
      observer.observe(anchorElement);
      return () => observer.disconnect();
    }

    fallbackTimeoutId = window.setTimeout(revealDeferredSections, 1200);
    return () => {
      if (typeof fallbackTimeoutId === 'number') {
        window.clearTimeout(fallbackTimeoutId);
      }
    };
  }, [showDeferredSections]);

  return (
    <div className="home-stack classic-page">
      <section className="classic-hero">
        <p className="classic-badge">Classic Mode</p>
        <h1 className="classic-title">クソでっけぇプッシャーゲーム 公式ウェブサイト</h1>
        <p className="classic-subtitle">
          通常版ページです。イベント演出あり版は <a href="/" className="classic-inline-link">トップページ</a> から開けます。
        </p>
      </section>

      <section className="link-card classic-links-card">
        <h2 className="section-title">でかプ公式リンク集</h2>
        <div className="link-grid link-grid-two">
          <div className="link-column">
            <MmpPrimaryButton
              href="https://discord.com/invite/CgnYyXecKm"
              target="_blank"
              icon={<IconBrandDiscord size={32} />}
              size="xl"
              className="mmp-link-hero-button"
              heightMultiplier={1.35}
            >
              公式Discord
            </MmpPrimaryButton>
            <MmpPrimaryButton
              href="https://wikiwiki.jp/vr_bigpusher/"
              target="_blank"
              icon={<IconBook2 size={22} />}
              size="lg"
              className="mmp-link-hero-button"
            >
              公式Wiki
            </MmpPrimaryButton>
          </div>
          <div className="link-column">
            <MmpPrimaryButton
              href="https://vrchat.com/home/group/grp_5900a25d-0bb9-48d4-bab1-f3bd5c9a5e73"
              target="_blank"
              icon={<IconWorld size={22} />}
              size="lg"
              className="mmp-link-hero-button"
            >
              公式グループ(VRChat)
            </MmpPrimaryButton>
            <MmpOutlineButton
              href="https://github.com/pikachu0310/very-big-medal-pusher-data-server"
              target="_blank"
              icon={<IconBrandGithub size={20} />}
              size="lg"
            >
              Data Server GitHub
            </MmpOutlineButton>
          </div>
        </div>
      </section>

      <div ref={deferredSectionAnchorRef} className="deferred-anchor" aria-hidden="true" />

      {showDeferredSections ? (
        <Suspense fallback={<DeferredPlaceholders />}>
          <DeferredSections />
        </Suspense>
      ) : (
        <DeferredPlaceholders />
      )}
    </div>
  );
}

export default ClassicHomePage;

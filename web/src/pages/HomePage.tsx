import { lazy, Suspense, useEffect, useRef, useState } from 'react';
import {
  IconBrandDiscord,
  IconBrandGithub,
  IconWorld,
  IconBook2,
  IconExternalLink,
  IconLock
} from '@tabler/icons-react';
import { MmpOutlineButton, MmpPrimaryButton } from '../components/MmpButton';

const DeferredSections = lazy(() => import('../components/DeferredSections'));

const sectionTitleColor = '#1f5da8';
const pageTitleColor = '#2c4256';

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

function HomePage() {
  const [showDeferredSections, setShowDeferredSections] = useState(false);
  const deferredSectionAnchorRef = useRef<HTMLDivElement | null>(null);
  const twitterHashUrl = 'https://x.com/search?q=%23%E3%81%A7%E3%81%8B%E3%83%97%20OR%20%23VR%E3%81%A7%E3%81%8B%E3%83%97&f=live';

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
    <div className="home-stack">
      <div className="logo-center">
        <div className="logo-box">
          <picture>
            <source
              media="(max-width: 768px)"
              srcSet="/MMP_logo_480.webp 480w, /MMP_logo_596.webp 596w"
              sizes="(max-width: 768px) 66vw"
              type="image/webp"
            />
            <source
              media="(min-width: 769px)"
              srcSet="/MMP_logo_640.webp 640w, /MMP_logo_768.webp 768w, /MMP_logo_960.webp 960w, /MMP_logo_1192.webp 1192w"
              sizes="(max-width: 1200px) 48vw, 620px"
              type="image/webp"
            />
            <img
              src="/MMP_logo_480.webp"
              alt="Massive Medal Pusher ロゴ"
              width={1192}
              height={520}
              loading="eager"
              fetchPriority="high"
              decoding="async"
              className="home-logo-image"
            />
          </picture>
        </div>
      </div>

      <h1 className="home-title" style={{ color: pageTitleColor }}>
        クソでっけぇプッシャーゲーム 公式ウェブサイト
      </h1>
      <p className="home-subtitle">公式リンクや統計情報、開発者向けの情報をまとめて確認できます</p>

      <section className="link-card">
        <h2 className="section-title" style={{ color: sectionTitleColor }}>
          でかプ公式リンク集 / MMP Quick Links
        </h2>
        <div className="link-grid link-grid-two">
          <div className="link-column">
            <MmpPrimaryButton
              href="https://discord.com/invite/CgnYyXecKm"
              target="_blank"
              icon={<IconBrandDiscord size={36} />}
              size="xl"
              heightMultiplier={2}
              className="mmp-link-hero-button"
            >
              公式Discord でかプ同好会
            </MmpPrimaryButton>
            <MmpPrimaryButton
              href="https://wikiwiki.jp/vr_bigpusher/"
              target="_blank"
              icon={<IconBook2 size={24} />}
              size="xl"
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
              heightMultiplier={1.1}
              className="mmp-link-hero-button"
            >
              公式グループ(VRChat)
            </MmpPrimaryButton>
            <MmpPrimaryButton
              href="https://vrchat.com/home/launch?worldId=wrld_1af53798-92a3-4c3f-99ae-a7c42ec6084d"
              target="_blank"
              icon={<IconWorld size={22} />}
              size="lg"
              heightMultiplier={1.1}
              className="mmp-link-hero-button"
            >
              VRChatワールドリンク
            </MmpPrimaryButton>
            <MmpPrimaryButton
              href={twitterHashUrl}
              target="_blank"
              icon={<IconExternalLink size={22} />}
              size="lg"
              heightMultiplier={1.1}
              className="mmp-link-hero-button"
            >
              #でかプ / #VRでかプ (X投稿)
            </MmpPrimaryButton>
          </div>
        </div>
      </section>

      <section className="link-card">
        <h2 className="section-title" style={{ color: sectionTitleColor }}>
          開発者向けリンク集 / Links for Developers
        </h2>
        <div className="link-grid link-grid-three">
          <MmpOutlineButton
            href="/swagger/index.html"
            target="_blank"
            icon={<IconExternalLink size={18} />}
            size="lg"
            heightMultiplier={1.1}
          >
            SwaggerUI (API一覧)
          </MmpOutlineButton>
          <MmpOutlineButton
            href="https://push.trap.show/?server=mariadb.ns-system.svc.cluster.local&username=nsapp_c27d6f571f88ffff360fe2&db=nsapp_c27d6f571f88ffff360fe2"
            target="_blank"
            icon={<IconLock size={18} />}
            size="lg"
            heightMultiplier={1.1}
          >
            データベース
          </MmpOutlineButton>
          <MmpOutlineButton
            href="https://github.com/pikachu0310/very-big-medal-pusher-data-server"
            target="_blank"
            icon={<IconBrandGithub size={18} />}
            size="lg"
            heightMultiplier={1.1}
          >
            Data Server GitHub
          </MmpOutlineButton>
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

export default HomePage;

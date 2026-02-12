import { lazy, Suspense, useEffect, useRef, useState } from 'react';
import {
  IconBrandDiscord,
  IconBrandGithub,
  IconWorld,
  IconBook2,
  IconExternalLink,
  IconLock
} from '@tabler/icons-react';

const DeferredSections = lazy(() => import('../components/DeferredSections'));

const sectionTitleColor = '#1f5da8';
const pageTitleColor = '#2c4256';
const primaryButtonColor = 'blue';

type HeroButtonProps = {
  children: React.ReactNode;
  href: string;
  icon?: React.ReactNode;
  size?: 'lg' | 'xl';
  variant?: 'filled' | 'outline';
  heightMultiplier?: number;
  className?: string;
};

function HeroButton({
  children,
  href,
  icon,
  size = 'lg',
  variant = 'filled',
  heightMultiplier = 1,
  className
}: HeroButtonProps) {
  const mergedClassName = [
    'hero-link-button',
    size === 'xl' ? 'hero-link-button-xl' : 'hero-link-button-lg',
    variant === 'outline' ? 'hero-link-button-outline' : 'hero-link-button-filled',
    className
  ].filter(Boolean).join(' ');

  return (
    <a
      href={href}
      target="_blank"
      rel="noreferrer"
      className={mergedClassName}
      style={{ minHeight: `calc(${size === 'xl' ? 52 : 44}px * ${heightMultiplier})` }}
    >
      <span className="hero-link-button-content">
        {icon ? <span className="hero-link-button-icon" aria-hidden="true">{icon}</span> : null}
        <span className="hero-link-button-label">{children}</span>
      </span>
    </a>
  );
}

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
            <HeroButton
              href="https://discord.com/invite/CgnYyXecKm"
              icon={<IconBrandDiscord size={36} />}
              size="xl"
              variant="filled"
              heightMultiplier={2}
              className="mmp-link-hero-button"
            >
              公式Discord でかプ同好会
            </HeroButton>
            <HeroButton
              href="https://wikiwiki.jp/vr_bigpusher/"
              icon={<IconBook2 size={24} />}
              size="xl"
              variant="filled"
              className="mmp-link-hero-button"
            >
              公式Wiki
            </HeroButton>
          </div>
          <div className="link-column">
            <HeroButton
              href="https://vrchat.com/home/group/grp_5900a25d-0bb9-48d4-bab1-f3bd5c9a5e73"
              icon={<IconWorld size={22} />}
              size="lg"
              variant="filled"
              className="mmp-link-hero-button"
            >
              公式グループ(VRChat)
            </HeroButton>
            <HeroButton
              href="https://vrchat.com/home/launch?worldId=wrld_1af53798-92a3-4c3f-99ae-a7c42ec6084d"
              icon={<IconWorld size={22} />}
              size="lg"
              variant="filled"
              className="mmp-link-hero-button"
            >
              VRChatワールドリンク
            </HeroButton>
            <HeroButton
              href={twitterHashUrl}
              icon={<IconExternalLink size={22} />}
              size="lg"
              variant="filled"
              heightMultiplier={1}
              className="mmp-link-hero-button"
            >
              #でかプ / #VRでかプ (X投稿)
            </HeroButton>
          </div>
        </div>
      </section>

      <section className="link-card">
        <h2 className="section-title" style={{ color: sectionTitleColor }}>
          開発者向けリンク集 / Links for Developers
        </h2>
        <div className="link-grid link-grid-three">
          <HeroButton
            href="/swagger/index.html"
            icon={<IconExternalLink size={18} />}
            size="lg"
            variant="outline"
            heightMultiplier={1.1}
          >
            SwaggerUI (API一覧)
          </HeroButton>
          <HeroButton
            href="https://push.trap.show/?server=mariadb.ns-system.svc.cluster.local&username=nsapp_c27d6f571f88ffff360fe2&db=nsapp_c27d6f571f88ffff360fe2"
            icon={<IconLock size={18} />}
            size="lg"
            variant="outline"
            heightMultiplier={1.1}
          >
            データベース
          </HeroButton>
          <HeroButton
            href="https://github.com/pikachu0310/very-big-medal-pusher-data-server"
            icon={<IconBrandGithub size={18} />}
            size="lg"
            variant="outline"
            heightMultiplier={1.1}
          >
            Data Server GitHub
          </HeroButton>
        </div>
      </section>

      <div ref={deferredSectionAnchorRef} className="deferred-anchor" aria-hidden="true" />

      {showDeferredSections ? (
        <Suspense fallback={<DeferredPlaceholders />}>
          <DeferredSections primaryButtonColor={primaryButtonColor} />
        </Suspense>
      ) : (
        <DeferredPlaceholders />
      )}
    </div>
  );
}

export default HomePage;

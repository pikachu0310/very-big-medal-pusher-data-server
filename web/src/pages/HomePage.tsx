import { lazy, Suspense, useEffect, useState } from 'react';
import {
  Title,
  Text,
  Button,
  Stack,
  Card,
  Grid,
  Center,
  Box,
  Paper,
  Loader
} from '@mantine/core';
import {
  IconBrandDiscord,
  IconBrandGithub,
  IconWorld,
  IconBook2,
  IconExternalLink,
  IconLock
} from '@tabler/icons-react';

const StatsTabsSection = lazy(() => import('../components/StatsTabsSection'));
const DeveloperToolsSection = lazy(() => import('../components/DeveloperToolsSection'));

const linkSectionCardStyle = {
  backgroundColor: '#e9f3ff',
  border: '1px solid #cddff7'
} as const;

const sectionTitleColor = '#1f5da8';
const pageTitleColor = '#2c4256';
const primaryButtonColor = 'blue';

type HeroButtonProps = {
  children: React.ReactNode;
  href: string;
  icon: React.ReactNode;
  size?: 'lg' | 'xl';
  variant?: 'filled' | 'outline' | 'light';
  color?: string;
  heightMultiplier?: number;
  className?: string;
};

function HeroButton({
  children,
  href,
  icon,
  size = 'lg',
  variant = 'filled',
  color,
  heightMultiplier = 1,
  className
}: HeroButtonProps) {
  const textColor = variant === 'outline'
    ? color === 'black' ? '#1f2937' : undefined
    : variant === 'light'
      ? '#1e3a8a'
      : '#fff';
  const buttonClassName = variant === 'outline' && color === 'black'
    ? 'mmp-outline-black-button'
    : variant === 'filled' && color === primaryButtonColor
      ? 'mmp-primary-button'
      : undefined;
  const mergedClassName = [buttonClassName, className].filter(Boolean).join(' ');

  return (
    <Button
      component="a"
      href={href}
      target="_blank"
      rel="noreferrer"
      leftSection={icon}
      size={size}
      radius="md"
      fullWidth
      variant={variant}
      color={color}
      className={mergedClassName || undefined}
      fw={700}
      c={textColor}
      aria-label={typeof children === 'string' ? children : undefined}
      title={typeof children === 'string' ? children : undefined}
      style={{
        minHeight: `calc(${size === 'xl' ? 52 : 44}px * ${heightMultiplier})`,
        whiteSpace: 'normal',
        lineHeight: 1.3,
        paddingInline: 'clamp(0.75rem, 3vw, 1.4rem)',
        textAlign: 'center',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        wordBreak: 'break-word'
      }}
    >
      {children}
    </Button>
  );
}

function HomePage() {
  const [showDeferredSections, setShowDeferredSections] = useState(false);
  const twitterHashUrl = 'https://x.com/search?q=%23%E3%81%A7%E3%81%8B%E3%83%97%20OR%20%23VR%E3%81%A7%E3%81%8B%E3%83%97&f=live';

  useEffect(() => {
    const win = window as Window & {
      requestIdleCallback?: (callback: () => void, options?: { timeout: number }) => number;
      cancelIdleCallback?: (id: number) => void;
    };
    let timeoutId: number | undefined;
    let idleCallbackId: number | undefined;
    const revealDeferredSections = () => setShowDeferredSections(true);

    if (typeof win.requestIdleCallback === 'function') {
      idleCallbackId = win.requestIdleCallback(revealDeferredSections, { timeout: 500 });
    } else {
      timeoutId = window.setTimeout(revealDeferredSections, 350);
    }

    return () => {
      if (typeof timeoutId === 'number') {
        window.clearTimeout(timeoutId);
      }
      if (typeof idleCallbackId === 'number' && typeof win.cancelIdleCallback === 'function') {
        win.cancelIdleCallback(idleCallbackId);
      }
    };
  }, []);

  return (
    <Stack gap="xl" pt={0}>
      <Center mt={0} pt={0} mb={0}>
        <Box
          style={{
            width: 'min(100%, 700px)',
            aspectRatio: '1192 / 520',
            marginBottom: '0.5rem',
            marginTop: 0
          }}
        >
          <picture>
            <source
              srcSet="/MMP_logo_596.webp 596w, /MMP_logo_768.webp 768w, /MMP_logo_1192.webp 1192w"
              sizes="(max-width: 460px) calc(100vw - 2rem), (max-width: 768px) 85vw, 700px"
              type="image/webp"
            />
            <img
              src="/MMP_logo_768.webp"
              alt="Massive Medal Pusher ロゴ"
              width={1192}
              height={520}
              loading="eager"
              fetchPriority="high"
              decoding="sync"
              style={{
                width: '100%',
                height: '100%',
                display: 'block'
              }}
            />
          </picture>
        </Box>
      </Center>
      <Title order={1} ta="center" fz={{ base: 26, sm: 32 }} c={pageTitleColor}>
        クソでっけぇプッシャーゲーム 公式ウェブサイト
      </Title>
      <Text size="sm" c="#47678f" ta="center" mt={-10}>
        公式リンクや統計情報、開発者向けの情報をまとめて確認できます
      </Text>

      <Card padding="xl" radius="md" shadow="sm" style={linkSectionCardStyle}>
        <Stack gap="md">
          <Title order={2} fz="1.35rem" c={sectionTitleColor}>
            でかプ公式リンク集 / MMP Quick Links
          </Title>
          <Grid gutter={{ base: 'xs', md: 'md' }}>
            <Grid.Col span={{ base: 12, md: 6 }}>
              <HeroButton
                href="https://discord.com/invite/CgnYyXecKm"
                icon={<IconBrandDiscord size={36} />}
                size="xl"
                variant="filled"
                color={primaryButtonColor}
                heightMultiplier={2}
                className="mmp-link-hero-button"
              >
                公式Discord でかプ同好会
              </HeroButton>
              <Box mt="sm">
                <HeroButton
                  href="https://wikiwiki.jp/vr_bigpusher/"
                  icon={<IconBook2 size={24} />}
                  size="xl"
                  variant="filled"
                  color={primaryButtonColor}
                  className="mmp-link-hero-button"
                >
                  公式Wiki
                </HeroButton>
              </Box>
            </Grid.Col>
            <Grid.Col span={{ base: 12, md: 6 }}>
              <HeroButton
                href="https://vrchat.com/home/group/grp_5900a25d-0bb9-48d4-bab1-f3bd5c9a5e73"
                icon={<IconWorld size={22} />}
                size="lg"
                variant="filled"
                color={primaryButtonColor}
                className="mmp-link-hero-button"
              >
                公式グループ(VRChat)
              </HeroButton>
              <Box mt="sm">
                <HeroButton
                  href="https://vrchat.com/home/launch?worldId=wrld_1af53798-92a3-4c3f-99ae-a7c42ec6084d"
                  icon={<IconWorld size={22} />}
                  size="lg"
                  variant="filled"
                  color={primaryButtonColor}
                  className="mmp-link-hero-button"
                >
                  VRChatワールドリンク
                </HeroButton>
              </Box>
              <Box mt="sm">
                <HeroButton
                  href={twitterHashUrl}
                  icon={<IconExternalLink size={22} />}
                  size="lg"
                  variant="filled"
                  color={primaryButtonColor}
                  heightMultiplier={1}
                  className="mmp-link-hero-button"
                >
                  #でかプ / #VRでかプ (X投稿)
                </HeroButton>
              </Box>
            </Grid.Col>
          </Grid>
        </Stack>
      </Card>

      <Card padding="xl" radius="md" shadow="sm" style={linkSectionCardStyle}>
        <Stack gap="md">
          <Title order={2} fz="1.35rem" c={sectionTitleColor}>
            開発者向けリンク集 / Links for Developers
          </Title>
          <Grid gutter={{ base: 'xs', sm: 'sm' }}>
            <Grid.Col span={{ base: 12, sm: 4 }}>
              <HeroButton
                href="/swagger/index.html"
                icon={<IconExternalLink size={18} />}
                size="lg"
                variant="outline"
                color="black"
                heightMultiplier={1.1}
              >
                SwaggerUI (API一覧)
              </HeroButton>
            </Grid.Col>
            <Grid.Col span={{ base: 12, sm: 4 }}>
              <HeroButton
                href="https://push.trap.show/?server=mariadb.ns-system.svc.cluster.local&username=nsapp_c27d6f571f88ffff360fe2&db=nsapp_c27d6f571f88ffff360fe2"
                icon={<IconLock size={18} />}
                size="lg"
                variant="outline"
                color="black"
                heightMultiplier={1.1}
              >
                データベース
              </HeroButton>
            </Grid.Col>
            <Grid.Col span={{ base: 12, sm: 4 }}>
              <HeroButton
                href="https://github.com/pikachu0310/very-big-medal-pusher-data-server"
                icon={<IconBrandGithub size={18} />}
                size="lg"
                variant="outline"
                color="black"
                heightMultiplier={1.1}
              >
                Data Server GitHub
              </HeroButton>
            </Grid.Col>
          </Grid>
        </Stack>
      </Card>

      {showDeferredSections ? (
        <Suspense
          fallback={(
            <Paper p="xl" shadow="sm" radius="md" mih={360}>
              <Center role="status" aria-live="polite" mih={300}>
                <Loader />
              </Center>
            </Paper>
          )}
        >
          <StatsTabsSection primaryButtonColor={primaryButtonColor} />
        </Suspense>
      ) : (
        <Paper p="xl" shadow="sm" radius="md" mih={360}>
          <Center role="status" aria-live="polite" mih={300}>
            <Loader size="sm" />
          </Center>
        </Paper>
      )}

      {showDeferredSections ? (
        <Suspense
          fallback={(
            <Paper p="xl" shadow="sm" radius="md">
              <Center role="status" aria-live="polite" mih={120}>
                <Loader size="sm" />
              </Center>
            </Paper>
          )}
        >
          <DeveloperToolsSection />
        </Suspense>
      ) : (
        <Paper p="xl" shadow="sm" radius="md">
          <Center role="status" aria-live="polite" mih={120}>
            <Loader size="sm" />
          </Center>
        </Paper>
      )}
    </Stack>
  );
}

export default HomePage;

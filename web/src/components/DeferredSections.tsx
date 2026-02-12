import { MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css';
import DeveloperToolsSection from './DeveloperToolsSection';
import StatsTabsSection from './StatsTabsSection';

function DeferredSections({ primaryButtonColor }: { primaryButtonColor: string }) {
  return (
    <MantineProvider>
      <div className="deferred-sections">
        <StatsTabsSection primaryButtonColor={primaryButtonColor} />
        <DeveloperToolsSection />
      </div>
    </MantineProvider>
  );
}

export default DeferredSections;

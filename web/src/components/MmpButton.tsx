import type { MouseEventHandler, ReactNode } from 'react';

type MmpButtonSize = 'sm' | 'lg' | 'xl';
type MmpButtonVariant = 'primary' | 'outline';

type MmpButtonCommonProps = {
  children: ReactNode;
  icon?: ReactNode;
  size?: MmpButtonSize;
  variant?: MmpButtonVariant;
  className?: string;
  heightMultiplier?: number;
  loading?: boolean;
};

type MmpAnchorButtonProps = MmpButtonCommonProps & {
  href: string;
  target?: '_blank' | '_self' | '_parent' | '_top';
  rel?: string;
  onClick?: never;
  disabled?: never;
  type?: never;
};

type MmpActionButtonProps = MmpButtonCommonProps & {
  href?: never;
  onClick?: MouseEventHandler<HTMLButtonElement>;
  disabled?: boolean;
  type?: 'button' | 'submit' | 'reset';
};

export type MmpButtonProps = MmpAnchorButtonProps | MmpActionButtonProps;

function buildButtonClassName({
  size,
  variant,
  className
}: {
  size: MmpButtonSize;
  variant: MmpButtonVariant;
  className?: string;
}) {
  return [
    'hero-link-button',
    size === 'xl'
      ? 'hero-link-button-xl'
      : size === 'sm'
        ? 'hero-link-button-sm'
        : 'hero-link-button-lg',
    variant === 'outline' ? 'hero-link-button-outline' : 'hero-link-button-filled',
    className
  ].filter(Boolean).join(' ');
}

function renderButtonContent({
  icon,
  loading,
  children
}: {
  icon?: ReactNode;
  loading?: boolean;
  children: ReactNode;
}) {
  return (
    <span className="hero-link-button-content">
      {loading ? (
        <span className="hero-link-button-icon" aria-hidden="true">
          <span className="hero-link-button-spinner" />
        </span>
      ) : icon ? (
        <span className="hero-link-button-icon" aria-hidden="true">{icon}</span>
      ) : null}
      <span className="hero-link-button-label">{children}</span>
    </span>
  );
}

function isAnchorButtonProps(props: MmpButtonProps): props is MmpAnchorButtonProps {
  return typeof (props as MmpAnchorButtonProps).href === 'string';
}

function MmpButton(props: MmpButtonProps) {
  const {
    children,
    icon,
    size = 'lg',
    variant = 'primary',
    className,
    heightMultiplier = 1,
    loading = false
  } = props;
  const mergedClassName = buildButtonClassName({ size, variant, className });
  const style = { minHeight: `calc(${size === 'xl' ? 52 : size === 'sm' ? 36 : 44}px * ${heightMultiplier})` };
  const content = renderButtonContent({ icon, loading, children });

  if (isAnchorButtonProps(props)) {
    const rel = props.target === '_blank'
      ? props.rel ?? 'noreferrer'
      : props.rel;

    return (
      <a
        href={props.href}
        target={props.target}
        rel={rel}
        className={mergedClassName}
        style={style}
      >
        {content}
      </a>
    );
  }

  return (
    <button
      type={props.type ?? 'button'}
      onClick={props.onClick}
      className={mergedClassName}
      style={style}
      disabled={props.disabled || loading}
      aria-busy={loading || undefined}
    >
      {content}
    </button>
  );
}

type MmpPrimaryButtonProps = MmpAnchorButtonProps | MmpActionButtonProps;
type MmpOutlineButtonProps = MmpAnchorButtonProps | MmpActionButtonProps;

export function MmpPrimaryButton(props: MmpPrimaryButtonProps) {
  return <MmpButton {...props} variant="primary" />;
}

export function MmpOutlineButton(props: MmpOutlineButtonProps) {
  return <MmpButton {...props} variant="outline" />;
}

export default MmpButton;

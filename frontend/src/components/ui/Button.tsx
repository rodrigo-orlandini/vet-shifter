import { type ButtonHTMLAttributes, type ReactNode } from "react";
import Link, { type LinkProps } from "next/link";
import { SpinnerIcon } from "@/components/icons/SpinnerIcon";

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "secondary" | "ghost" | "back" | "link";
  loading?: boolean;
  children: ReactNode;
  className?: string;
}

const base =
  "inline-flex items-center justify-center gap-2 rounded-lg px-5 h-12 text-[15px] font-semibold transition-colors hover:cursor-pointer disabled:cursor-not-allowed disabled:opacity-50";

const variants = {
  primary:
    "bg-primary text-surface hover:bg-primary-hover focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/40",
  secondary:
    "border border-edge-input bg-surface text-ink-body hover:bg-page",
  ghost:
    "border-[1.5px] border-primary bg-transparent text-primary hover:bg-primary/5 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/30",
  back:
    "border-[1.5px] border-primary bg-transparent text-primary hover:bg-primary/5 sm:border-edge sm:bg-surface sm:text-ink-muted sm:hover:bg-page focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/30",
  link: "",
};

const linkClass =
  "inline-flex items-center gap-2 text-sm font-medium text-primary underline-offset-2 transition-colors hover:underline focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/30 disabled:cursor-not-allowed disabled:opacity-50 hover:cursor-pointer";

function Spinner() {
  return <SpinnerIcon className="h-4 w-4 animate-spin" />;
}

export interface ButtonLinkProps extends LinkProps {
  variant?: "primary" | "secondary" | "ghost" | "back";
  children: ReactNode;
  className?: string;
}

export function ButtonLink({
  variant = "primary",
  children,
  className = "",
  ...rest
}: ButtonLinkProps) {
  return (
    <Link className={`${base} ${variants[variant]} ${className}`} {...rest}>
      {children}
    </Link>
  );
}


export function Button({
  variant = "primary",
  loading = false,
  disabled,
  children,
  className = "",
  type = "button",
  ...rest
}: ButtonProps) {
  const resolvedClass =
    variant === "link"
      ? `${linkClass} ${className}`
      : `${base} ${variants[variant]} ${className}`;

  return (
    <button
      type={type}
      disabled={disabled || loading}
      className={resolvedClass}
      {...rest}
    >
      {loading && <Spinner />}
      {children}
    </button>
  );
}

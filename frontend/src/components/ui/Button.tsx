import { type ButtonHTMLAttributes, type ReactNode } from "react";
import { SpinnerIcon } from "@/components/icons/SpinnerIcon";

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "secondary" | "ghost" | "back";
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
};

function Spinner() {
  return <SpinnerIcon className="h-4 w-4 animate-spin" />;
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
  return (
    <button
      type={type}
      disabled={disabled || loading}
      className={`${base} ${variants[variant]} ${className}`}
      {...rest}
    >
      {loading && <Spinner />}
      {children}
    </button>
  );
}

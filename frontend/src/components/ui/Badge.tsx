import { type ReactNode } from "react";

export type BadgeVariant = "warning" | "success" | "danger" | "info" | "default";

export interface BadgeProps {
  variant?: BadgeVariant;
  children: ReactNode;
  className?: string;
}

const variantClasses: Record<BadgeVariant, string> = {
  warning: "bg-warning-subtle text-warning-ink",
  success: "bg-success-subtle text-success",
  danger: "bg-danger/10 text-danger",
  info: "bg-info-subtle text-info",
  default: "bg-edge text-ink-muted",
};

export function Badge({ variant = "default", children, className = "" }: BadgeProps) {
  return (
    <span
      className={`rounded px-2 py-0.5 text-[11px] font-medium ${variantClasses[variant]} ${className}`}
    >
      {children}
    </span>
  );
}

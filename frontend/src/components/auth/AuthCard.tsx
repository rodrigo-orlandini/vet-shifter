import type { ReactNode } from "react";

export function AuthCard({
  children,
  className = "",
}: Readonly<{
  children: ReactNode;
  className?: string;
}>) {
  return (
    <div
      className={`rounded-xl bg-surface shadow-[0_2px_12px_rgba(0,0,0,0.05)] ${className}`}
    >
      {children}
    </div>
  );
}

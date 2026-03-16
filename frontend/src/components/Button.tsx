"use client";

import { type ButtonHTMLAttributes, type ReactNode } from "react";

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "secondary";
  loading?: boolean;
  children: ReactNode;
  className?: string;
}

const base =
  "rounded-lg px-4 py-2 text-sm font-medium transition-colors disabled:opacity-70 disabled:pointer-events-none";

const variants = {
  primary: "bg-emerald-600 text-white hover:bg-emerald-700",
  secondary: "border border-neutral-300 bg-white text-neutral-700 hover:bg-neutral-50",
};

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
      disabled={disabled ?? loading}
      className={`${base} ${variants[variant]} ${className}`}
      {...rest}
    >
      {children}
    </button>
  );
}

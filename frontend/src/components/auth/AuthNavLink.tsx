"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

export function AuthNavLink() {
  const isLogin = usePathname() === "/login";
  const linkClass =
    "text-sm font-medium text-primary underline-offset-2 transition-colors hover:text-primary-hover hover:underline focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/30";

  return isLogin ? (
    <Link href="/signup/company" className={linkClass}>
      Cadastrar
    </Link>
  ) : (
    <Link href="/login" className={linkClass}>
      Entrar
    </Link>
  );
}

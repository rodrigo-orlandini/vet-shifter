"use client";

import { usePathname } from "next/navigation";
import { ButtonLink } from "@/components/ui/Button";

export function AuthNavLink() {
  const isLogin = usePathname() === "/login";

  return isLogin ? (
    <ButtonLink href="/signup/company" variant="ghost">
      Cadastrar
    </ButtonLink>
  ) : (
    <ButtonLink href="/login" variant="ghost">
      Entrar
    </ButtonLink>
  );
}

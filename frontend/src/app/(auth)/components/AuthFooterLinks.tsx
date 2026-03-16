"use client";

import Link from "next/link";

export type AuthFooterVariant = "login" | "company" | "veterinary";

const linkClass = "font-medium text-emerald-600 hover:underline";

export function AuthFooterLinks({ variant }: { variant: AuthFooterVariant }) {
  if (variant === "login") {
    return (
      <p className="mt-6 text-center text-sm text-neutral-600">
        Não tem uma conta?{" "}
        <Link href="/signup/company" className={linkClass}>
          Cadastre sua empresa
        </Link>
        {" ou "}
        <Link href="/signup/veterinary" className={linkClass}>
          Cadastre-se como veterinário
        </Link>
      </p>
    );
  }

  if (variant === "company") {
    return (
      <p className="mt-6 text-center text-sm text-neutral-600">
        Já tem uma conta?{" "}
        <Link href="/login" className={linkClass}>
          Entrar
        </Link>
        {" · "}
        É veterinário?{" "}
        <Link href="/signup/veterinary" className={linkClass}>
          Cadastre-se aqui
        </Link>
      </p>
    );
  }

  return (
    <p className="mt-6 text-center text-sm text-neutral-600">
      Já tem uma conta?{" "}
      <Link href="/login" className={linkClass}>
        Entrar
      </Link>
      {" · "}
      É empresa?{" "}
      <Link href="/signup/company" className={linkClass}>
        Cadastre sua clínica
      </Link>
    </p>
  );
}

import Link from "next/link";

export type AuthFooterVariant = "login" | "company" | "veterinary";

const linkClass =
  "font-medium text-primary underline-offset-2 transition-colors hover:text-primary-hover hover:underline focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/30";

export function AuthFooterLinks({ variant }: { variant: AuthFooterVariant }) {
  if (variant === "login") {
    return (
      <p className="mt-6 text-center text-sm text-ink-muted">
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
      <p className="mt-6 text-center text-sm text-ink-muted">
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
    <p className="mt-6 text-center text-sm text-ink-muted">
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

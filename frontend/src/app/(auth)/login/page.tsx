"use client";

import { Suspense, useState } from "react";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { getVetShifterAPI } from "@/api/generated/api";
import { useToast } from "@/components/toast/ToastProvider";
import { FieldWithError } from "@/components/FieldWithError";
import { Button } from "@/components/Button";
import { AuthFooterLinks } from "../components/AuthFooterLinks";
import { isValidEmail, validationMessages } from "@/lib/validation";
import { getBackendErrorMessage } from "@/lib/backendErrorMessage";

const api = getVetShifterAPI();

function LoginForm() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const registered = searchParams.get("registered");
  const resetSuccess = searchParams.get("reset") === "success";
  const { pushToast } = useToast();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [rememberMe, setRememberMe] = useState(false);
  const [fieldErrors, setFieldErrors] = useState<{ email?: string; password?: string }>({});
  const [error, setError] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    const err: { email?: string; password?: string } = {};
    if (!email.trim()) err.email = validationMessages.required;
    else if (!isValidEmail(email)) err.email = validationMessages.email;
    if (!password.trim()) err.password = validationMessages.required;
    if (Object.keys(err).length > 0) {
      setFieldErrors(err);
      return;
    }
    setFieldErrors({});

    setSubmitting(true);

    try {
      const userType = await api.getAuthUserType({ email });

      if (!userType?.user_type) {
        const message = "E-mail ou senha incorretos.";
        pushToast({ tone: "error", message });
        setError(message);
        setSubmitting(false);
        return;
      }

      const credentials = { email, password, remember_me: rememberMe };

      if (userType.user_type === "company_owner") {
        const res = await api.postAuthLoginOwner(credentials);
        if (res) {
          router.push("/dashboard/company");
          return;
        }
      }

      if (userType.user_type === "shift_veterinary") {
        const res = await api.postAuthLoginVeterinary(credentials);
        if (res) {
          router.push("/dashboard/veterinary");
          return;
        }
      }

      setError("E-mail ou senha inválidos.");
    } catch (err) {
      const message = getBackendErrorMessage(err);
      pushToast({ tone: "error", message });
      setError(message);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="overflow-hidden rounded-2xl border border-neutral-200/80 bg-white shadow-xl shadow-neutral-200/50">
      <div className="border-t-4 border-emerald-500 bg-linear-to-r from-emerald-500/5 to-teal-500/5 px-8 pt-8 pb-6">
        <h1 className="mb-1 text-2xl font-bold tracking-tight text-neutral-900">Entrar</h1>
        <p className="text-sm text-neutral-600">
          Informe seus dados para acessar sua conta.
        </p>
      </div>
      <div className="p-8 pt-6">

      {registered === "company" && (
        <p className="mb-4 rounded-lg bg-emerald-50 p-3 text-sm text-emerald-800">
          Conta da empresa criada. Agora você pode entrar.
        </p>
      )}
      {registered === "veterinary" && (
        <p className="mb-4 rounded-lg bg-emerald-50 p-3 text-sm text-emerald-800">
          Conta de veterinário criada. Agora você pode entrar.
        </p>
      )}
      {resetSuccess && (
        <p className="mb-4 rounded-lg bg-emerald-50 p-3 text-sm text-emerald-800">
          Senha alterada. Agora você pode entrar com a nova senha.
        </p>
      )}

      <form onSubmit={handleSubmit} className="flex flex-col gap-4">
        <FieldWithError
          label="E-mail"
          error={fieldErrors.email}
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="seu@email.com"
        />
        <FieldWithError
          label="Senha"
          error={fieldErrors.password}
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="••••••••"
        />
        <label className="flex items-center gap-2">
          <input
            type="checkbox"
            checked={rememberMe}
            onChange={(e) => setRememberMe(e.target.checked)}
            className="rounded border-neutral-300"
          />
          <span className="text-sm text-neutral-700">Lembrar de mim</span>
        </label>

        {error && (
          <p className="text-sm text-red-600" role="alert">
            {error}
          </p>
        )}

        <Button type="submit" disabled={submitting} loading={submitting}>
          {submitting ? "Entrando…" : "Entrar"}
        </Button>
      </form>

      <p className="mt-4 text-center text-sm text-neutral-600">
        <Link href="/forgot-password" className="font-medium text-emerald-600 hover:underline">
          Esqueceu a senha?
        </Link>
      </p>

      <AuthFooterLinks variant="login" />
      </div>
    </div>
  );
}

export default function LoginPage() {
  return (
    <Suspense fallback={<div className="mx-auto max-w-md p-6 text-center text-neutral-500">Carregando…</div>}>
      <LoginForm />
    </Suspense>
  );
}

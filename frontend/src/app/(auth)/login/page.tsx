"use client";

import { Suspense, useState } from "react";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { getVetShifterAPI } from "@/api/generated/api";
import { useToast } from "@/components/toast/ToastProvider";
import { FieldWithError } from "@/components/FieldWithError";
import { Button } from "@/components/Button";
import { AuthFooterLinks } from "@/components/auth/AuthFooterLinks";
import { AuthCard } from "@/components/auth/AuthCard";
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
    <AuthCard>
      <div className="p-5 sm:p-10">
        <h1 className="text-xl font-bold text-[#18181B] sm:text-2xl">Entrar</h1>
        <p className="mt-1 text-sm text-[#6C757D]">Informe seus dados para acessar sua conta.</p>
        <div className="mt-6">
        {registered === "company" && (
          <p className="mb-4 rounded-lg bg-[#E8F4FD] p-3 text-sm text-[#2B6CB0]">
            Conta da clínica criada. Agora você pode entrar.
          </p>
        )}
        {registered === "veterinary" && (
          <p className="mb-4 rounded-lg bg-[#E8F4FD] p-3 text-sm text-[#2B6CB0]">
            Cadastro de veterinário concluído. Agora você pode entrar.
          </p>
        )}
        {resetSuccess && (
          <p className="mb-4 rounded-lg bg-[#E8F4FD] p-3 text-sm text-[#2B6CB0]">
            Senha alterada. Agora você pode entrar com a nova senha.
          </p>
        )}

        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
          <FieldWithError
            label="E-mail"
            requiredMark
            error={fieldErrors.email}
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="seu@email.com"
          />
          <FieldWithError
            label="Senha"
            requiredMark
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
              className="h-4 w-4 rounded border-[#DEE2E6] text-[#2A9D8F] focus:ring-[#2A9D8F]"
            />
            <span className="text-sm text-[#6C757D]">Lembrar de mim</span>
          </label>

          {error && (
            <p className="text-sm text-[#E53E3E]" role="alert">
              {error}
            </p>
          )}

          <Button type="submit" disabled={submitting} loading={submitting} className="w-full">
            {submitting ? "Entrando…" : "Entrar"}
          </Button>
        </form>

        <p className="mt-4 text-center text-sm text-[#6C757D]">
          <Link href="/forgot-password" className="font-medium text-[#2A9D8F] hover:underline">
            Esqueceu a senha?
          </Link>
        </p>

        <AuthFooterLinks variant="login" />
        </div>
      </div>
    </AuthCard>
  );
}

export default function LoginPage() {
  return (
    <Suspense fallback={<div className="p-6 text-center text-[#6C757D]">Carregando…</div>}>
      <LoginForm />
    </Suspense>
  );
}

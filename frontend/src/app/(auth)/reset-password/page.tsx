"use client";

import { useState, Suspense } from "react";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { getVetShifterAPI } from "@/api/generated/api";
import { useToast } from "@/components/toast/ToastProvider";
import { FieldWithError } from "@/components/FieldWithError";
import { Button } from "@/components/Button";
import { AuthCard } from "@/components/auth/AuthCard";
import { validationMessages } from "@/lib/validation";
import { getBackendErrorMessage } from "@/lib/backendErrorMessage";

const api = getVetShifterAPI();

function ResetPasswordForm() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const token = searchParams.get("token") ?? "";

  const { pushToast } = useToast();
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [fieldErrors, setFieldErrors] = useState<{ password?: string; confirmPassword?: string }>({});
  const [error, setError] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    const err: { password?: string; confirmPassword?: string } = {};

    if (password.length < 8) err.password = validationMessages.password;
    if (password !== confirmPassword) err.confirmPassword = validationMessages.passwordMatch;

    if (Object.keys(err).length > 0) {
      setFieldErrors(err);
      return;
    }

    if (!token) {
      setError("Link inválido. Solicite uma nova redefinição de senha.");
      return;
    }

    setFieldErrors({});
    setSubmitting(true);

    try {
      await api.postAuthResetPassword({ token, new_password: password });
      router.push("/login?reset=success");
    } catch (e) {
      const message = getBackendErrorMessage(e);
      pushToast({ tone: "error", message });
      setError(message);
    } finally {
      setSubmitting(false);
    }
  };

  if (!token) {
    return (
      <AuthCard>
        <div className="p-5 sm:p-10">
          <h1 className="text-xl font-bold text-[#18181B] sm:text-2xl">Link inválido</h1>
          <p className="mt-1 text-sm text-[#6C757D]">
            Este link de redefinição é inválido ou está incompleto. Solicite um novo link.
          </p>
          <div className="mt-6">
            <Link
              href="/forgot-password"
              className="inline-flex h-12 w-full items-center justify-center rounded-lg bg-[#2A9D8F] px-5 text-[15px] font-semibold text-white hover:bg-primary-hover"
            >
              Solicitar novo link
            </Link>
          </div>
        </div>
      </AuthCard>
    );
  }

  return (
    <AuthCard>
      <div className="p-5 sm:p-10">
        <h1 className="text-xl font-bold text-[#18181B] sm:text-2xl">Redefinir senha</h1>
        <p className="mt-1 text-sm text-[#6C757D]">Informe sua nova senha abaixo.</p>

        <form onSubmit={handleSubmit} className="mt-6 flex flex-col gap-4">
          <FieldWithError
            label="Nova senha"
            requiredMark
            error={fieldErrors.password}
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            minLength={8}
            placeholder="Mínimo 8 caracteres"
          />
          <FieldWithError
            label="Confirmar senha"
            requiredMark
            error={fieldErrors.confirmPassword}
            type="password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            minLength={8}
            placeholder="Repita a senha"
          />

          {error && (
            <p className="text-sm text-[#E53E3E]" role="alert">
              {error}
            </p>
          )}

          <Button type="submit" disabled={submitting} loading={submitting} className="w-full">
            {submitting ? "Atualizando…" : "Atualizar senha"}
          </Button>
        </form>

        <p className="mt-6 text-center text-sm text-[#6C757D]">
          <Link href="/login" className="font-medium text-[#2A9D8F] hover:underline">
            Voltar para entrar
          </Link>
        </p>
      </div>
    </AuthCard>
  );
}

export default function ResetPasswordPage() {
  return (
    <Suspense fallback={<div className="p-6 text-center text-[#6C757D]">Carregando…</div>}>
      <ResetPasswordForm />
    </Suspense>
  );
}

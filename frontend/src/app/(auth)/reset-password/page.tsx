"use client";

import { useState, Suspense } from "react";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { AuthenticationService } from "@/auth/api";
import { useToast } from "@/components/toast/ToastProvider";
import { FieldWithError } from "@/components/FieldWithError";
import { Button } from "@/components/Button";
import { validationMessages } from "@/lib/validation";
import { getBackendErrorMessage } from "@/lib/backendErrorMessage";

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
      await AuthenticationService.resetPassword({ token, new_password: password });
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
      <div className="overflow-hidden rounded-2xl border border-neutral-200/80 bg-white shadow-xl shadow-neutral-200/50">
        <div className="border-t-4 border-amber-500 bg-linear-to-r from-amber-500/5 to-orange-500/5 px-8 pt-8 pb-6">
          <h1 className="mb-1 text-2xl font-bold tracking-tight text-neutral-900">Link inválido</h1>
          <p className="text-sm text-neutral-600">
            Este link de redefinição é inválido ou está incompleto. Solicite um novo link.
          </p>
        </div>
        <div className="p-8 pt-6">
          <Link
            href="/forgot-password"
            className="inline-block rounded-lg bg-emerald-600 px-4 py-2 text-sm font-medium text-white hover:bg-emerald-700"
          >
            Solicitar novo link
          </Link>
        </div>
      </div>
    );
  }

  return (
    <div className="overflow-hidden rounded-2xl border border-neutral-200/80 bg-white shadow-xl shadow-neutral-200/50">
      <div className="border-t-4 border-emerald-500 bg-linear-to-r from-emerald-500/5 to-teal-500/5 px-8 pt-8 pb-4">
        <h1 className="mb-1 text-2xl font-bold tracking-tight text-neutral-900">Redefinir senha</h1>
        <p className="text-sm text-neutral-600">
          Informe sua nova senha abaixo.
        </p>
      </div>
      <div className="p-8 pt-6">
        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
          <FieldWithError
            label="Nova senha"
            error={fieldErrors.password}
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            minLength={8}
            placeholder="Mínimo 8 caracteres"
          />
          <FieldWithError
            label="Confirmar senha"
            error={fieldErrors.confirmPassword}
            type="password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            minLength={8}
            placeholder="Repita a senha"
          />

          {error && (
            <p className="text-sm text-red-600" role="alert">
              {error}
            </p>
          )}

          <Button type="submit" disabled={submitting} loading={submitting}>
            {submitting ? "Atualizando…" : "Atualizar senha"}
          </Button>
        </form>

        <p className="mt-6 text-center text-sm text-neutral-600">
          <Link href="/login" className="font-medium text-emerald-600 hover:underline">
            Voltar para entrar
          </Link>
        </p>
      </div>
    </div>
  );
}

export default function ResetPasswordPage() {
  return (
    <Suspense fallback={<div className="mx-auto max-w-md p-6 text-center text-neutral-500">Carregando…</div>}>
      <ResetPasswordForm />
    </Suspense>
  );
}

"use client";

import { useState } from "react";
import Link from "next/link";
import { getVetShifterAPI } from "@/api/generated/api";
import { useToast } from "@/components/toast/ToastProvider";
import { FieldWithError } from "@/components/ui/FieldWithError";
import { Button } from "@/components/ui/Button";
import { AuthCard } from "@/components/auth/AuthCard";
import { isRequired, isValidEmail, validationMessages } from "@/lib/validation";
import { getBackendErrorMessage } from "@/lib/backendErrorMessage";

const api = getVetShifterAPI();

export default function ForgotPasswordPage() {
  const { pushToast } = useToast();
  const [email, setEmail] = useState("");
  const [submitted, setSubmitted] = useState(false);
  const [emailError, setEmailError] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    if (!isRequired(email)) {
      setEmailError(validationMessages.required);
      return;
    }

    if (!isValidEmail(email)) {
      setEmailError(validationMessages.email);
      return;
    }

    setEmailError(null);
    setSubmitting(true);

    try {
      await api.postAuthForgotPassword({ email });
      setSubmitted(true);
    } catch (e) {
      const message = getBackendErrorMessage(e);
      pushToast({ tone: "error", message });
      setError(message);
    } finally {
      setSubmitting(false);
    }
  };

  if (submitted) {
    return (
      <AuthCard>
        <div className="p-5 sm:p-10">
          <h1 className="text-xl font-bold text-[#18181B] sm:text-2xl">Verifique seu e-mail</h1>
          <p className="mt-1 text-sm text-[#6C757D]">
            Se existir uma conta com este e-mail, enviamos as instruções para redefinir sua senha.
          </p>
          <div className="mt-6">
            <Link
              href="/login"
              className="inline-flex h-12 w-full items-center justify-center rounded-lg bg-[#2A9D8F] px-5 text-[15px] font-semibold text-white hover:bg-primary-hover"
            >
              Voltar para entrar
            </Link>
          </div>
        </div>
      </AuthCard>
    );
  }

  return (
    <AuthCard>
      <div className="p-5 sm:p-10">
        <h1 className="text-xl font-bold text-[#18181B] sm:text-2xl">Esqueci a senha</h1>
        <p className="mt-1 text-sm text-[#6C757D]">
          Informe seu e-mail e enviaremos um link para redefinir sua senha.
        </p>

        <form onSubmit={handleSubmit} className="mt-6 flex flex-col gap-4">
          <FieldWithError
            label="E-mail"
            requiredMark
            error={emailError}
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="seu@email.com"
          />

          {error && (
            <p className="text-sm text-[#E53E3E]" role="alert">
              {error}
            </p>
          )}

          <Button type="submit" disabled={submitting} loading={submitting} className="w-full">
            {submitting ? "Enviando…" : "Enviar link"}
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

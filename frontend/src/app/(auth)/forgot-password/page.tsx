"use client";

import { useState } from "react";
import Link from "next/link";
import { getVetShifterAPI } from "@/api/generated/api";
import { useToast } from "@/components/toast/ToastProvider";
import { FieldWithError } from "@/components/FieldWithError";
import { Button } from "@/components/Button";
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
      <div className="overflow-hidden rounded-2xl border border-neutral-200/80 bg-white shadow-xl shadow-neutral-200/50">
        <div className="border-t-4 border-emerald-500 bg-linear-to-r from-emerald-500/5 to-teal-500/5 px-8 pt-8 pb-6">
          <h1 className="mb-1 text-2xl font-bold tracking-tight text-neutral-900">Verifique seu e-mail</h1>
          <p className="text-sm text-neutral-600">
            Se existir uma conta com este e-mail, enviamos as instruções para redefinir sua senha.
          </p>
        </div>
        <div className="p-8 pt-6">
          <Link
            href="/login"
            className="inline-block rounded-lg bg-emerald-600 px-4 py-2 text-sm font-medium text-white hover:bg-emerald-700"
          >
            Voltar para entrar
          </Link>
        </div>
      </div>
    );
  }

  return (
    <div className="overflow-hidden rounded-2xl border border-neutral-200/80 bg-white shadow-xl shadow-neutral-200/50">
      <div className="border-t-4 border-emerald-500 bg-linear-to-r from-emerald-500/5 to-teal-500/5 px-8 pt-8 pb-4">
        <h1 className="mb-1 text-2xl font-bold tracking-tight text-neutral-900">Esqueci a senha</h1>
        <p className="text-sm text-neutral-600">
          Informe seu e-mail e enviaremos um link para redefinir sua senha.
        </p>
      </div>
      <div className="p-8 pt-6">
        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
          <FieldWithError
            label="E-mail"
            error={emailError}
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="seu@email.com"
          />

          {error && (
            <p className="text-sm text-red-600" role="alert">
              {error}
            </p>
          )}

          <Button type="submit" disabled={submitting} loading={submitting}>
            {submitting ? "Enviando…" : "Enviar link"}
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

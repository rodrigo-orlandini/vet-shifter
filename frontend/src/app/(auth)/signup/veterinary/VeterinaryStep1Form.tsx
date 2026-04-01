"use client";

import Link from "next/link";
import { FieldWithError } from "@/components/ui/FieldWithError";
import { PasswordFields } from "@/components/auth/PasswordFields";
import { Checkbox } from "@/components/ui/Checkbox";
import { formatCpf, formatPhoneBr } from "@/lib/masks";
import type { ControllersRegisterShiftVeterinaryRequest } from "@/api/generated/api";

type FieldErrors = Partial<Record<keyof ControllersRegisterShiftVeterinaryRequest | "specialties" | "confirmPassword", string>>;

export interface VeterinaryStep1FormProps {
  form: ControllersRegisterShiftVeterinaryRequest;
  fieldErrors: FieldErrors;
  update: (partial: Partial<ControllersRegisterShiftVeterinaryRequest>) => void;
  confirmPassword: string;
  onConfirmPasswordChange: (value: string) => void;
}

export function VeterinaryStep1Form({
  form,
  fieldErrors,
  update,
  confirmPassword,
  onConfirmPasswordChange,
}: VeterinaryStep1FormProps) {
  return (
    <div className="flex flex-col gap-4">
      <FieldWithError
        label="Nome completo"
        requiredMark
        error={fieldErrors.full_name}
        value={form.full_name}
        onChange={(e) => update({ full_name: e.target.value })}
        placeholder="Seu nome completo"
      />
      <FieldWithError
        label="CPF"
        requiredMark
        error={fieldErrors.cpf}
        value={form.cpf}
        onChange={(e) => update({ cpf: formatCpf(e.target.value) })}
        placeholder="000.000.000-00"
      />
      <FieldWithError
        label="E-mail"
        requiredMark
        error={fieldErrors.email}
        type="email"
        value={form.email}
        onChange={(e) => update({ email: e.target.value })}
        placeholder="seu@email.com"
      />
      <FieldWithError
        label="Telefone"
        requiredMark
        error={fieldErrors.phone}
        value={form.phone}
        onChange={(e) => update({ phone: formatPhoneBr(e.target.value) })}
        placeholder="(11) 99999-9999"
      />

      <PasswordFields
        password={form.password}
        confirmPassword={confirmPassword}
        onPasswordChange={(v) => update({ password: v })}
        onConfirmChange={onConfirmPasswordChange}
        passwordError={fieldErrors.password}
        confirmError={fieldErrors.confirmPassword}
      />

      <div className="flex flex-col gap-2">
        {fieldErrors.consent_lgpd && (
          <p className="text-sm text-danger" role="alert">
            {fieldErrors.consent_lgpd}
          </p>
        )}
        <label className="flex items-start gap-3">
          <Checkbox
            checked={form.consent_lgpd}
            onChange={(e) => update({ consent_lgpd: e.target.checked })}
            checkboxClassName="mt-0.5"
          />
          <span className="text-[13px] leading-relaxed text-ink-muted">
            Li e concordo com os{" "}
            <Link href="/terms" className="font-medium text-primary underline">
              Termos de Uso
            </Link>{" "}
            e a{" "}
            <Link href="/privacy" className="font-medium text-primary underline">
              Política de Privacidade
            </Link>
            , e autorizo o tratamento dos meus dados conforme a LGPD.
          </span>
        </label>
      </div>
    </div>
  );
}

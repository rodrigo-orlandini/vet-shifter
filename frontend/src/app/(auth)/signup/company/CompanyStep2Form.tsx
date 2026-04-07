"use client";

import Link from "next/link";
import { FieldWithError } from "@/components/ui/FieldWithError";
import { PasswordFields } from "@/components/auth/PasswordFields";
import { Checkbox } from "@/components/ui/Checkbox";
import { formatPhoneBr } from "@/lib/masks";
import type { ControllersRegisterCompanyRequest } from "@/api/generated/api";

type FieldErrors = Partial<Record<keyof ControllersRegisterCompanyRequest, string>> & {
  confirmPassword?: string;
};

export interface CompanyStep2FormProps {
  form: ControllersRegisterCompanyRequest;
  fieldErrors: FieldErrors;
  update: (partial: Partial<ControllersRegisterCompanyRequest>) => void;
  confirmPassword: string;
  onConfirmPasswordChange: (value: string) => void;
}

export function CompanyStep2Form({
  form,
  fieldErrors,
  update,
  confirmPassword,
  onConfirmPasswordChange,
}: CompanyStep2FormProps) {
  return (
    <div className="flex flex-col gap-4">
      <FieldWithError
        label="Nome do responsável"
        requiredMark
        error={fieldErrors.owner_name}
        value={form.owner_name}
        onChange={(e) => update({ owner_name: e.target.value })}
        placeholder="Ex: Maria Silva"
      />

      <FieldWithError
        label="E-mail"
        requiredMark
        error={fieldErrors.email}
        type="email"
        value={form.email}
        onChange={(e) => update({ email: e.target.value })}
        placeholder="nome@clinica.com.br"
      />

      <FieldWithError
        label="Telefone"
        requiredMark
        error={fieldErrors.phone}
        value={form.phone}
        onChange={(e) => update({ phone: formatPhoneBr(e.target.value) })}
        placeholder="(00) 00000-0000"
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

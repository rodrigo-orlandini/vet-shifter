"use client";

import { FieldWithError } from "@/components/FieldWithError";
import type { RegisterVeterinaryRequest } from "@/auth/types/veterinary-signup";

type FieldErrors = Partial<Record<keyof RegisterVeterinaryRequest | "specialties", string>>;

export interface VeterinaryStep3FormProps {
  form: RegisterVeterinaryRequest;
  fieldErrors: FieldErrors;
  update: (partial: Partial<RegisterVeterinaryRequest>) => void;
}

export function VeterinaryStep3Form({ form, fieldErrors, update }: VeterinaryStep3FormProps) {
  return (
    <div className="flex flex-col gap-4">
      <FieldWithError
        label="Senha"
        error={fieldErrors.password}
        type="password"
        value={form.password}
        onChange={(e) => update({ password: e.target.value })}
        placeholder="Mínimo 8 caracteres"
      />
      <label className="flex flex-col gap-1">
        {fieldErrors.consent_lgpd && (
          <span className="text-xs text-red-600" role="alert">
            {fieldErrors.consent_lgpd}
          </span>
        )}
        <label className="flex items-start gap-2">
          <input
            type="checkbox"
            checked={form.consent_lgpd}
            onChange={(e) => update({ consent_lgpd: e.target.checked })}
            className={`mt-1 rounded ${fieldErrors.consent_lgpd ? "border-red-500" : "border-neutral-300"}`}
          />
          <span className="text-sm text-neutral-700">
            Concordo com o tratamento dos meus dados de acordo com a LGPD (Lei Geral de Proteção
            de Dados).
          </span>
        </label>
      </label>
    </div>
  );
}

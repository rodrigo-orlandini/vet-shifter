"use client";

import { FieldWithError } from "@/components/FieldWithError";
import { formatPhoneBr } from "@/lib/masks";
import type { ControllersRegisterCompanyRequest } from "@/api/generated/api";

type FieldErrors = Partial<Record<keyof ControllersRegisterCompanyRequest, string>>;

export interface CompanyStep2FormProps {
  form: ControllersRegisterCompanyRequest;
  fieldErrors: FieldErrors;
  update: (partial: Partial<ControllersRegisterCompanyRequest>) => void;
}

export function CompanyStep2Form({ form, fieldErrors, update }: CompanyStep2FormProps) {
  return (
    <div className="flex flex-col gap-4">
      <FieldWithError
        label="Nome do responsável"
        error={fieldErrors.owner_name}
        value={form.owner_name}
        onChange={(e) => update({ owner_name: e.target.value })}
        placeholder="Seu nome completo"
      />
      <FieldWithError
        label="E-mail"
        error={fieldErrors.email}
        type="email"
        value={form.email}
        onChange={(e) => update({ email: e.target.value })}
        placeholder="seu@clinica.com"
      />
      <FieldWithError
        label="Telefone"
        error={fieldErrors.phone}
        value={form.phone}
        onChange={(e) => update({ phone: formatPhoneBr(e.target.value) })}
        placeholder="(11) 99999-9999"
      />
    </div>
  );
}

"use client";

import { FieldWithError } from "@/components/FieldWithError";
import { formatCpf, formatPhoneBr } from "@/lib/masks";
import type { ControllersRegisterShiftVeterinaryRequest } from "@/api/generated/api";

type FieldErrors = Partial<Record<keyof ControllersRegisterShiftVeterinaryRequest | "specialties", string>>;

export interface VeterinaryStep1FormProps {
  form: ControllersRegisterShiftVeterinaryRequest;
  fieldErrors: FieldErrors;
  update: (partial: Partial<ControllersRegisterShiftVeterinaryRequest>) => void;
}

export function VeterinaryStep1Form({ form, fieldErrors, update }: VeterinaryStep1FormProps) {
  return (
    <div className="flex flex-col gap-4">
      <FieldWithError
        label="Nome completo"
        error={fieldErrors.full_name}
        value={form.full_name}
        onChange={(e) => update({ full_name: e.target.value })}
        placeholder="Seu nome completo"
      />
      <FieldWithError
        label="CPF"
        error={fieldErrors.cpf}
        value={form.cpf}
        onChange={(e) => update({ cpf: formatCpf(e.target.value) })}
        placeholder="000.000.000-00"
      />
      <FieldWithError
        label="E-mail"
        error={fieldErrors.email}
        type="email"
        value={form.email}
        onChange={(e) => update({ email: e.target.value })}
        placeholder="seu@email.com"
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

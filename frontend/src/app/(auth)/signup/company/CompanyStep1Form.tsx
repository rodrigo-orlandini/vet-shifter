"use client";

import { useState } from "react";
import { FieldWithError } from "@/components/ui/FieldWithError";
import { FormField } from "@/components/ui/FormField";
import { Input } from "@/components/ui/Input";
import { Select } from "@/components/ui/Select";
import { formatCep, formatCnpj } from "@/lib/masks";
import { fetchAddressByCep } from "@/lib/cep";
import type { ControllersRegisterCompanyRequest } from "@/api/generated/api";
import { BRAZIL_UFS } from "@/constants/ufs";

type FieldErrors = Partial<Record<keyof ControllersRegisterCompanyRequest, string>>;

export interface CompanyStep1FormProps {
  form: ControllersRegisterCompanyRequest;
  fieldErrors: FieldErrors;
  update: (partial: Partial<ControllersRegisterCompanyRequest>) => void;
}

export function CompanyStep1Form({ form, fieldErrors, update }: CompanyStep1FormProps) {
  const [cepLoading, setCepLoading] = useState(false);

  const handleCepBlur = async () => {
    const digits = (form.zip_code ?? "").replace(/\D/g, "");
    if (digits.length !== 8) return;

    setCepLoading(true);

    try {
      const addr = await fetchAddressByCep(form.zip_code ?? "");
      if (addr) {
        update({
          street: addr.street,
          city: addr.city,
          state: addr.state,
        });
      }
    } finally {
      setCepLoading(false);
    }
  };

  return (
    <div className="flex flex-col gap-4">
      <FieldWithError
        label="Nome da clínica"
        requiredMark
        error={fieldErrors.company_name}
        value={form.company_name}
        onChange={(e) => update({ company_name: e.target.value })}
        placeholder="Ex: Clínica Vida Animal"
      />

      <FieldWithError
        label="CNPJ"
        requiredMark
        error={fieldErrors.cnpj}
        value={form.cnpj}
        onChange={(e) => update({ cnpj: formatCnpj(e.target.value) })}
        placeholder="00.000.000/0001-00"
      />

      <FieldWithError
        label="CEP"
        requiredMark
        error={fieldErrors.zip_code}
        value={form.zip_code ?? ""}
        onChange={(e) => update({ zip_code: formatCep(e.target.value) })}
        onBlur={handleCepBlur}
        placeholder="00000-000"
        helperText="Digite o CEP para preencher o endereço automaticamente."
      />

      {cepLoading && (
        <p className="flex items-center gap-2 text-xs text-ink-muted">
          <span className="inline-block h-3.5 w-3.5 animate-spin rounded-full border-2 border-primary border-t-transparent" />
          Buscando endereço…
        </p>
      )}

      <FormField label="Logradouro / Rua" required htmlFor="company-street" error={fieldErrors.street}>
        <Input
          id="company-street"
          value={form.street ?? ""}
          onChange={(e) => update({ street: e.target.value })}
          placeholder="Ex: Rua das Flores"
          hasError={Boolean(fieldErrors.street)}
        />
      </FormField>

      <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
        <FieldWithError
          label="Número"
          requiredMark
          error={fieldErrors.number}
          value={form.number ?? ""}
          onChange={(e) => update({ number: e.target.value })}
          placeholder="Nº"
        />
        <FieldWithError
          label="Cidade"
          requiredMark
          error={fieldErrors.city}
          value={form.city ?? ""}
          onChange={(e) => update({ city: e.target.value })}
          placeholder="Cidade"
        />
      </div>

      <FormField label="Estado" required htmlFor="company-uf" error={fieldErrors.state}>
        <Select
          id="company-uf"
          value={form.state ?? ""}
          onChange={(e) => update({ state: e.target.value })}
          hasError={Boolean(fieldErrors.state)}
        >
          <option value="">Selecionar estado</option>
          {BRAZIL_UFS.map((uf) => (
            <option key={uf.code} value={uf.code}>
              {uf.code}
            </option>
          ))}
        </Select>
      </FormField>
    </div>
  );
}

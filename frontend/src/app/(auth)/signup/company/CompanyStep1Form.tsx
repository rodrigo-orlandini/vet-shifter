"use client";

import { FieldWithError } from "@/components/FieldWithError";
import { formatCep, formatCnpj } from "@/lib/masks";
import type { ControllersRegisterCompanyRequest } from "@/api/generated/api";
import { BRAZIL_UFS } from "@/auth/constants/ufs";

type FieldErrors = Partial<Record<keyof ControllersRegisterCompanyRequest, string>>;

export interface CompanyStep1FormProps {
  form: ControllersRegisterCompanyRequest;
  fieldErrors: FieldErrors;
  update: (partial: Partial<ControllersRegisterCompanyRequest>) => void;
}

export function CompanyStep1Form({ form, fieldErrors, update }: CompanyStep1FormProps) {
  return (
    <div className="flex flex-col gap-4">
      <FieldWithError
        label="CNPJ"
        error={fieldErrors.cnpj}
        value={form.cnpj}
        onChange={(e) => update({ cnpj: formatCnpj(e.target.value) })}
        placeholder="00.000.000/0001-00"
      />
      <FieldWithError
        label="Razão social"
        error={fieldErrors.company_name}
        value={form.company_name}
        onChange={(e) => update({ company_name: e.target.value })}
        placeholder="Nome da clínica"
      />

      <label className="flex flex-col gap-1">
        <span className="text-sm font-medium text-neutral-700">Rua (opcional)</span>
        <input
          type="text"
          value={form.street ?? ""}
          onChange={(e) => update({ street: e.target.value })}
          className="rounded-lg border border-neutral-300 px-3 py-2 text-neutral-900"
        />
      </label>

      <div className="grid grid-cols-2 gap-3">
        <label className="flex flex-col gap-1">
          <span className="text-sm font-medium text-neutral-700">Número</span>
          <input
            type="text"
            value={form.number ?? ""}
            onChange={(e) => update({ number: e.target.value })}
            className="rounded-lg border border-neutral-300 px-3 py-2 text-neutral-900"
          />
        </label>
        <FieldWithError
          label="CEP"
          error={fieldErrors.zip_code}
          value={form.zip_code ?? ""}
          onChange={(e) => update({ zip_code: formatCep(e.target.value) })}
          placeholder="00000-000"
        />
      </div>

      <div className="grid grid-cols-2 gap-3">
        <label className="flex flex-col gap-1">
          <span className="text-sm font-medium text-neutral-700">Cidade</span>
          <input
            type="text"
            value={form.city ?? ""}
            onChange={(e) => update({ city: e.target.value })}
            className="rounded-lg border border-neutral-300 px-3 py-2 text-neutral-900"
          />
        </label>
        <label className="flex flex-col gap-1">
          <span className="text-sm font-medium text-neutral-700">Estado</span>
          <select
            value={form.state ?? ""}
            onChange={(e) => update({ state: e.target.value })}
            className="rounded-lg border border-neutral-300 px-3 py-2 text-neutral-900"
          >
            <option value="">Selecione a UF</option>
            {BRAZIL_UFS.map((uf) => (
              <option key={uf.code} value={uf.code}>
                {uf.code} - {uf.name}
              </option>
            ))}
          </select>
        </label>
      </div>
    </div>
  );
}

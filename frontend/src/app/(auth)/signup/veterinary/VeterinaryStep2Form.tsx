"use client";

import { FieldWithError } from "@/components/FieldWithError";
import { VETERINARY_SPECIALTIES, SPECIALTY_LABELS } from "@/auth/constants/specialties";
import { BRAZIL_UFS } from "@/auth/constants/ufs";
import type { ControllersRegisterShiftVeterinaryRequest } from "@/api/generated/api";

type FieldErrors = Partial<Record<keyof ControllersRegisterShiftVeterinaryRequest | "specialties", string>>;

export interface VeterinaryStep2FormProps {
  form: ControllersRegisterShiftVeterinaryRequest;
  fieldErrors: FieldErrors;
  update: (partial: Partial<ControllersRegisterShiftVeterinaryRequest>) => void;
  toggleSpecialty: (value: string) => void;
}

export function VeterinaryStep2Form({
  form,
  fieldErrors,
  update,
  toggleSpecialty,
}: VeterinaryStep2FormProps) {
  return (
    <div className="flex flex-col gap-4">
      <FieldWithError
        label="Número CRMV"
        error={fieldErrors.crmv_number}
        value={form.crmv_number}
        onChange={(e) => update({ crmv_number: e.target.value })}
        placeholder="12345"
      />
      <label className="flex flex-col gap-1">
        <span className={`text-sm font-medium ${fieldErrors.crmv_number ? "text-red-700" : "text-neutral-700"}`}>
          UF do CRMV
        </span>
        <select
          value={form.crmv_state}
          onChange={(e) => update({ crmv_state: e.target.value })}
          className={`rounded-lg border px-3 py-2 text-neutral-900 ${
            fieldErrors.crmv_number ? "border-red-500" : "border-neutral-300"
          }`}
        >
          <option value="">Selecione a UF</option>
          {BRAZIL_UFS.map((uf) => (
            <option key={uf.code} value={uf.code}>
              {uf.code} - {uf.name}
            </option>
          ))}
        </select>
      </label>
      <fieldset>
        {fieldErrors.specialties && (
          <span className="mb-1 block text-xs text-red-600" role="alert">
            {fieldErrors.specialties}
          </span>
        )}
        <legend className={`text-sm font-medium ${fieldErrors.specialties ? "text-red-700" : "text-neutral-700"}`}>
          Especialidades (escolha ao menos 1)
        </legend>
        <div className="mt-3 flex flex-wrap justify-center gap-2">
          {VETERINARY_SPECIALTIES.map((value) => {
            const selected = form.specialties.includes(value);
            const label = (SPECIALTY_LABELS[value] ?? value).toUpperCase();

            return (
              <button
                key={value}
                type="button"
                onClick={() => toggleSpecialty(value)}
                aria-pressed={selected}
                className={[
                  "rounded-lg px-2 py-0.5 text-xs font-semibold uppercase transition-colors",
                  "border",
                  selected
                    ? "bg-emerald-600 border-emerald-600 text-white hover:bg-emerald-700"
                    : "bg-neutral-200 border-neutral-200 text-neutral-800 hover:bg-neutral-300",
                ].join(" ")}
              >
                <div>{label}</div>
              </button>
            );
          })}
        </div>
      </fieldset>
    </div>
  );
}

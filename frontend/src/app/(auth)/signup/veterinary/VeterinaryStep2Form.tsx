"use client";

import { FieldWithError } from "@/components/FieldWithError";
import { VETERINARY_SPECIALTIES, SPECIALTY_LABELS } from "@/auth/constants/specialties";
import type { RegisterVeterinaryRequest } from "@/auth/types/veterinary-signup";

type FieldErrors = Partial<Record<keyof RegisterVeterinaryRequest | "specialties", string>>;

export interface VeterinaryStep2FormProps {
  form: RegisterVeterinaryRequest;
  fieldErrors: FieldErrors;
  update: (partial: Partial<RegisterVeterinaryRequest>) => void;
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
        <span
          className={`text-sm font-medium ${fieldErrors.crmv_number ? "text-red-700" : "text-neutral-700"}`}
        >
          UF do CRMV
        </span>
        <input
          type="text"
          value={form.crmv_state}
          onChange={(e) => update({ crmv_state: e.target.value.toUpperCase().slice(0, 2) })}
          className={`rounded-lg border px-3 py-2 text-neutral-900 ${
            fieldErrors.crmv_number ? "border-red-500" : "border-neutral-300"
          }`}
          placeholder="SP"
          maxLength={2}
        />
      </label>
      <fieldset>
        {fieldErrors.specialties && (
          <span className="mb-1 block text-xs text-red-600" role="alert">
            {fieldErrors.specialties}
          </span>
        )}
        <legend
          className={`text-sm font-medium ${fieldErrors.specialties ? "text-red-700" : "text-neutral-700"}`}
        >
          Especialidades (pelo menos uma)
        </legend>
        <div className="mt-2 grid grid-cols-1 gap-2">
          {VETERINARY_SPECIALTIES.map((value) => (
            <label key={value} className="flex items-center gap-2">
              <input
                type="checkbox"
                checked={form.specialties.includes(value)}
                onChange={() => toggleSpecialty(value)}
                className="rounded border-neutral-300"
              />
              <span className="text-sm text-neutral-800">{SPECIALTY_LABELS[value] ?? value}</span>
            </label>
          ))}
        </div>
      </fieldset>
    </div>
  );
}

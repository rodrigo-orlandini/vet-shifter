"use client";

import { FormField } from "@/components/ui/FormField";
import { Input } from "@/components/ui/Input";
import { Select } from "@/components/ui/Select";
import { VETERINARY_SPECIALTIES, SPECIALTY_LABELS } from "@/constants/specialties";
import { BRAZIL_UFS } from "@/constants/ufs";
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
    <div className="flex flex-col gap-5">
      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <FormField label="Número do CRMV" required htmlFor="crmv-number" error={fieldErrors.crmv_number}>
          <Input
            id="crmv-number"
            value={form.crmv_number}
            onChange={(e) => update({ crmv_number: e.target.value })}
            placeholder="12345"
            hasError={Boolean(fieldErrors.crmv_number)}
          />
        </FormField>

        <FormField label="Estado do CRMV" required htmlFor="crmv-uf" error={fieldErrors.crmv_state}>
          <Select
            id="crmv-uf"
            value={form.crmv_state}
            onChange={(e) => update({ crmv_state: e.target.value })}
            hasError={Boolean(fieldErrors.crmv_state)}
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

      <p className="text-[11px] text-ink-muted">Digite o número conforme consta em sua carteira profissional.</p>

      <fieldset className="min-w-0">
        {fieldErrors.specialties && (
          <span className="mb-2 block text-sm text-danger" role="alert">
            {fieldErrors.specialties}
          </span>
        )}
        
        <legend className="text-[13px] font-medium text-ink-body">
          Especialidades <span className="text-danger">*</span>
        </legend>

        <p className="mt-1 text-[11px] text-ink-muted">Selecione todas as especialidades em que você atua.</p>

        <div className="mt-3 grid grid-cols-2 gap-2 sm:grid-cols-3">
          {VETERINARY_SPECIALTIES.map((value) => {
            const selected = form.specialties.includes(value);
            const label = SPECIALTY_LABELS[value] ?? value;

            return (
              <button
                key={value}
                type="button"
                onClick={() => toggleSpecialty(value)}
                aria-pressed={selected}
                className={[
                  "rounded-full px-3 py-2 text-center text-xs font-medium transition-colors",
                  "cursor-pointer",
                  "focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/30 focus-visible:ring-offset-2 focus-visible:ring-offset-page",
                  selected
                    ? "border border-primary bg-primary text-surface hover:bg-primary/90"
                    : "border border-primary bg-surface text-primary hover:bg-primary/5 hover:text-primary",
                ].join(" ")}
              >
                {label}
              </button>
            );
          })}
        </div>
      </fieldset>
    </div>
  );
}

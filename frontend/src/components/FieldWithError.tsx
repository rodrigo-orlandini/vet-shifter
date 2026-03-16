import { type InputHTMLAttributes } from "react";

export interface FieldWithErrorProps extends Omit<InputHTMLAttributes<HTMLInputElement>, "className"> {
  label: string;
  error?: string | null;
  inputClassName?: string;
}

export function FieldWithError({
  label,
  error,
  inputClassName = "",
  ...inputProps
}: FieldWithErrorProps) {
  return (
    <label className="flex flex-col gap-1">
      {error && (
        <span className="text-xs text-red-600" role="alert">
          {error}
        </span>
      )}

      <span className={`text-sm font-medium ${error ? "text-red-700" : "text-neutral-700"}`}>
        {label}
      </span>
      
      <input
        {...inputProps}
        className={`rounded-lg border px-3 py-2 text-neutral-900 ${
          error ? "border-red-500 focus:outline-red-500" : "border-neutral-300"
        } ${inputClassName}`}
      />
    </label>
  );
}

import { useId, type InputHTMLAttributes } from "react";
import { FormField } from "./FormField";
import { Input } from "./Input";

export interface FieldWithErrorProps extends Omit<InputHTMLAttributes<HTMLInputElement>, "className"> {
  label: string;
  error?: string | null;
  requiredMark?: boolean;
  helperText?: string;
  inputClassName?: string;
}

export function FieldWithError({
  label,
  error,
  requiredMark = false,
  helperText,
  inputClassName = "",
  id,
  ...inputProps
}: FieldWithErrorProps) {
  const uid = useId();
  const inputId = id ?? (inputProps.name ? String(inputProps.name) : `field-${uid}`);

  return (
    <FormField
      label={label}
      required={requiredMark}
      error={error}
      helperText={helperText}
      htmlFor={inputId}
    >
      <Input
        id={inputId}
        hasError={Boolean(error)}
        inputClassName={inputClassName}
        {...inputProps}
      />
    </FormField>
  );
}

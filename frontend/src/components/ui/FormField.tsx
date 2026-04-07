import { type ReactNode } from "react";
import { Label } from "./Label";

export interface FormFieldProps {
  label?: string;
  required?: boolean;
  error?: string | null;
  helperText?: string;
  htmlFor?: string;
  children: ReactNode;
  className?: string;
}

export function FormField({
  label,
  required,
  error,
  helperText,
  htmlFor,
  children,
  className = "",
}: FormFieldProps) {
  return (
    <div className={`flex flex-col gap-1 ${className}`}>
      {label && (
        <Label htmlFor={htmlFor} required={required}>
          {label}
        </Label>
      )}
      {children}
      {helperText && !error && (
        <p className="text-[11px] text-ink-muted">{helperText}</p>
      )}
      {error && (
        <p className="text-sm text-danger" role="alert">
          {error}
        </p>
      )}
    </div>
  );
}

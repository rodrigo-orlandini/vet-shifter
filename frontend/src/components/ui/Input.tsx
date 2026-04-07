import { useId, type InputHTMLAttributes } from "react";

export interface InputProps extends Omit<InputHTMLAttributes<HTMLInputElement>, "className"> {
  hasError?: boolean;
  inputClassName?: string;
}

export function Input({ hasError = false, inputClassName = "", id, ...props }: InputProps) {
  const uid = useId();
  return (
    <input
      id={id ?? uid}
      {...props}
      aria-invalid={hasError || undefined}
      className={`h-11 w-full rounded-lg border px-3.5 text-sm text-ink-body placeholder:text-placeholder focus:outline-none focus:ring-2 focus:ring-primary/30 ${
        hasError
          ? "border-danger focus:border-danger"
          : "border-edge-input focus:border-primary"
      } ${inputClassName}`}
    />
  );
}

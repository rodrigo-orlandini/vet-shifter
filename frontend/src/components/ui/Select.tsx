import { useId, type SelectHTMLAttributes } from "react";

export interface SelectProps extends Omit<SelectHTMLAttributes<HTMLSelectElement>, "className"> {
  hasError?: boolean;
  selectClassName?: string;
}

export function Select({ hasError = false, selectClassName = "", id, children, ...props }: SelectProps) {
  const uid = useId();
  return (
    <select
      id={id ?? uid}
      {...props}
      className={`h-11 w-full rounded-lg border px-3.5 text-sm text-ink-body focus:outline-none focus:ring-2 focus:ring-primary/30 ${
        hasError
          ? "border-danger focus:border-danger"
          : "border-edge-input focus:border-primary"
      } ${selectClassName}`}
    >
      {children}
    </select>
  );
}

import { type InputHTMLAttributes } from "react";

export interface CheckboxProps extends Omit<InputHTMLAttributes<HTMLInputElement>, "type" | "className"> {
  checkboxClassName?: string;
}

export function Checkbox({ checkboxClassName = "", ...props }: CheckboxProps) {
  return (
    <input
      type="checkbox"
      {...props}
      className={`h-5 w-5 shrink-0 rounded border-edge-input accent-primary focus:ring-primary ${checkboxClassName}`}
    />
  );
}

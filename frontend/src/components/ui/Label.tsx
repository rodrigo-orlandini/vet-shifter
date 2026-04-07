import { type LabelHTMLAttributes } from "react";

export interface LabelProps extends LabelHTMLAttributes<HTMLLabelElement> {
  required?: boolean;
}

export function Label({ children, required, className = "", ...props }: LabelProps) {
  return (
    <label className={`text-[13px] font-medium text-ink-body ${className}`} {...props}>
      {children}
      {required && <span className="text-danger"> *</span>}
    </label>
  );
}

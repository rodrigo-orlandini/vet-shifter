"use client";

import { useId, useState } from "react";
import { getPasswordStrength, passwordRules } from "@/lib/passwordPolicy";

export interface PasswordFieldsProps {
  password: string;
  confirmPassword: string;
  onPasswordChange: (value: string) => void;
  onConfirmChange: (value: string) => void;
  passwordError?: string | null;
  confirmError?: string | null;
}

export function PasswordFields({
  password,
  confirmPassword,
  onPasswordChange,
  onConfirmChange,
  passwordError,
  confirmError,
}: PasswordFieldsProps) {
  const uid = useId();
  const [showPwd, setShowPwd] = useState(false);
  const [showConf, setShowConf] = useState(false);

  const rules = passwordRules(password);
  const strength = getPasswordStrength(password);

  const filledSegments = strength === "weak" ? 1 : strength === "medium" ? 2 : 3;
  const segmentColor =
    strength === "weak" ? "bg-danger" : strength === "medium" ? "bg-warning" : "bg-success";
  const strengthLabel =
    strength === "weak" ? "Fraca" : strength === "medium" ? "Média" : "Forte";
  const strengthTextColor =
    strength === "weak" ? "text-danger" : strength === "medium" ? "text-warning" : "text-success";

  const EyeIcon = ({ open }: { open: boolean }) =>
    open ? (
      <svg className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.5}>
        <path strokeLinecap="round" strokeLinejoin="round" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
      </svg>
    ) : (
      <svg className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.5}>
        <path strokeLinecap="round" strokeLinejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
        <path strokeLinecap="round" strokeLinejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
      </svg>
    );

  return (
    <div className="flex flex-col gap-4">
      <div className="flex flex-col gap-1.5">
        <label htmlFor={`${uid}-pwd`} className="text-[13px] font-medium text-ink-body">
          Senha <span className="text-danger">*</span>
        </label>

        <div className="relative">
          <input
            id={`${uid}-pwd`}
            type={showPwd ? "text" : "password"}
            autoComplete="new-password"
            value={password}
            onChange={(e) => onPasswordChange(e.target.value)}
            className={`h-11 w-full rounded-lg border px-3.5 pr-11 text-sm text-ink-body placeholder:text-placeholder focus:outline-none focus:ring-2 focus:ring-primary/30 ${
              passwordError ? "border-danger" : "border-edge-input focus:border-primary"
            }`}
            placeholder="••••••••"
          />

          <button
            type="button"
            className="absolute right-2 top-1/2 -translate-y-1/2 rounded p-1 text-ink-muted hover:bg-page"
            onClick={() => setShowPwd((s) => !s)}
            aria-label={showPwd ? "Ocultar senha" : "Mostrar senha"}
          >
            <EyeIcon open={showPwd} />
          </button>
        </div>

        <div className="flex gap-1">
          {[0, 1, 2].map((i) => (
            <div
              key={i}
              className={`h-1 flex-1 rounded-sm transition-colors ${i < filledSegments ? segmentColor : "bg-edge-alt"}`}
            />
          ))}
        </div>

        <p className={`text-[11px] font-medium ${strengthTextColor}`}>
          Força: {strengthLabel}
        </p>

        <ul className="space-y-1">
          {[
            { met: rules.minLen, text: "Mínimo 8 caracteres" },
            { met: rules.upper, text: "Pelo menos uma letra maiúscula" },
            { met: rules.digit, text: "Pelo menos um número" },
          ].map(({ met, text }) => (
            <li key={text} className={`flex items-center gap-1.5 text-[11px] ${met ? "text-success" : "text-ink-muted"}`}>
              {met ? (
                <svg className="h-3.5 w-3.5 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2.5}>
                  <path strokeLinecap="round" strokeLinejoin="round" d="M5 13l4 4L19 7" />
                </svg>
              ) : (
                <span className="h-3.5 w-3.5 shrink-0 rounded-full border border-placeholder" />
              )}
              {text}
            </li>
          ))}
        </ul>

        {passwordError && (
          <p className="text-sm text-danger" role="alert">
            {passwordError}
          </p>
        )}
      </div>

      <div className="flex flex-col gap-1.5">
        <label htmlFor={`${uid}-conf`} className="text-[13px] font-medium text-ink-body">
          Confirmar senha <span className="text-danger">*</span>
        </label>

        <div className="relative">
          <input
            id={`${uid}-conf`}
            type={showConf ? "text" : "password"}
            autoComplete="new-password"
            value={confirmPassword}
            onChange={(e) => onConfirmChange(e.target.value)}
            className={`h-11 w-full rounded-lg border px-3.5 pr-11 text-sm text-ink-body focus:outline-none focus:ring-2 focus:ring-primary/30 ${
              confirmError ? "border-danger" : "border-edge-input focus:border-primary"
            }`}
            placeholder="Repita a senha"
          />
          <button
            type="button"
            className="absolute right-2 top-1/2 -translate-y-1/2 rounded p-1 text-ink-muted hover:bg-page"
            onClick={() => setShowConf((s) => !s)}
            aria-label={showConf ? "Ocultar confirmação" : "Mostrar confirmação"}
          >
            <EyeIcon open={showConf} />
          </button>
        </div>

        {confirmError && (
          <p className="text-sm text-danger" role="alert">
            {confirmError}
          </p>
        )}
      </div>
    </div>
  );
}

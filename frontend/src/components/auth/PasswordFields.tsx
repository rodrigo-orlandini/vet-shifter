"use client";

import { useId, useState } from "react";
import { getPasswordStrength, passwordRules } from "@/lib/passwordPolicy";
import { EyeOffIcon } from "@/components/icons/EyeOffIcon";
import { EyeIcon } from "@/components/icons/EyeIcon";
import { CheckIcon } from "@/components/icons/CheckIcon";

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

  const ToggleEyeIcon = ({ open }: { open: boolean }) =>
    open ? (
      <EyeOffIcon className="h-5 w-5" />
    ) : (
      <EyeIcon className="h-5 w-5" />
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
            <ToggleEyeIcon open={showPwd} />
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
                <CheckIcon className="h-3.5 w-3.5 shrink-0" />
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
            <ToggleEyeIcon open={showConf} />
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

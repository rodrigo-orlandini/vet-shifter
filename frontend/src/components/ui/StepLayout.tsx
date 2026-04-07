"use client";

import { useCallback, useEffect, useRef, type ReactNode } from "react";
import { ArrowRightIcon } from "@/components/icons/ArrowRightIcon";
import { Button } from "./Button";

export interface StepLayoutProps {
  onBack?: () => void;
  onNext?: () => void;
  onSubmit?: () => void;
  isFirstStep: boolean;
  isLastStep: boolean;
  nextLabel?: string;
  submitLabel?: string;
  loading?: boolean;
  submitDisabled?: boolean;
  children: ReactNode;
  className?: string;
}

export function StepLayout({
  onBack,
  onNext,
  onSubmit,
  isFirstStep,
  isLastStep,
  nextLabel = "Próximo",
  submitLabel = "Enviar",
  loading = false,
  submitDisabled = false,
  children,
  className = "",
}: StepLayoutProps) {
  const rootRef = useRef<HTMLDivElement | null>(null);

  const handlePrimary = useCallback(() => {
    if (isLastStep) {
      onSubmit?.();
    } else {
      onNext?.();
    }
  }, [isLastStep, onNext, onSubmit]);

  useEffect(() => {
    const onKeyDown = (e: KeyboardEvent) => {
      if (e.key !== "Enter" || e.repeat || loading || (isLastStep && submitDisabled)) return;

      const rootEl = rootRef.current;
      if (!rootEl) return;

      const active = document.activeElement as HTMLElement | null;
      const tag = active?.tagName?.toUpperCase() ?? "";

      if (tag === "TEXTAREA") return;

      const target = e.target as Node | null;
      const isInside = target ? rootEl.contains(target) : false;

      const isBodyFocus = tag === "BODY" || tag === "HTML";

      if (!isInside && !isBodyFocus) return;

      e.preventDefault();
      handlePrimary();
    };

    window.addEventListener("keydown", onKeyDown);
    return () => window.removeEventListener("keydown", onKeyDown);
  }, [handlePrimary, loading, isLastStep, submitDisabled]);

  return (
    <div ref={rootRef} className={`flex flex-col gap-6 ${className}`}>
      <div className="min-h-[120px]">{children}</div>

      <div className="flex w-full flex-col gap-3 sm:flex-row sm:items-stretch sm:justify-between sm:gap-4">
        <div className="sm:flex-1">
          {!isFirstStep && onBack && (
            <Button variant="back" type="button" onClick={onBack} disabled={loading} className="w-full sm:w-auto">
              ← Voltar
            </Button>
          )}
        </div>

        <Button
          type="button"
          onClick={handlePrimary}
          loading={loading}
          disabled={isLastStep && submitDisabled}
          className={`w-full sm:min-w-[160px] ${!isFirstStep ? "sm:flex-1" : "sm:ml-auto"}`}
        >
          {isLastStep ? (
            submitLabel
          ) : (
            <span className="inline-flex items-center gap-2">
              {nextLabel}
              <ArrowRightIcon className="h-4 w-4" />
            </span>
          )}
        </Button>
      </div>
    </div>
  );
}

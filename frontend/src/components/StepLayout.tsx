"use client";

import { useCallback, useEffect, useRef, type ReactNode } from "react";
import { StepIndicator } from "./StepIndicator";
import { Button } from "./Button";

export interface StepLayoutProps {
  currentStep: number;
  totalSteps: number;
  stepLabels?: string[];
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
  currentStep,
  totalSteps,
  stepLabels,
  onBack,
  onNext,
  onSubmit,
  isFirstStep,
  isLastStep,
  nextLabel = "Continuar",
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
      <StepIndicator
        currentStep={currentStep}
        totalSteps={totalSteps}
        stepLabels={stepLabels}
      />

      <div className="min-h-[200px]">{children}</div>

      <div className="flex gap-3 justify-between">
        <div>
          {!isFirstStep && onBack && (
            <Button variant="secondary" type="button" onClick={onBack} disabled={loading}>
              Voltar
            </Button>
          )}
        </div>

        <Button
          type="button"
          onClick={handlePrimary}
          loading={loading}
          disabled={isLastStep && submitDisabled}
        >
          {isLastStep ? submitLabel : nextLabel}
        </Button>
      </div>
    </div>
  );
}

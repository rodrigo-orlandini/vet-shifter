"use client";

import { type ReactNode } from "react";
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
  children,
  className = "",
}: StepLayoutProps) {
  const handlePrimary = () => {
    if (isLastStep) {
      onSubmit?.();
    } else {
      onNext?.();
    }
  };

  return (
    <div className={`flex flex-col gap-6 ${className}`}>
      <StepIndicator
        currentStep={currentStep}
        totalSteps={totalSteps}
        stepLabels={stepLabels}
      />

      <div className="min-h-[200px]">{children}</div>

      <div className="flex gap-3 justify-between">
        <div>
          {!isFirstStep && onBack && (
            <Button variant="secondary" type="button" onClick={onBack}>
              Voltar
            </Button>
          )}
        </div>

        <Button type="button" onClick={handlePrimary}>
          {isLastStep ? submitLabel : nextLabel}
        </Button>
      </div>
    </div>
  );
}

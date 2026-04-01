import { Fragment, type ReactElement } from "react";

function desktopStepItems(
  currentStep: number,
  totalSteps: number,
  stepLabels: string[] | undefined,
): ReactElement[] {
  const items: ReactElement[] = [];

  for (let i = 0; i < totalSteps; i++) {
    const step = i + 1;
    const isCompleted = step < currentStep;
    const isCurrent = step === currentStep;
    const prevCompleted = step - 1 < currentStep;

    if (i > 0) {
      items.push(
        <div
          key={`line-${step}`}
          className={`h-0.5 w-20 shrink-0 self-center ${prevCompleted ? "bg-primary" : "bg-edge"}`}
          aria-hidden
        />,
      );
    }

    items.push(
      <div key={`step-${step}`} className="flex flex-nowrap items-center gap-2.5">
        <div
          className={`flex h-8 w-8 shrink-0 items-center justify-center rounded-full text-sm font-bold transition-colors ${
            isCompleted || isCurrent
              ? "bg-primary text-surface"
              : "border-[1.5px] border-edge bg-surface text-ink-subtle"
          }`}
        >
          {isCompleted ? (
            <svg className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2.5}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M5 13l4 4L19 7" />
            </svg>
          ) : (
            step
          )}
        </div>
        {stepLabels?.[i] ? (
          <span
            className={`whitespace-nowrap text-sm mr-2 ${
              isCompleted
                ? "font-medium text-primary"
                : isCurrent
                  ? "font-semibold text-primary"
                  : "font-medium text-ink-subtle"
            }`}
          >
            {stepLabels[i]}
          </span>
        ) : null}
      </div>,
    );
  }

  return items;
}

export interface StepIndicatorProps {
  currentStep: number;
  totalSteps: number;
  stepLabels?: string[];
  className?: string;
}

export function StepIndicator({
  currentStep,
  totalSteps,
  stepLabels,
  className = "",
}: StepIndicatorProps) {
  const currentLabel = stepLabels?.[currentStep - 1];

  return (
    <>
      {/* Mobile: full-width white bar with bottom border */}
      <div
        className={`md:hidden -mx-4 sm:-mx-6 -mt-6 sm:-mt-10 bg-surface border-b border-edge h-12 flex items-center justify-center gap-3 mb-5 ${className}`}
        role="progressbar"
        aria-valuenow={currentStep}
        aria-valuemin={1}
        aria-valuemax={totalSteps}
        aria-label={`Etapa ${currentStep} de ${totalSteps}`}
      >
        {Array.from({ length: totalSteps }, (_, i) => {
          const step = i + 1;
          const isCompleted = step < currentStep;
          const isCurrent = step === currentStep;

          return (
            <Fragment key={step}>
              {isCompleted ? (
                <span className="flex h-2.5 w-2.5 shrink-0 items-center justify-center rounded-full bg-success" aria-hidden>
                  <svg className="h-[7px] w-[7px] text-surface" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={3.5}>
                    <path strokeLinecap="round" strokeLinejoin="round" d="M5 13l4 4L19 7" />
                  </svg>
                </span>
              ) : isCurrent ? (
                <span className="h-2.5 w-2.5 shrink-0 rounded-full bg-primary" aria-hidden />
              ) : (
                <span className="h-2 w-2 shrink-0 rounded-full bg-edge-input" aria-hidden />
              )}
            </Fragment>
          );
        })}

        {currentLabel && (
          <span className="text-[13px] font-medium text-primary">{currentLabel}</span>
        )}
      </div>

      <div
        className={`hidden md:flex flex-nowrap items-center justify-center gap-0 ${className}`}
        role="progressbar"
        aria-valuenow={currentStep}
        aria-valuemin={1}
        aria-valuemax={totalSteps}
        aria-label={`Etapa ${currentStep} de ${totalSteps}`}
      >
        {desktopStepItems(currentStep, totalSteps, stepLabels)}
      </div>
    </>
  );
}

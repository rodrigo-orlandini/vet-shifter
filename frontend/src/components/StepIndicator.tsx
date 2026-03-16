import React from "react";

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
  const columnCount = totalSteps * 2 - 1;
  const gridCols = Array.from(
    { length: columnCount },
    (_, i) => (i % 2 === 0 ? "auto" : "1fr")
  ).join(" ");

  return (
    <div
      className={`w-full ${className}`}
      role="progressbar"
      aria-valuenow={currentStep}
      aria-valuemin={1}
      aria-valuemax={totalSteps}
      aria-label={`Etapa ${currentStep} de ${totalSteps}`}
    >
      <div
        className="grid items-start"
        style={{ gridTemplateColumns: gridCols }}
      >
        {Array.from({ length: totalSteps }, (_, i) => {
          const step = i + 1;
          const isCompleted = step < currentStep;
          const isCurrent = step === currentStep;
          const prevCompleted = step - 1 < currentStep;

          return (
            <React.Fragment key={step}>
              {i > 0 && (
                <div
                  className="flex h-10 items-center"
                  aria-hidden
                >
                  <div
                    className={`w-full ${prevCompleted ? "bg-emerald-500" : "bg-neutral-200"}`}
                    style={{ height: 2 }}
                  />
                </div>
              )}

              <div className="flex flex-col items-center justify-center">
                <div
                  className={`
                    flex h-10 w-10 shrink-0 items-center justify-center rounded-full border-2 text-sm font-semibold
                    transition-colors duration-200
                    ${
                      isCompleted
                        ? "border-emerald-500 bg-emerald-500 text-white"
                        : isCurrent
                          ? "border-emerald-500 bg-white text-emerald-600 ring-4 ring-emerald-100"
                          : "border-neutral-300 bg-white text-neutral-400"
                    }
                  `}
                >
                  {isCompleted ? (
                    <svg
                      className="h-5 w-5"
                      fill="none"
                      viewBox="0 0 24 24"
                      stroke="currentColor"
                      strokeWidth={2.5}
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M5 13l4 4L19 7"
                      />
                    </svg>
                  ) : (
                    step
                  )}
                </div>
                {stepLabels?.[i] && (
                  <span
                    className={`mt-2 hidden text-center text-xs font-medium sm:block ${
                      isCurrent ? "text-emerald-600" : isCompleted ? "text-neutral-600" : "text-neutral-400"
                    }`}
                  >
                    {stepLabels[i]}
                  </span>
                )}
              </div>
            </React.Fragment>
          );
        })}
      </div>

      {!stepLabels && (
        <p className="mt-3 text-center text-sm font-medium text-neutral-500">
          Etapa {currentStep} de {totalSteps}
        </p>
      )}
    </div>
  );
}

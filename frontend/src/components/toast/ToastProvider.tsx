"use client";

import { createContext, useCallback, useContext, useMemo, useRef, useState } from "react";

export type ToastTone = "error" | "success" | "info";

export type ToastPayload = {
  id?: string;
  message: string;
  tone?: ToastTone;
  durationMs?: number;
};

type Toast = Required<Pick<ToastPayload, "id" | "message" | "tone">> & {
  durationMs: number;
};

type ToastContextValue = {
  pushToast: (payload: ToastPayload) => void;
};

const ToastContext = createContext<ToastContextValue | null>(null);

function makeId() {
  return `${Date.now()}_${Math.random().toString(16).slice(2)}`;
}

export function ToastProvider({ children }: { children: React.ReactNode }) {
  const [toasts, setToasts] = useState<Toast[]>([]);
  const timeoutsRef = useRef<Record<string, number>>({});

  const removeToast = useCallback((id: string) => {
    setToasts((prev) => prev.filter((t) => t.id !== id));
    const timeoutId = timeoutsRef.current[id];
    if (timeoutId) window.clearTimeout(timeoutId);
    delete timeoutsRef.current[id];
  }, []);

  const pushToast = useCallback(
    (payload: ToastPayload) => {
      const toast: Toast = {
        id: payload.id ?? makeId(),
        message: payload.message,
        tone: payload.tone ?? "info",
        durationMs: payload.durationMs ?? 2800,
      };

      setToasts((prev) => [...prev, toast]);

      timeoutsRef.current[toast.id] = window.setTimeout(
        () => removeToast(toast.id),
        toast.durationMs,
      );
    },
    [removeToast],
  );

  const value = useMemo<ToastContextValue>(() => ({ pushToast }), [pushToast]);

  return (
    <ToastContext.Provider value={value}>
      {children}
      <div
        className="pointer-events-none fixed right-4 top-4 z-50 flex flex-col gap-2"
        aria-live="polite"
        aria-relevant="additions"
      >
        {toasts.map((t) => {
          const toneStyles =
            t.tone === "success"
              ? "bg-emerald-600 text-white"
              : t.tone === "error"
                ? "bg-red-600 text-white"
                : "bg-neutral-900 text-white";

          return (
            <div
              key={t.id}
              className={`${toneStyles} pointer-events-auto rounded-lg px-3 py-2 text-sm shadow-lg`}
              role={t.tone === "error" ? "alert" : "status"}
            >
              {t.message}
            </div>
          );
        })}
      </div>
    </ToastContext.Provider>
  );
}

export function useToast() {
  const ctx = useContext(ToastContext);
  if (!ctx) {
    throw new Error("useToast must be used within <ToastProvider />");
  }
  return ctx;
}


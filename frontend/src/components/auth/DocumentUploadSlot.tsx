"use client";

import { useId, useRef } from "react";
import { Badge } from "@/components/ui/Badge";

export interface DocumentUploadSlotProps {
  title?: string;
  description?: string;
  file: File | null;
  onFile: (file: File | null) => void;
  required?: boolean;
  compact?: boolean;
}

export function DocumentUploadSlot({
  title,
  description,
  file,
  onFile,
  required = true,
  compact = false,
}: DocumentUploadSlotProps) {
  const uid = useId();
  const inputRef = useRef<HTMLInputElement>(null);

  const body = (
    <>
      <input
        ref={inputRef}
        id={uid}
        type="file"
        accept=".pdf,.jpg,.jpeg,.png,application/pdf,image/jpeg,image/png"
        className="sr-only"
        onChange={(e) => {
          const f = e.target.files?.[0] ?? null;
          onFile(f);
        }}
      />

      <button
        type="button"
        onClick={() => inputRef.current?.click()}
        className="flex w-full flex-col items-center justify-center gap-2 rounded-lg border border-edge-input bg-page/50 px-4 py-6 text-center transition-colors hover:border-primary/50"
      >
        <svg className="h-7 w-7 text-placeholder" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.25}>
          <path strokeLinecap="round" strokeLinejoin="round" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
        </svg>

        <span className="text-sm font-medium text-ink-muted">Arraste ou clique para enviar</span>
        <span className="text-[11px] text-placeholder">PDF, JPG ou PNG — máx. 10 MB</span>
      </button>

      <div className="mt-3 flex items-center justify-between text-[11px]">
        <span className="flex items-center gap-1.5 text-ink-muted">
          <span className="h-1.5 w-1.5 rounded-full bg-placeholder" aria-hidden />
          {file ? (
            <span className="text-ink-body">
              {file.name} ({(file.size / 1024 / 1024).toFixed(2)} MB)
            </span>
          ) : (
            "Não enviado"
          )}
        </span>

        {file && (
          <button
            type="button"
            className="font-medium text-primary hover:underline"
            onClick={() => {
              onFile(null);
              if (inputRef.current) inputRef.current.value = "";
            }}
          >
            Remover
          </button>
        )}
      </div>
    </>
  );

  if (compact) {
    return <div className="min-w-0">{body}</div>;
  }

  return (
    <div className="rounded-lg border border-edge bg-surface p-4">
      {(title || description || required) && (
        <div className="mb-3 flex items-start justify-between gap-2">
          <div className="flex items-center gap-2">
            <svg className="h-5 w-5 shrink-0 text-primary" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.5}>
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
              />
            </svg>

            <div>
              {title ? <p className="text-sm font-semibold text-ink-body">{title}</p> : null}
              {description && <p className="text-[11px] text-ink-muted">{description}</p>}
            </div>
          </div>
          {required && <Badge variant="warning" className="shrink-0">Obrigatório</Badge>}
        </div>
      )}

      {body}
    </div>
  );
}

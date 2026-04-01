"use client";

import { useId, useRef } from "react";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { CloudUploadIcon } from "@/components/icons/CloudUploadIcon";
import { FileTextIcon } from "@/components/icons/FileTextIcon";

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
        <CloudUploadIcon className="h-7 w-7 text-placeholder" />

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
          <Button
            type="button"
            variant="link"
            onClick={() => {
              onFile(null);
              if (inputRef.current) inputRef.current.value = "";
            }}
          >
            Remover
          </Button>
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
            <FileTextIcon className="h-5 w-5 shrink-0 text-primary" />

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

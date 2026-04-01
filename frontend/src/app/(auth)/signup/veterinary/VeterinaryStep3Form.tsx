"use client";

import { DocumentUploadSlot } from "@/components/auth/DocumentUploadSlot";
import { DocumentReviewNotice } from "@/components/auth/DocumentReviewNotice";
import { SkipUploadButton } from "@/components/auth/SkipUploadButton";

export type VetDocKey = "idDoc" | "crmvFront" | "crmvBack" | "diploma";

export interface VeterinaryStep3FormProps {
  files: Record<VetDocKey, File | null>;
  onFile: (key: VetDocKey, file: File | null) => void;
  onSkipUploads?: () => void;
}

export function VeterinaryStep3Form({ files, onFile, onSkipUploads }: VeterinaryStep3FormProps) {
  return (
    <div className="flex flex-col gap-4">
      <DocumentReviewNotice />

      <DocumentUploadSlot title="RG ou CNH" required={false} file={files.idDoc} onFile={(f) => onFile("idDoc", f)} />

      <div className="rounded-lg border border-[#E9ECEF] p-4">
        <div className="mb-3 flex items-center gap-2">
          <svg className="h-5 w-5 text-[#2A9D8F]" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.5}>
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
            />
          </svg>
          <p className="text-sm font-semibold text-ink-body">Carteira do CRMV</p>
        </div>

        <div className="flex flex-col gap-4 md:flex-row">
          <div className="flex-1">
            <p className="mb-2 text-xs font-medium text-[#6C757D]">Frente</p>
            <DocumentUploadSlot
              compact
              required={false}
              file={files.crmvFront}
              onFile={(f) => onFile("crmvFront", f)}
            />
          </div>
          <div className="flex-1">
            <p className="mb-2 text-xs font-medium text-[#6C757D]">Verso</p>
            <DocumentUploadSlot
              compact
              required={false}
              file={files.crmvBack}
              onFile={(f) => onFile("crmvBack", f)}
            />
          </div>
        </div>
        <p className="mt-2 text-[11px] text-[#6C757D]">Envie frente e verso da sua carteira do CRMV.</p>
      </div>

      <DocumentUploadSlot
        title="Diploma de Graduação em Medicina Veterinária"
        description="Medicina Veterinária"
        required={false}
        file={files.diploma}
        onFile={(f) => onFile("diploma", f)}
      />

      {onSkipUploads && <SkipUploadButton onClick={onSkipUploads} />}
    </div>
  );
}

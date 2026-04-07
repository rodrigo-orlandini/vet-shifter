"use client";

import { DocumentUploadSlot } from "@/components/auth/DocumentUploadSlot";
import { DocumentReviewNotice } from "@/components/auth/DocumentReviewNotice";
import { SkipUploadButton } from "@/components/auth/SkipUploadButton";

export type CompanyDocKey = "cnpjCard" | "alvara" | "idDoc";

export interface CompanyStep3FormProps {
  files: Record<CompanyDocKey, File | null>;
  onFile: (key: CompanyDocKey, file: File | null) => void;
  onSkipUploads?: () => void;
}

export function CompanyStep3Form({ files, onFile, onSkipUploads }: CompanyStep3FormProps) {
  return (
    <div className="flex flex-col gap-4">
      <DocumentReviewNotice />

      <DocumentUploadSlot
        title="Cartão CNPJ"
        required={false}
        file={files.cnpjCard}
        onFile={(f) => onFile("cnpjCard", f)}
      />
      <DocumentUploadSlot
        title="Alvará de Funcionamento"
        required={false}
        file={files.alvara}
        onFile={(f) => onFile("alvara", f)}
      />
      <DocumentUploadSlot
        title="RG ou CNH do Responsável"
        description="Documento de identidade do responsável"
        required={false}
        file={files.idDoc}
        onFile={(f) => onFile("idDoc", f)}
      />

      {onSkipUploads && <SkipUploadButton onClick={onSkipUploads} />}
    </div>
  );
}

export interface SkipUploadButtonProps {
  onClick: () => void;
}

export function SkipUploadButton({ onClick }: SkipUploadButtonProps) {
  return (
    <button
      type="button"
      onClick={onClick}
      className="flex items-center justify-center gap-2 text-sm font-medium text-[#2A9D8F] hover:underline"
    >
      Pular esta etapa e enviar depois
      <svg className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
        <path strokeLinecap="round" strokeLinejoin="round" d="M14 5l7 7m0 0l-7 7m7-7H3" />
      </svg>
    </button>
  );
}

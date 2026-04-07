import { ArrowRightIcon } from "@/components/icons/ArrowRightIcon";

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
      <ArrowRightIcon className="h-4 w-4" />
    </button>
  );
}

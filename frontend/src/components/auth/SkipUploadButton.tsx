import { Button } from "@/components/ui/Button";
import { ArrowRightIcon } from "@/components/icons/ArrowRightIcon";

export interface SkipUploadButtonProps {
  onClick: () => void;
}

export function SkipUploadButton({ onClick }: SkipUploadButtonProps) {
  return (
    <Button type="button" variant="link" onClick={onClick}>
      Pular esta etapa e enviar depois
      <ArrowRightIcon className="h-4 w-4" />
    </Button>
  );
}

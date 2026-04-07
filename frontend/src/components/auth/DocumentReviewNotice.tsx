import { InfoCircleIcon } from "@/components/icons/InfoCircleIcon";

export function DocumentReviewNotice() {
  return (
    <div className="flex gap-3 rounded-lg bg-[#E8F4FD] px-4 py-3 text-[#2B6CB0]">
      <InfoCircleIcon className="mt-0.5 h-5 w-5 shrink-0" />
      <p className="text-sm leading-relaxed">
        Seus documentos serão analisados pela nossa equipe em até 2 dias úteis. Você receberá um e-mail com o
        resultado.
      </p>
    </div>
  );
}


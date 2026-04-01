export function DocumentReviewNotice() {
  return (
    <div className="flex gap-3 rounded-lg bg-[#E8F4FD] px-4 py-3 text-[#2B6CB0]">
      <svg className="mt-0.5 h-5 w-5 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.5}>
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>
      <p className="text-sm leading-relaxed">
        Seus documentos serão analisados pela nossa equipe em até 2 dias úteis. Você receberá um e-mail com o
        resultado.
      </p>
    </div>
  );
}


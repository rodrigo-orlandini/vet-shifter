import Link from "next/link";

export default function TermsPage() {
  return (
    <div className="mx-auto min-h-screen max-w-2xl px-4 py-12 text-ink-body">
      <h1 className="text-2xl font-bold">Termos de Uso</h1>
      <p className="mt-4 text-sm leading-relaxed text-[#6C757D]">
        Esta página está em elaboração. Em breve você encontrará aqui os Termos de Uso completos do VetPlant.
      </p>
      <p className="mt-8">
        <Link href="/login" className="font-medium text-[#2A9D8F] hover:underline">
          Voltar ao login
        </Link>
      </p>
    </div>
  );
}

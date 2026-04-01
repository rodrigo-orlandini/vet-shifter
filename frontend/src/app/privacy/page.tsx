import Link from "next/link";

export default function PrivacyPage() {
  return (
    <div className="mx-auto min-h-screen max-w-2xl px-4 py-12 text-ink-body">
      <h1 className="text-2xl font-bold">Política de Privacidade</h1>
      <p className="mt-4 text-sm leading-relaxed text-[#6C757D]">
        Esta página está em elaboração. Em breve você encontrará aqui a Política de Privacidade completa do VetPlant,
        incluindo informações sobre tratamento de dados conforme a LGPD.
      </p>

      <p className="mt-8">
        <Link href="/login" className="font-medium text-[#2A9D8F] hover:underline">
          Voltar ao login
        </Link>
      </p>
    </div>
  );
}

import Link from "next/link";

export default function CompanyDashboardPage() {
  return (
    <div className="mx-auto max-w-2xl px-4 py-8">
      <h1 className="mb-2 text-2xl font-semibold text-neutral-900">Painel da empresa</h1>
      <p className="mb-6 text-neutral-600">
        Espaço reservado para a área da clínica/empresa. Aqui você gerenciará plantões e plantonistas.
      </p>
      <Link
        href="/login"
        className="text-sm font-medium text-emerald-600 hover:underline"
      >
        Voltar para entrar
      </Link>
    </div>
  );
}

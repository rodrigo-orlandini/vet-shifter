import Link from "next/link";

export default function VeterinaryDashboardPage() {
  return (
    <div className="mx-auto max-w-2xl px-4 py-8">
      <h1 className="mb-2 text-2xl font-semibold text-neutral-900">Painel do veterinário</h1>
      <p className="mb-6 text-neutral-600">
        Espaço reservado para a área do plantonista. Aqui você encontrará e gerenciará plantões.
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

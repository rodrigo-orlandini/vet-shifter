'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from '@/components/Link';
import Card from '@/components/Card';
import { api, getToken } from '@/utils/api';

type Shift = {
  id: string;
  company_id: string;
  starts_at: string;
  ends_at: string;
  type: string;
  offered_value_cents: number;
  requirements: string;
  description: string;
  location: string;
  status: string;
  created_at: string;
};

export default function ShiftsPage() {
  const router = useRouter();
  const [shifts, setShifts] = useState<Shift[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!getToken()) {
      router.push('/login');
      return;
    }
    api<{ shifts: Shift[] }>('/shifts')
      .then((data) => setShifts(data.shifts || []))
      .catch((err) => setError(err instanceof Error ? err.message : 'Erro ao carregar'))
      .finally(() => setLoading(false));
  }, [router]);

  const typeLabel: Record<string, string> = {
    emergency: 'Emergência',
    consultation: 'Consulta',
    surgery: 'Cirurgia',
  };
  const statusLabel: Record<string, string> = {
    open: 'Aberto',
    confirmed: 'Confirmado',
    cancelled: 'Cancelado',
    completed: 'Concluído',
  };

  return (
    <div className="min-h-screen bg-[#f0f9f7] px-4 py-8">
      <div className="max-w-3xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <Link href="/dashboard" className="text-teal-700 font-medium">
            ← Voltar
          </Link>
        </div>
        <h1 className="text-2xl font-bold text-teal-800 mb-6">Plantões disponíveis</h1>
        {loading && <p className="text-slate-600">Carregando...</p>}
        {error && (
          <p className="text-red-600 bg-red-50 p-3 rounded-lg mb-4">{error}</p>
        )}
        {!loading && !error && shifts.length === 0 && (
          <Card>
            <p className="text-slate-600 text-center">Nenhum plantão publicado no momento.</p>
          </Card>
        )}
        {!loading && shifts.length > 0 && (
          <div className="space-y-4">
            {shifts.map((s) => (
              <Card key={s.id} className="flex flex-col gap-2">
                <div className="flex justify-between items-start">
                  <span className="text-sm font-medium text-teal-700">
                    {typeLabel[s.type] || s.type}
                  </span>
                  <span className="text-xs px-2 py-1 rounded bg-slate-100 text-slate-600">
                    {statusLabel[s.status] || s.status}
                  </span>
                </div>
                <p className="text-slate-700">
                  {new Date(s.starts_at).toLocaleString('pt-BR')} –{' '}
                  {new Date(s.ends_at).toLocaleString('pt-BR', { timeStyle: 'short' })}
                </p>
                <p className="text-teal-800 font-semibold">
                  R$ {(s.offered_value_cents / 100).toFixed(2).replace('.', ',')}
                </p>
                {s.location && (
                  <p className="text-sm text-slate-600">{s.location}</p>
                )}
              </Card>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

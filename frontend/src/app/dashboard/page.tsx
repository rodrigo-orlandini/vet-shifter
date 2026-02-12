'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from '@/components/Link';
import Card from '@/components/Card';
import { getToken, clearToken } from '@/utils/api';

export default function DashboardPage() {
  const router = useRouter();
  const [role, setRole] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const token = getToken();
    if (!token) {
      router.push('/login');
      return;
    }
    try {
      const payload = JSON.parse(atob(token.split('.')[1]));
      setRole((payload as { role?: string }).role || '');
    } catch {
      clearToken();
      router.push('/login');
    } finally {
      setLoading(false);
    }
  }, [router]);

  const handleLogout = () => {
    clearToken();
    router.push('/login');
  };

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-[#f0f9f7]">
        <p className="text-slate-600">Carregando...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#f0f9f7] px-4 py-8">
      <div className="max-w-2xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-2xl font-bold text-teal-800">VetTroca</h1>
          <button
            type="button"
            onClick={handleLogout}
            className="text-sm text-slate-600 hover:text-teal-700"
          >
            Sair
          </button>
        </div>
        <Card>
          <p className="text-slate-600 mb-6">
            Você está logado como{' '}
            <strong className="text-teal-800">
              {role === 'clinic' ? 'Clínica' : 'Plantonista'}
            </strong>
            .
          </p>
          <div className="flex flex-col gap-3">
            <Link
              href="/shifts"
              className="block w-full py-3 px-4 rounded-lg bg-teal-600 text-white text-center font-medium hover:bg-teal-700 transition-colors"
            >
              Ver plantões
            </Link>
            {role === 'clinic' && (
              <Link
                href="/shifts/new"
                className="block w-full py-3 px-4 rounded-lg border-2 border-teal-600 text-teal-700 text-center font-medium hover:bg-teal-50 transition-colors"
              >
                Publicar plantão
              </Link>
            )}
          </div>
        </Card>
      </div>
    </div>
  );
}

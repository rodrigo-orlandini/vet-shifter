'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import Input from '@/components/Input';
import Button from '@/components/Button';
import Link from '@/components/Link';
import Card from '@/components/Card';
import { api, getToken } from '@/utils/api';

export default function NewShiftPage() {
  const router = useRouter();
  const [companyId, setCompanyId] = useState('');
  const [formData, setFormData] = useState({
    starts_at: '',
    ends_at: '',
    type: 'consultation',
    offered_value_cents: '',
    requirements: '',
    description: '',
    location: '',
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!getToken()) {
      router.push('/login');
      return;
    }
    setCompanyId(typeof window !== 'undefined' ? localStorage.getItem('vet_troca_company_id') || '' : '');
  }, [router]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    const cents = parseInt(formData.offered_value_cents.replace(/\D/g, ''), 10) || 0;
    if (!companyId) {
      setError('Informe o ID da clínica. (Em produção isso virá do seu perfil.)');
      return;
    }
    setIsLoading(true);
    try {
      await api('/shifts', {
        method: 'POST',
        body: JSON.stringify({
          company_id: companyId,
          starts_at: new Date(formData.starts_at).toISOString(),
          ends_at: new Date(formData.ends_at).toISOString(),
          type: formData.type,
          offered_value_cents: cents || 0,
          requirements: formData.requirements || undefined,
          description: formData.description || undefined,
          location: formData.location || undefined,
        }),
      });
      router.push('/shifts');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Falha ao publicar');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-[#f0f9f7] px-4 py-8">
      <div className="max-w-md mx-auto">
        <Link href="/dashboard" className="text-teal-700 font-medium mb-6 inline-block">
          ← Voltar
        </Link>
        <h1 className="text-2xl font-bold text-teal-800 mb-6">Publicar plantão</h1>
        <Card>
          <form onSubmit={handleSubmit}>
            <div className="space-y-6">
              {error && (
                <p className="text-sm text-red-600 bg-red-50 p-3 rounded-lg">{error}</p>
              )}
              <Input
                id="company_id"
                label="ID da clínica (UUID)"
                value={companyId}
                onChange={(e) => setCompanyId(e.target.value)}
                placeholder="550e8400-e29b-41d4-a716-446655440000"
                required
              />
              <Input
                id="starts_at"
                name="starts_at"
                label="Início"
                type="datetime-local"
                value={formData.starts_at}
                onChange={handleChange}
                required
              />
              <Input
                id="ends_at"
                name="ends_at"
                label="Fim"
                type="datetime-local"
                value={formData.ends_at}
                onChange={handleChange}
                required
              />
              <div>
                <label htmlFor="type" className="block text-sm font-medium text-gray-700 mb-2">
                  Tipo
                </label>
                <select
                  id="type"
                  name="type"
                  value={formData.type}
                  onChange={handleChange}
                  className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500"
                >
                  <option value="consultation">Consulta</option>
                  <option value="emergency">Emergência</option>
                  <option value="surgery">Cirurgia</option>
                </select>
              </div>
              <Input
                id="offered_value_cents"
                name="offered_value_cents"
                label="Valor ofertado (centavos ou R$)"
                type="text"
                value={formData.offered_value_cents}
                onChange={handleChange}
                placeholder="15000 ou 150,00"
              />
              <Input
                id="location"
                name="location"
                label="Local"
                type="text"
                value={formData.location}
                onChange={handleChange}
                placeholder="Endereço ou nome do local"
              />
              <div>
                <label htmlFor="description" className="block text-sm font-medium text-gray-700 mb-2">
                  Descrição (opcional)
                </label>
                <textarea
                  id="description"
                  name="description"
                  value={formData.description}
                  onChange={handleChange}
                  rows={3}
                  className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-teal-500"
                />
              </div>
              <Button type="submit" isLoading={isLoading} loadingText="Publicando...">
                Publicar plantão
              </Button>
            </div>
          </form>
        </Card>
      </div>
    </div>
  );
}

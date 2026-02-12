'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Input from '@/components/Input';
import Button from '@/components/Button';
import Link from '@/components/Link';
import Card from '@/components/Card';
import { formatCNPJ, formatPhone } from '@/utils/formatters';
import { api } from '@/utils/api';

export default function SignUpPage() {
  const router = useRouter();
  const [formData, setFormData] = useState({
    company_name: '',
    owner_name: '',
    email: '',
    password: '',
    phone: '',
    cnpj: '',
    consent_lgpd: false,
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type } = e.target;
    setFormData({
      ...formData,
      [name]: type === 'checkbox' ? (e.target as HTMLInputElement).checked : value,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!formData.consent_lgpd) {
      setError('É necessário aceitar o uso dos dados (LGPD).');
      return;
    }
    setError('');
    setIsLoading(true);
    try {
      const res = await api<{ company_id: string }>('/companies', {
        method: 'POST',
        body: JSON.stringify({
          cnpj: formData.cnpj.replace(/\D/g, ''),
          company_name: formData.company_name,
          owner_name: formData.owner_name,
          email: formData.email,
          phone: formData.phone.replace(/\D/g, ''),
          password: formData.password,
          consent_lgpd: true,
        }),
      });
      if (res.company_id) {
        localStorage.setItem('vet_troca_company_id', res.company_id);
      }
      router.push('/login');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Falha no cadastro');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-[#f0f9f7] px-4 py-8">
      <div className="w-full max-w-md">
        <div className="text-center mb-8">
          <h1 className="text-3xl font-bold text-teal-800">VetTroca</h1>
          <p className="text-slate-600 mt-1">Cadastro de clínica</p>
        </div>
        <Card>
          <form onSubmit={handleSubmit}>
            <div className="space-y-6">
              {error && (
                <p className="text-sm text-red-600 bg-red-50 p-3 rounded-lg">{error}</p>
              )}
              <Input
                id="company_name"
                name="company_name"
                label="Nome da clínica"
                type="text"
                value={formData.company_name}
                onChange={handleChange}
                required
                placeholder="Ex.: Clínica Veterinária São Paulo"
              />
              <Input
                id="owner_name"
                name="owner_name"
                label="Nome do responsável"
                type="text"
                value={formData.owner_name}
                onChange={handleChange}
                required
                placeholder="Seu nome completo"
              />

              <Input
                id="email"
                name="email"
                label="E-mail"
                type="email"
                value={formData.email}
                onChange={handleChange}
                required
                placeholder="seu@email.com"
              />

              <Input
                id="phone"
                name="phone"
                label="Telefone"
                type="tel"
                value={formData.phone}
                onChange={(e) =>
                  setFormData({ ...formData, phone: formatPhone(e.target.value) })
                }
                required
                placeholder="(11) 99999-9999"
                maxLength={15}
              />

              <Input
                id="cnpj"
                name="cnpj"
                label="CNPJ"
                type="text"
                value={formData.cnpj}
                onChange={(e) =>
                  setFormData({ ...formData, cnpj: formatCNPJ(e.target.value) })
                }
                required
                placeholder="00.000.000/0000-00"
                maxLength={18}
              />

              <Input
                id="password"
                name="password"
                label="Senha"
                type="password"
                value={formData.password}
                onChange={handleChange}
                required
                minLength={6}
                placeholder="••••••••"
              />

              <label className="flex items-center gap-2 text-sm text-slate-700">
                <input
                  type="checkbox"
                  name="consent_lgpd"
                  checked={formData.consent_lgpd}
                  onChange={handleChange}
                  className="rounded border-slate-300 text-teal-600 focus:ring-teal-500"
                />
                Aceito o uso dos meus dados conforme a política de privacidade (LGPD).
              </label>
              <Button type="submit" isLoading={isLoading} loadingText="Cadastrando...">
                Cadastrar clínica
              </Button>
            </div>
          </form>
        </Card>
        <p className="text-center mt-6 text-slate-600">
          Já tem conta? <Link href="/login">Entrar</Link>
        </p>
      </div>
    </div>
  );
}


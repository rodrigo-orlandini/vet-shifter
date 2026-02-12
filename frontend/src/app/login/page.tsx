'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Input from '@/components/Input';
import Button from '@/components/Button';
import Link from '@/components/Link';
import Card from '@/components/Card';
import { api, setToken } from '@/utils/api';

export default function LoginPage() {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setIsLoading(true);
    try {
      const res = await api<{ token: string; role: string; sub: string }>('/auth/login', {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      });
      setToken(res.token);
      router.push('/dashboard');
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Falha ao entrar');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-[#f0f9f7] px-4 py-8">
      <div className="w-full max-w-md">
        <div className="text-center mb-8">
          <h1 className="text-3xl font-bold text-teal-800">VetTroca</h1>
          <p className="text-slate-600 mt-1">Entre na sua conta</p>
        </div>
        <Card>
          <form onSubmit={handleSubmit}>
            <div className="space-y-6">
              {error && (
                <p className="text-sm text-red-600 bg-red-50 p-3 rounded-lg">{error}</p>
              )}
              <Input
                id="email"
                label="E-mail"
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
                placeholder="seu@email.com"
              />
              <Input
                id="password"
                label="Senha"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                placeholder="••••••••"
              />
              <Button type="submit" isLoading={isLoading} loadingText="Entrando...">
                Entrar
              </Button>
            </div>
          </form>
        </Card>
        <p className="text-center mt-6 text-slate-600">
          Não tem uma conta?{' '}
          <Link href="/signup">Clínica</Link>
          {' · '}
          <Link href="/signup/vet">Plantonista</Link>
        </p>
      </div>
    </div>
  );
}

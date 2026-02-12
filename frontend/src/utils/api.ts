const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export async function api<T>(
  path: string,
  options: RequestInit & { token?: string } = {}
): Promise<T> {
  const { token, ...init } = options;
  const headers: HeadersInit = {
    'Content-Type': 'application/json',
    ...(init.headers as Record<string, string>),
  };
  if (token) {
    (headers as Record<string, string>)['Authorization'] = `Bearer ${token}`;
  }
  const res = await fetch(`${API_BASE}${path}`, { ...init, headers });
  const data = await res.json().catch(() => ({}));
  if (!res.ok) {
    throw new Error((data as { error?: string }).error || res.statusText);
  }
  return data as T;
}

export function getToken(): string | null {
  if (typeof window === 'undefined') return null;
  return localStorage.getItem('vet_troca_token');
}

export function setToken(token: string): void {
  localStorage.setItem('vet_troca_token', token);
}

export function clearToken(): void {
  localStorage.removeItem('vet_troca_token');
}

export type LoginResponse = { token: string; role: string; sub: string };

export interface ViaCepResponse {
  cep?: string;
  logradouro?: string;
  complemento?: string;
  bairro?: string;
  localidade?: string;
  uf?: string;
  erro?: boolean;
}

/**
 * Busca endereço pelo CEP (8 dígitos). Retorna null se inválido ou não encontrado.
 */
export async function fetchAddressByCep(rawCep: string): Promise<{
  street: string;
  city: string;
  state: string;
} | null> {
  const digits = rawCep.replace(/\D/g, "");
  if (digits.length !== 8) return null;

  const res = await fetch(`https://viacep.com.br/ws/${digits}/json/`);
  if (!res.ok) return null;

  const data = (await res.json()) as ViaCepResponse;
  if (!data || data.erro) return null;

  return {
    street: data.logradouro ?? "",
    city: data.localidade ?? "",
    state: data.uf ?? "",
  };
}

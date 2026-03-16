const EMAIL_REGEX = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

/** Remove caracteres não numéricos */
function digitsOnly(value: string): string {
  return value.replace(/\D/g, "");
}

/** CNPJ: 14 dígitos. Aceita com ou sem pontuação. */
export function isValidCnpj(value: string): boolean {
  const d = digitsOnly(value);
  return d.length === 14;
}

/** CPF: 11 dígitos. Aceita com ou sem pontuação. */
export function isValidCpf(value: string): boolean {
  const d = digitsOnly(value);
  return d.length === 11;
}

/** Email: formato básico. */
export function isValidEmail(value: string): boolean {
  return EMAIL_REGEX.test(value.trim());
}

/** Telefone BR: 10 ou 11 dígitos (DDD + número). */
export function isValidPhoneBr(value: string): boolean {
  const d = digitsOnly(value);
  return d.length === 10 || d.length === 11;
}

/** Senha: mínimo 8 caracteres. */
export function isValidPassword(value: string): boolean {
  return value.length >= 8;
}

/** Campo obrigatório (não vazio após trim). */
export function isRequired(value: string | undefined | null): boolean {
  return typeof value === "string" && value.trim().length > 0;
}

/** CRMV: número e estado preenchidos (estado 2 caracteres). */
export function isValidCrmv(number: string, state: string): boolean {
  return isRequired(number) && isRequired(state) && state.trim().length === 2;
}

export const validationMessages = {
  required: "Campo obrigatório.",
  email: "Informe um e-mail válido.",
  cnpj: "CNPJ deve ter 14 dígitos (com ou sem pontuação).",
  cpf: "CPF deve ter 11 dígitos (com ou sem pontuação).",
  phone: "Telefone deve ter 10 ou 11 dígitos (DDD + número).",
  password: "A senha deve ter no mínimo 8 caracteres.",
  passwordMatch: "As senhas não coincidem.",
  lgpd: "É necessário aceitar o uso dos dados conforme a LGPD.",
  specialties: "Selecione pelo menos uma especialidade.",
  crmv: "Informe número e UF do CRMV (2 letras).",
} as const;

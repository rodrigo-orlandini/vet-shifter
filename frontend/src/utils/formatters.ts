export function formatCNPJ(value: string): string {
  const numbers = value.replace(/\D/g, '');
  if (numbers.length <= 14) {
    return numbers
      .replace(/^(\d{2})(\d)/, '$1.$2')
      .replace(/^(\d{2})\.(\d{3})(\d)/, '$1.$2.$3')
      .replace(/\.(\d{3})(\d)/, '.$1/$2')
      .replace(/(\d{4})(\d)/, '$1-$2');
  }
  return value;
}

export function formatPhone(value: string): string {
  const numbers = value.replace(/\D/g, '');
  if (numbers.length <= 11) {
    return numbers
      .replace(/^(\d{2})(\d)/, '($1) $2')
      .replace(/(\d{5})(\d)/, '$1-$2');
  }
  return value;
}

export function formatCPF(value: string): string {
  const numbers = value.replace(/\D/g, '');
  if (numbers.length <= 11) {
    return numbers
      .replace(/^(\d{3})(\d)/, '$1.$2')
      .replace(/^(\d{3})\.(\d{3})(\d)/, '$1.$2.$3')
      .replace(/\.(\d{3})(\d)/, '.$1-$2');
  }
  return value;
}


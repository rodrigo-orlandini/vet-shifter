export function passwordRules(value: string) {
  return {
    minLen: value.length >= 8,
    upper: /[A-ZÀ-Ü]/.test(value),
    digit: /[0-9]/.test(value),
  };
}

export type PasswordStrength = "weak" | "medium" | "strong";

export function getPasswordStrength(value: string): PasswordStrength {
  const r = passwordRules(value);
  const score = [r.minLen, r.upper, r.digit].filter(Boolean).length;

  if (score <= 1) return "weak";
  if (score === 2) return "medium";
  
  return "strong";
}

export function meetsPasswordPolicy(value: string): boolean {
  const r = passwordRules(value);
  return r.minLen && r.upper && r.digit;
}

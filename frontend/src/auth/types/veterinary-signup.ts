export interface RegisterVeterinaryRequest {
  full_name: string;
  cpf: string;
  email: string;
  phone: string;
  crmv_number: string;
  crmv_state: string;
  specialties: string[];
  password: string;
  consent_lgpd: boolean;
}

export interface RegisterVeterinaryResponse {
  veterinary_id?: string;
}

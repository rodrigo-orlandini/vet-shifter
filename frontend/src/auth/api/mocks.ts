/* eslint-disable */
import type { RegisterVeterinaryRequest } from "../types/veterinary-signup";

const delay = (ms: number) => new Promise((r) => setTimeout(r, ms));

export const MOCK_USER_TYPE: "company_owner" | "shift_veterinary" = "company_owner";

export async function mockGetUserType(_email: string): Promise<{ user_type: string }> {
  await delay(300);
  return { user_type: MOCK_USER_TYPE };
}

export async function mockGetUserTypeNotFound(_email: string): Promise<never> {
  await delay(300);
  const err = new Error("NOT_FOUND") as Error & { response?: { status: number } };
  err.response = { status: 404 };
  throw err;
}

export async function mockLoginOwner(_body: {
  email: string;
  password: string;
  remember_me?: boolean;
}): Promise<{ access_token: string; expires_at: string }> {
  await delay(400);
  return {
    access_token: "mock-jwt-owner",
    expires_at: new Date(Date.now() + 86400 * 1000).toISOString(),
  };
}

export async function mockLoginVeterinary(_body: {
  email: string;
  password: string;
  remember_me?: boolean;
}): Promise<{ access_token: string; expires_at: string }> {
  await delay(400);
  return {
    access_token: "mock-jwt-veterinary",
    expires_at: new Date(Date.now() + 86400 * 1000).toISOString(),
  };
}

export async function mockRegisterCompany(_body: unknown): Promise<{ company_id: string }> {
  await delay(500);
  return { company_id: "mock-company-id" };
}

export async function mockRegisterVeterinary(
  _body: RegisterVeterinaryRequest
): Promise<{ veterinary_id: string }> {
  await delay(500);
  return { veterinary_id: "mock-veterinary-id" };
}

export async function mockForgotPassword(_body: { email: string }): Promise<Record<string, string>> {
  await delay(400);
  return { message: "If an account exists with this email, you will receive a password reset link." };
}

export async function mockResetPassword(_body: {
  token: string;
  new_password: string;
}): Promise<Record<string, string>> {
  await delay(400);
  return { message: "Password updated successfully" };
}

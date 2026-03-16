import {
  mockForgotPassword,
  mockGetUserType,
  mockLoginOwner,
  mockLoginVeterinary,
  mockRegisterCompany,
  mockRegisterVeterinary,
  mockResetPassword,
} from "./mocks";
import type { RegisterVeterinaryRequest } from "../types/veterinary-signup";
import type { InternalCompaniesInfrastructureControllersRegisterCompanyRequest } from "@/api/generated/api";

export const authApi = {
  getUserType: (email: string) => mockGetUserType(email),

  loginOwner: (body: { email: string; password: string; remember_me?: boolean }) =>
    mockLoginOwner(body),

  loginVeterinary: (body: { email: string; password: string; remember_me?: boolean }) =>
    mockLoginVeterinary(body),

  registerCompany: (body: InternalCompaniesInfrastructureControllersRegisterCompanyRequest) =>
    mockRegisterCompany(body),

  registerVeterinary: (body: RegisterVeterinaryRequest) => mockRegisterVeterinary(body),

  forgotPassword: (body: { email: string }) => mockForgotPassword(body),

  resetPassword: (body: { token: string; new_password: string }) => mockResetPassword(body),
};

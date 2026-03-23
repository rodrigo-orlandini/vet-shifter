"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { StepLayout } from "@/components/StepLayout";
import { AuthFooterLinks } from "../../components/AuthFooterLinks";
import { CompanyStep1Form } from "./CompanyStep1Form";
import { CompanyStep2Form } from "./CompanyStep2Form";
import { CompanyStep3Form } from "./CompanyStep3Form";
import { AuthenticationService } from "@/auth/api";
import { useToast } from "@/components/toast/ToastProvider";
import type { InternalCompaniesInfrastructureControllersRegisterCompanyRequest } from "@/api/generated/api";
import {
  isRequired,
  isValidCnpj,
  isValidEmail,
  isValidPassword,
  isValidPhoneBr,
  validationMessages,
} from "@/lib/validation";
import { getBackendErrorMessage } from "@/lib/backendErrorMessage";

const initialForm: InternalCompaniesInfrastructureControllersRegisterCompanyRequest = {
  cnpj: "",
  company_name: "",
  owner_name: "",
  email: "",
  phone: "",
  password: "",
  consent_lgpd: false,
  street: "",
  number: "",
  city: "",
  state: "",
  zip_code: "",
};

type FieldErrors = Partial<Record<keyof typeof initialForm, string>>;

export default function CompanySignUpPage() {
  const router = useRouter();
  const { pushToast } = useToast();
  const [step, setStep] = useState(1);
  const [form, setForm] = useState(initialForm);
  const [fieldErrors, setFieldErrors] = useState<FieldErrors>({});
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const totalSteps = 3;
  const isFirstStep = step === 1;
  const isLastStep = step === totalSteps;

  const update = (partial: Partial<typeof form>) => {
    setForm((prev: typeof form) => ({ ...prev, ...partial }));
    setFieldErrors((prev) => {
      const next = { ...prev };
      Object.keys(partial).forEach((k) => delete next[k as keyof FieldErrors]);
      return next;
    });
  };

  const setStep1Errors = (): boolean => {
    const err: FieldErrors = {};
    if (!isRequired(form.cnpj)) err.cnpj = validationMessages.required;
    else if (!isValidCnpj(form.cnpj)) err.cnpj = validationMessages.cnpj;
    if (!isRequired(form.company_name)) err.company_name = "Informe a razão social da empresa.";
    setFieldErrors(err);
    return Object.keys(err).length === 0;
  };

  const setStep2Errors = (): boolean => {
    const err: FieldErrors = {};
    if (!isRequired(form.owner_name)) err.owner_name = "Informe o nome do responsável.";
    if (!isRequired(form.email)) err.email = validationMessages.required;
    else if (!isValidEmail(form.email)) err.email = validationMessages.email;
    if (!isRequired(form.phone)) err.phone = validationMessages.required;
    else if (!isValidPhoneBr(form.phone)) err.phone = validationMessages.phone;
    setFieldErrors(err);
    return Object.keys(err).length === 0;
  };

  const setStep3Errors = (): boolean => {
    const err: FieldErrors = {};
    if (!isRequired(form.password)) err.password = validationMessages.required;
    else if (!isValidPassword(form.password)) err.password = validationMessages.password;
    if (!form.consent_lgpd) err.consent_lgpd = validationMessages.lgpd;
    setFieldErrors(err);
    return Object.keys(err).length === 0;
  };

  const handleNext = () => {
    setError(null);
    if (step === 1 && !setStep1Errors()) return;
    if (step === 2 && !setStep2Errors()) return;
    if (step < totalSteps) setStep((s) => s + 1);
  };

  const handleBack = () => {
    setError(null);
    setFieldErrors({});
    if (step > 1) setStep((s) => s - 1);
  };

  const handleSubmit = async () => {
    setError(null);
    if (!setStep3Errors()) return;
    setSubmitting(true);

    // Backend value-objects expect digits-only values.
    const payload: InternalCompaniesInfrastructureControllersRegisterCompanyRequest = {
      ...form,
      cnpj: form.cnpj.replace(/\D/g, ""),
      phone: form.phone.replace(/\D/g, ""),
      zip_code: form.zip_code?.replace(/\D/g, "") ?? "",
    };

    try {
      await AuthenticationService.registerCompany(payload);
      router.push("/login?registered=company");
    } catch (e) {
      const message = getBackendErrorMessage(e);
      pushToast({ tone: "error", message });
      setError(message);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="overflow-hidden rounded-2xl border border-neutral-200/80 bg-white shadow-xl shadow-neutral-200/50">
      <div className="border-t-4 border-emerald-500 bg-linear-to-r from-emerald-500/5 to-teal-500/5 px-8 pt-8 pb-4">
        <h1 className="mb-1 text-2xl font-bold tracking-tight text-neutral-900">Cadastro de empresa</h1>
        <p className="text-sm text-neutral-600">Cadastre sua clínica para encontrar plantonistas.</p>
      </div>
      <div className="p-8 pt-6">

      <StepLayout
        currentStep={step}
        totalSteps={totalSteps}
        stepLabels={["Empresa", "Contato", "Segurança"]}
        onBack={handleBack}
        onNext={handleNext}
        onSubmit={handleSubmit}
        isFirstStep={isFirstStep}
        isLastStep={isLastStep}
        submitLabel="Criar conta"
        loading={submitting}
        submitDisabled={!form.consent_lgpd}
      >
        {step === 1 && <CompanyStep1Form form={form} fieldErrors={fieldErrors} update={update} />}
        {step === 2 && <CompanyStep2Form form={form} fieldErrors={fieldErrors} update={update} />}
        {step === 3 && <CompanyStep3Form form={form} fieldErrors={fieldErrors} update={update} />}
      </StepLayout>

      {error && (
        <p className="mt-4 text-sm text-red-600" role="alert">
          {error}
        </p>
      )}

      <AuthFooterLinks variant="company" />
      </div>
    </div>
  );
}

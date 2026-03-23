"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { StepLayout } from "@/components/StepLayout";
import { AuthFooterLinks } from "../../components/AuthFooterLinks";
import { VeterinaryStep1Form } from "./VeterinaryStep1Form";
import { VeterinaryStep2Form } from "./VeterinaryStep2Form";
import { VeterinaryStep3Form } from "./VeterinaryStep3Form";
import { getVetShifterAPI, type ControllersRegisterShiftVeterinaryRequest } from "@/api/generated/api";
import { useToast } from "@/components/toast/ToastProvider";
import {
  isRequired,
  isValidCpf,
  isValidCrmv,
  isValidEmail,
  isValidPassword,
  isValidPhoneBr,
  validationMessages,
} from "@/lib/validation";
import { getBackendErrorMessage } from "@/lib/backendErrorMessage";

const api = getVetShifterAPI();

const initialForm: ControllersRegisterShiftVeterinaryRequest = {
  full_name: "",
  cpf: "",
  email: "",
  phone: "",
  crmv_number: "",
  crmv_state: "",
  specialties: [],
  password: "",
  consent_lgpd: false,
};

type FieldErrors = Partial<Record<keyof ControllersRegisterShiftVeterinaryRequest | "specialties", string>>;

export default function VeterinarySignUpPage() {
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

  const update = (partial: Partial<ControllersRegisterShiftVeterinaryRequest>) => {
    setForm((prev) => ({ ...prev, ...partial }));
    setFieldErrors((prev) => {
      const next = { ...prev };
      Object.keys(partial).forEach((k) => delete next[k as keyof FieldErrors]);
      return next;
    });
  };

  const toggleSpecialty = (value: string) => {
    setForm((prev) => ({
      ...prev,
      specialties: prev.specialties.includes(value)
        ? prev.specialties.filter((s) => s !== value)
        : [...prev.specialties, value],
    }));
    if (fieldErrors.specialties) setFieldErrors((prev) => ({ ...prev, specialties: undefined }));
  };

  const setStep1Errors = (): boolean => {
    const err: FieldErrors = {};
    if (!isRequired(form.full_name)) err.full_name = "Informe seu nome completo.";
    if (!isRequired(form.cpf)) err.cpf = validationMessages.required;
    else if (!isValidCpf(form.cpf)) err.cpf = validationMessages.cpf;
    if (!isRequired(form.email)) err.email = validationMessages.required;
    else if (!isValidEmail(form.email)) err.email = validationMessages.email;
    if (!isRequired(form.phone)) err.phone = validationMessages.required;
    else if (!isValidPhoneBr(form.phone)) err.phone = validationMessages.phone;
    setFieldErrors(err);
    return Object.keys(err).length === 0;
  };

  const setStep2Errors = (): boolean => {
    const err: FieldErrors = {};
    if (!isValidCrmv(form.crmv_number, form.crmv_state)) err.crmv_number = validationMessages.crmv;
    if (!form.specialties.length) err.specialties = validationMessages.specialties;
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

    const payload: ControllersRegisterShiftVeterinaryRequest = {
      ...form,
      cpf: form.cpf.replace(/\D/g, ""),
      phone: form.phone.replace(/\D/g, ""),
    };

    try {
      await api.postVeterinaries(payload);

      pushToast({ tone: "success", message: "Cadastro realizado com sucesso!" });

      const loginRes = await api.postAuthLoginVeterinary({
        email: form.email,
        password: form.password,
        remember_me: false,
      });

      if (loginRes?.access_token) {
        router.push("/dashboard/veterinary");
        return;
      }

      router.push("/login?registered=veterinary");
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
        <h1 className="mb-1 text-2xl font-bold tracking-tight text-neutral-900">Cadastro de veterinário</h1>
        <p className="text-sm text-neutral-600">Cadastre-se como plantonista para encontrar plantões.</p>
      </div>
      <div className="p-8 pt-6">

      <StepLayout
        currentStep={step}
        totalSteps={totalSteps}
        stepLabels={["Dados pessoais", "Profissional", "Segurança"]}
        onBack={handleBack}
        onNext={handleNext}
        onSubmit={handleSubmit}
        isFirstStep={isFirstStep}
        isLastStep={isLastStep}
        submitLabel="Criar conta"
        loading={submitting}
        submitDisabled={!form.consent_lgpd}
      >
        {step === 1 && <VeterinaryStep1Form form={form} fieldErrors={fieldErrors} update={update} />}
        {step === 2 && (
          <VeterinaryStep2Form
            form={form}
            fieldErrors={fieldErrors}
            update={update}
            toggleSpecialty={toggleSpecialty}
          />
        )}
        {step === 3 && <VeterinaryStep3Form form={form} fieldErrors={fieldErrors} update={update} />}
      </StepLayout>

      {error && (
        <p className="mt-4 text-sm text-red-600" role="alert">
          {error}
        </p>
      )}

      <AuthFooterLinks variant="veterinary" />
      </div>
    </div>
  );
}

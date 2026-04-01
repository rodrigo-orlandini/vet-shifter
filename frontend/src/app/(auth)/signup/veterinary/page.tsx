"use client";

import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { StepLayout } from "@/components/StepLayout";
import { StepIndicator } from "@/components/StepIndicator";
import { AuthCard } from "@/components/auth/AuthCard";
import { Button } from "@/components/Button";
import { VeterinaryStep1Form } from "./VeterinaryStep1Form";
import { VeterinaryStep2Form } from "./VeterinaryStep2Form";
import { VeterinaryStep3Form, type VetDocKey } from "./VeterinaryStep3Form";
import { getVetShifterAPI, type ControllersRegisterShiftVeterinaryRequest } from "@/api/generated/api";
import { useToast } from "@/components/toast/ToastProvider";
import {
  isRequired,
  isValidCpf,
  isValidCrmv,
  isValidEmail,
  isValidPhoneBr,
  validationMessages,
} from "@/lib/validation";
import { meetsPasswordPolicy } from "@/lib/passwordPolicy";
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

type FieldErrors = Partial<
  Record<keyof ControllersRegisterShiftVeterinaryRequest | "specialties" | "confirmPassword", string>
>;

export default function VeterinarySignUpPage() {
  const { pushToast } = useToast();
  const router = useRouter();
  const [step, setStep] = useState(1);
  const [form, setForm] = useState(initialForm);
  const [confirmPassword, setConfirmPassword] = useState("");
  const [fieldErrors, setFieldErrors] = useState<FieldErrors>({});
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [submitted, setSubmitted] = useState(false);
  const [postLoginTarget, setPostLoginTarget] = useState<"dashboard" | "profile" | null>(null);
  const [docs, setDocs] = useState<Record<VetDocKey, File | null>>({
    idDoc: null,
    crmvFront: null,
    crmvBack: null,
    diploma: null,
  });

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

  const onConfirmPasswordChange = (value: string) => {
    setConfirmPassword(value);
    setFieldErrors((prev) => ({ ...prev, confirmPassword: undefined }));
  };

  const setDoc = (key: VetDocKey, file: File | null) => {
    setDocs((prev) => ({ ...prev, [key]: file }));
  };

  const toggleSpecialty = (value: string) => {
    setForm((prev) => ({
      ...prev,
      specialties: prev.specialties.includes(value)
        ? prev.specialties.filter((s) => s !== value)
        : [...prev.specialties, value],
    }));
    if (fieldErrors.specialties) {
      setFieldErrors((prev) => ({ ...prev, specialties: undefined }));
    }
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

    if (!isRequired(form.password)) err.password = validationMessages.required;
    else if (!meetsPasswordPolicy(form.password)) err.password = validationMessages.passwordPolicy;

    if (form.password !== confirmPassword) err.confirmPassword = validationMessages.passwordMatch;

    if (!form.consent_lgpd) err.consent_lgpd = validationMessages.lgpd;

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
    setSubmitting(true);

    const payload: ControllersRegisterShiftVeterinaryRequest = {
      ...form,
      cpf: form.cpf.replace(/\D/g, ""),
      phone: form.phone.replace(/\D/g, ""),
    };

    try {
      await api.postVeterinaries(payload);
      pushToast({ tone: "success", message: "Cadastro enviado com sucesso!" });
      setSubmitted(true);
    } catch (e) {
      const message = getBackendErrorMessage(e);
      pushToast({ tone: "error", message });
      setError(message);
    } finally {
      setSubmitting(false);
    }
  };

  const loginAndRedirect = async (target: "dashboard" | "profile") => {
    setError(null);
    setPostLoginTarget(target);

    try {
      await api.postAuthLoginVeterinary({ email: form.email, password: form.password, remember_me: false });
      router.push(target === "dashboard" ? "/dashboard/veterinary" : "/profile/veterinary");
    } catch (e) {
      const message = getBackendErrorMessage(e);
      pushToast({ tone: "error", message });
      setError(message);
      setPostLoginTarget(null);
    }
  };

  const subtitles: Record<number, string> = {
    1: "Etapa 1 de 3 — Dados pessoais",
    2: "Etapa 2 de 3 — Formação profissional",
    3: "Etapa 3 de 3 — Documentos para verificação",
  };

  if (submitted) {
    return (
      <div className="flex w-full flex-col items-center">
        <AuthCard className="w-full p-8 text-center sm:max-w-[480px] sm:p-12">
          <div className="mx-auto mb-6 flex h-[72px] w-[72px] items-center justify-center rounded-full bg-[#E6F9F0]">
            <svg className="h-8 w-8 text-[#38A169]" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2.5}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          <h1 className="text-2xl font-bold text-[#18181B]">Cadastro enviado com sucesso!</h1>
          <p className="mt-3 text-sm leading-relaxed text-[#71717A]">
            Seus documentos estão em análise. Assim que aprovados, você terá acesso completo à
            plataforma. Verifique seu e-mail para mais informações.
          </p>
          <div className="mt-6 flex flex-col gap-4">
            <Button
              type="button"
              className="w-full"
              loading={postLoginTarget === "dashboard"}
              disabled={postLoginTarget !== null}
              onClick={() => loginAndRedirect("dashboard")}
            >
              Ir para o painel
            </Button>

            <button
              type="button"
              className="text-sm font-medium text-[#2A9D8F] hover:underline disabled:cursor-not-allowed disabled:opacity-70"
              disabled={postLoginTarget !== null}
              onClick={() => loginAndRedirect("profile")}
            >
              Completar perfil enquanto aguardo
            </button>
          </div>
        </AuthCard>
      </div>
    );
  }

  return (
    <div className="flex w-full flex-col md:gap-8">
      <StepIndicator
        currentStep={step}
        totalSteps={totalSteps}
        stepLabels={["Dados Pessoais", "Formação Profissional", "Documentos"]}
      />

      <AuthCard>
        <div className="p-5 sm:p-10">
          <h1 className="text-xl font-bold text-[#18181B] sm:text-2xl">
            Cadastro de Veterinário Plantonista
          </h1>
          <p className="mt-1 text-sm text-[#71717A]">{subtitles[step]}</p>

          <div className="mt-6">
            <StepLayout
              onBack={handleBack}
              onNext={handleNext}
              onSubmit={handleSubmit}
              isFirstStep={isFirstStep}
              isLastStep={isLastStep}
              nextLabel="Próximo"
              submitLabel="Finalizar cadastro"
              loading={submitting}
              submitDisabled={false}
            >
              {step === 1 && (
                <VeterinaryStep1Form
                  form={form}
                  fieldErrors={fieldErrors}
                  update={update}
                  confirmPassword={confirmPassword}
                  onConfirmPasswordChange={onConfirmPasswordChange}
                />
              )}
              {step === 2 && (
                <VeterinaryStep2Form
                  form={form}
                  fieldErrors={fieldErrors}
                  update={update}
                  toggleSpecialty={toggleSpecialty}
                />
              )}
              {step === 3 && (
                <VeterinaryStep3Form files={docs} onFile={setDoc} onSkipUploads={handleSubmit} />
              )}
            </StepLayout>

            {error && (
              <p className="mt-4 text-sm text-[#E53E3E]" role="alert">
                {error}
              </p>
            )}
          </div>

          {step < 3 && (
            <p className="mt-6 text-center text-sm text-[#71717A]">
              Já tem uma conta?{" "}
              <Link href="/login" className="font-medium text-[#2A9D8F] hover:underline">
                Entrar
              </Link>
            </p>
          )}
        </div>
      </AuthCard>
    </div>
  );
}

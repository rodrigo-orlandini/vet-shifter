"use client";

import {
  IonButton,
  IonContent,
  IonHeader,
  IonInput,
  IonItem,
  IonLabel,
  IonPage,
  IonTitle,
  IonToolbar,
} from "@ionic/react";
import { useRouter, useSearchParams } from "next/navigation";
import { Suspense, useState } from "react";
import { axiosInstance } from "@/api/axios-instance";

function ResetPasswordForm() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const tokenFromUrl = searchParams.get("token") ?? "";
  const [token, setToken] = useState(tokenFromUrl);
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError("");
    if (newPassword.length < 8) {
      setError("A senha deve ter no mínimo 8 caracteres.");
      return;
    }
    if (newPassword !== confirmPassword) {
      setError("As senhas não coincidem.");
      return;
    }
    if (!token.trim()) {
      setError("Token inválido ou ausente.");
      return;
    }
    setLoading(true);
    try {
      await axiosInstance.post("/auth/reset-password", {
        token: token.trim(),
        new_password: newPassword,
      });
      setSuccess(true);
    } catch (err: unknown) {
      const res = err && typeof err === "object" && "response" in err ? (err as { response?: { data?: { error?: string } } }).response : undefined;
      setError(res?.data?.error ?? "Link inválido ou expirado. Solicite um novo.");
    } finally {
      setLoading(false);
    }
  }

  if (success) {
    return (
      <IonPage>
        <IonHeader>
          <IonToolbar>
            <IonTitle>Senha redefinida</IonTitle>
          </IonToolbar>
        </IonHeader>
        <IonContent className="ion-padding">
          <p className="ion-padding">Sua senha foi alterada com sucesso.</p>
          <IonButton expand="block" onClick={() => router.push("/login")}>
            Ir para o login
          </IonButton>
        </IonContent>
      </IonPage>
    );
  }

  return (
    <IonPage>
      <IonHeader>
        <IonToolbar>
          <IonTitle>Redefinir senha</IonTitle>
        </IonToolbar>
      </IonHeader>
      <IonContent className="ion-padding">
        <form onSubmit={handleSubmit} className="ion-padding">
          {!tokenFromUrl ? (
            <IonItem>
              <IonLabel position="stacked">Token (do e-mail)</IonLabel>
              <IonInput
                type="text"
                value={token}
                onIonInput={(e) => setToken(e.detail.value ?? "")}
                placeholder="Cole o token recebido por e-mail"
              />
            </IonItem>
          ) : null}
          <IonItem>
            <IonLabel position="stacked">Nova senha</IonLabel>
            <IonInput
              type="password"
              value={newPassword}
              onIonInput={(e) => setNewPassword(e.detail.value ?? "")}
              required
              placeholder="Mínimo 8 caracteres"
            />
          </IonItem>
          <IonItem>
            <IonLabel position="stacked">Confirmar senha</IonLabel>
            <IonInput
              type="password"
              value={confirmPassword}
              onIonInput={(e) => setConfirmPassword(e.detail.value ?? "")}
              required
              placeholder="Repita a senha"
            />
          </IonItem>
          {error ? (
            <p className="text-red-600 ion-padding-start ion-padding-end">
              {error}
            </p>
          ) : null}
          <IonButton
            type="submit"
            expand="block"
            className="ion-margin-top"
            disabled={loading}
          >
            {loading ? "Salvando..." : "Redefinir senha"}
          </IonButton>
          <IonButton
            type="button"
            expand="block"
            fill="clear"
            onClick={() => router.push("/login")}
            className="ion-margin-top"
          >
            Voltar ao login
          </IonButton>
        </form>
      </IonContent>
    </IonPage>
  );
}

export default function ResetPasswordPage() {
  return (
    <Suspense fallback={
      <IonPage>
        <IonHeader><IonToolbar><IonTitle>Redefinir senha</IonTitle></IonToolbar></IonHeader>
        <IonContent className="ion-padding"><p>Carregando...</p></IonContent>
      </IonPage>
    }>
      <ResetPasswordForm />
    </Suspense>
  );
}

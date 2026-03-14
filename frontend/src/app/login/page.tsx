"use client";

import {
  IonButton,
  IonCheckbox,
  IonContent,
  IonHeader,
  IonInput,
  IonItem,
  IonLabel,
  IonPage,
  IonTitle,
  IonToolbar,
} from "@ionic/react";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { setToken } from "@/lib/auth-storage";
import { axiosInstance } from "@/api/axios-instance";

export default function LoginPage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [rememberMe, setRememberMe] = useState(false);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError("");
    setLoading(true);
    try {
      const { data } = await axiosInstance.post<{
        access_token: string;
        expires_at: string;
        user: { id: string; email: string; type: string };
      }>("/auth/login", { email, password, remember_me: rememberMe });
      await setToken(
        data.access_token,
        { id: data.user.id, email: data.user.email, type: data.user.type },
        data.expires_at,
        rememberMe
      );
      router.push("/");
    } catch (err: unknown) {
      const res = err && typeof err === "object" && "response" in err ? (err as { response?: { status?: number; data?: { error?: string } } }).response : undefined;
      if (res?.status === 401) {
        setError("E-mail ou senha inválidos.");
      } else {
        setError(res?.data?.error ?? "Erro ao fazer login. Tente novamente.");
      }
    } finally {
      setLoading(false);
    }
  }

  return (
    <IonPage>
      <IonHeader>
        <IonToolbar>
          <IonTitle>Entrar</IonTitle>
        </IonToolbar>
      </IonHeader>
      <IonContent className="ion-padding">
        <form onSubmit={handleSubmit} className="ion-padding">
          <IonItem>
            <IonLabel position="stacked">E-mail</IonLabel>
            <IonInput
              type="email"
              value={email}
              onIonInput={(e) => setEmail(e.detail.value ?? "")}
              required
              placeholder="seu@email.com"
            />
          </IonItem>
          <IonItem>
            <IonLabel position="stacked">Senha</IonLabel>
            <IonInput
              type="password"
              value={password}
              onIonInput={(e) => setPassword(e.detail.value ?? "")}
              required
              placeholder="Mínimo 8 caracteres"
            />
          </IonItem>
          <IonItem lines="none">
            <IonCheckbox
              checked={rememberMe}
              onIonChange={(e) => setRememberMe(e.detail.checked)}
            />
            <IonLabel>Manter me conectado</IonLabel>
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
            {loading ? "Entrando..." : "Entrar"}
          </IonButton>
          <IonButton
            type="button"
            expand="block"
            fill="clear"
            onClick={() => router.push("/forgot-password")}
            className="ion-margin-top"
          >
            Esqueci minha senha
          </IonButton>
        </form>
      </IonContent>
    </IonPage>
  );
}

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
import { useRouter } from "next/navigation";
import { useState } from "react";
import { axiosInstance } from "@/api/axios-instance";

export default function ForgotPasswordPage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [sent, setSent] = useState(false);
  const [loading, setLoading] = useState(false);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setLoading(true);
    try {
      await axiosInstance.post("/auth/forgot-password", { email });
      setSent(true);
    } catch {
      setSent(true);
    } finally {
      setLoading(false);
    }
  }

  return (
    <IonPage>
      <IonHeader>
        <IonToolbar>
          <IonTitle>Esqueci minha senha</IonTitle>
        </IonToolbar>
      </IonHeader>
      <IonContent className="ion-padding">
        {sent ? (
          <div className="ion-padding">
            <p>
              Se existir uma conta com este e-mail, você receberá um link para
              redefinir sua senha.
            </p>
            <IonButton expand="block" onClick={() => router.push("/login")}>
              Voltar ao login
            </IonButton>
          </div>
        ) : (
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
            <IonButton
              type="submit"
              expand="block"
              className="ion-margin-top"
              disabled={loading}
            >
              {loading ? "Enviando..." : "Enviar link"}
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
        )}
      </IonContent>
    </IonPage>
  );
}

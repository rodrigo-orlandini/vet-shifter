"use client";

import { IonButton, IonContent, IonHeader, IonPage, IonTitle, IonToolbar } from "@ionic/react";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { clearAuth, getUser } from "@/lib/auth-storage";
import { axiosInstance } from "@/api/axios-instance";

export default function Home() {
  const router = useRouter();
  const [user, setUser] = useState<{ id: string; email: string; type: string } | null>(null);

  useEffect(() => {
    getUser().then(setUser);
  }, []);

  async function handleLogout() {
    try {
      await axiosInstance.post("/auth/logout");
    } catch {
      // ignore
    }
    await clearAuth();
    router.push("/login");
  }

  return (
    <IonPage>
      <IonHeader>
        <IonToolbar>
          <IonTitle>Vet Shifter</IonTitle>
          {user ? (
            <IonButton slot="end" fill="clear" onClick={handleLogout}>
              Sair
            </IonButton>
          ) : (
            <IonButton slot="end" fill="clear" onClick={() => router.push("/login")}>
              Entrar
            </IonButton>
          )}
        </IonToolbar>
      </IonHeader>
      <IonContent className="ion-padding">
        <div className="flex min-h-[80vh] flex-col items-center justify-center">
          <h1 className="text-2xl font-bold">Vet Shifter</h1>
          {user ? (
            <p className="mt-2 text-gray-600">
              Olá, {user.email} ({user.type})
            </p>
          ) : (
            <p className="mt-2 text-gray-600">Conecte clínicas e plantonistas.</p>
          )}
        </div>
      </IonContent>
    </IonPage>
  );
}

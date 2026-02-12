import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "VetTroca — Plantões e clínicas veterinárias",
  description: "Conecte clínicas e veterinários plantonistas de forma segura",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="pt-BR">
      <body className="min-h-screen bg-[var(--color-background)] text-slate-800">{children}</body>
    </html>
  );
}

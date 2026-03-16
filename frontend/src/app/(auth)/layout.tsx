import Link from "next/link";
import { APP_NAME } from "@/app/config";

export default function AuthLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <div className="relative flex min-h-screen flex-col items-center justify-center overflow-hidden bg-linear-to-br from-emerald-50 via-white to-teal-50 px-4 py-8">
      <div
        className="absolute inset-0 opacity-30"
        style={{
          backgroundImage: `radial-gradient(circle at 1px 1px, rgb(0 0 0 / 0.06) 1px, transparent 0)`,
          backgroundSize: "24px 24px",
        }}
      />
      <div className="absolute left-1/4 top-1/4 h-72 w-72 rounded-full bg-emerald-200/40 blur-3xl" />
      <div className="absolute bottom-1/4 right-1/4 h-96 w-96 rounded-full bg-teal-200/30 blur-3xl" />
      <div className="relative z-10 flex w-full max-w-md flex-col items-center gap-8">
        <Link
          href="/login"
          className="text-2xl font-bold tracking-tight text-neutral-800 hover:text-emerald-600"
        >
          {APP_NAME}
        </Link>
        {children}
      </div>
    </div>
  );
}

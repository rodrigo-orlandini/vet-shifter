import { ToastProvider } from "@/components/toast/ToastProvider";
import { AuthTopNav } from "@/components/auth/AuthTopNav";

export default function AuthLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <ToastProvider>
      <div className="flex min-h-screen flex-col bg-page text-ink-body">
        <AuthTopNav />
        <main className="flex flex-1 flex-col px-4 py-6 sm:px-6 sm:py-10">
          <div className="mx-auto w-full max-w-[560px] flex-1">{children}</div>
        </main>
      </div>
    </ToastProvider>
  );
}

import Link from "next/link";
import { APP_NAME } from "@/app/config";
import { AuthNavLink } from "./AuthNavLink";

export function AuthTopNav() {
  return (
    <header className="border-b border-edge bg-surface">
      <div className="flex h-[52px] w-full items-center justify-between px-5 sm:h-16 sm:px-10">
        <Link
          href="/login"
          className="text-[18px] font-bold text-primary transition-colors hover:text-primary-hover focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary/30 sm:text-[22px]"
        >
          {APP_NAME}
        </Link>
        <AuthNavLink />
      </div>
    </header>
  );
}

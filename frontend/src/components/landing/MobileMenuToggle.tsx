"use client";

import { useState } from "react";
import Link from "next/link";
import { MenuIcon } from "@/components/icons/MenuIcon";
import { XIcon } from "@/components/icons/XIcon";

export function MobileMenuToggle() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="lg:hidden">
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="p-1 text-ink-body"
        aria-label={isOpen ? "Fechar menu" : "Abrir menu"}
      >
        {isOpen ? (
          <XIcon className="h-6 w-6" />
        ) : (
          <MenuIcon className="h-6 w-6" />
        )}
      </button>

      {isOpen && (
        <div className="fixed inset-x-0 top-[60px] z-50 border-b border-edge-input bg-surface px-5 py-5 shadow-lg">
          <nav className="flex flex-col gap-4 border-b border-edge-alt pb-5">
            <a
              href="#como-funciona"
              className="text-sm text-ink-muted transition-colors hover:text-ink-body"
              onClick={() => setIsOpen(false)}
            >
              Como funciona
            </a>
            <a
              href="#para-clinicas"
              className="text-sm text-ink-muted transition-colors hover:text-ink-body"
              onClick={() => setIsOpen(false)}
            >
              Para clínicas
            </a>
            <a
              href="#para-veterinarios"
              className="text-sm text-ink-muted transition-colors hover:text-ink-body"
              onClick={() => setIsOpen(false)}
            >
              Para veterinários
            </a>
            <a
              href="#precos"
              className="text-sm text-ink-muted transition-colors hover:text-ink-body"
              onClick={() => setIsOpen(false)}
            >
              Preços
            </a>
          </nav>
          <div className="flex flex-col gap-3 pt-4">
            <Link
              href="/signup/company"
              className="flex items-center justify-center rounded-lg bg-primary py-3 text-sm font-semibold text-white transition-colors hover:bg-primary-hover"
              onClick={() => setIsOpen(false)}
            >
              Sou clínica
            </Link>
            <Link
              href="/signup/veterinary"
              className="flex items-center justify-center rounded-lg border-[1.5px] border-primary py-3 text-sm font-semibold text-primary transition-colors hover:bg-primary-light"
              onClick={() => setIsOpen(false)}
            >
              Sou veterinário
            </Link>
            <Link
              href="/login"
              className="flex items-center justify-center py-2 text-sm font-medium text-ink-muted transition-colors hover:text-ink-body"
              onClick={() => setIsOpen(false)}
            >
              Já tenho conta — Entrar
            </Link>
          </div>
        </div>
      )}
    </div>
  );
}

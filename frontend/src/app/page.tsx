import Link from "next/link";
import { APP_NAME } from "@/app/config";
import { MobileMenuToggle } from "@/components/landing/MobileMenuToggle";
import { ArrowRightIcon } from "@/components/icons/ArrowRightIcon";
import { PawPrintIcon } from "@/components/icons/PawPrintIcon";
import { CircleCheckIcon } from "@/components/icons/CircleCheckIcon";
import { UserPlusIcon } from "@/components/icons/UserPlusIcon";
import { SearchIcon } from "@/components/icons/SearchIcon";
import { CalendarCheckIcon } from "@/components/icons/CalendarCheckIcon";
import { ShieldCheckIcon } from "@/components/icons/ShieldCheckIcon";
import { BanknoteIcon } from "@/components/icons/BanknoteIcon";
import { TimerIcon } from "@/components/icons/TimerIcon";
import { TrendingUpIcon } from "@/components/icons/TrendingUpIcon";
import { StarIcon } from "@/components/icons/StarIcon";
import { InstagramIcon } from "@/components/icons/InstagramIcon";
import { LinkedinIcon } from "@/components/icons/LinkedinIcon";
import { TwitterXIcon } from "@/components/icons/TwitterXIcon";

export default function HomePage() {
  return (
    <div className="min-h-screen bg-page font-sans">
      <header className="sticky top-0 z-40 border-b border-edge-input bg-surface">
        <div className="mx-auto flex h-[60px] max-w-[1440px] items-center justify-between px-5 lg:h-[72px] lg:px-20">
          <Link href="/" className="flex items-center gap-2.5">
            <PawPrintIcon className="h-5 w-5 text-primary lg:h-7 lg:w-7" />
            <span className="text-base font-bold text-ink-body lg:text-xl">{APP_NAME}</span>
          </Link>

          <nav className="hidden items-center gap-10 lg:flex">
            <a href="#como-funciona" className="text-sm text-ink-muted transition-colors hover:text-ink-body">
              Como funciona
            </a>
            <a href="#para-clinicas" className="text-sm text-ink-muted transition-colors hover:text-ink-body">
              Para clínicas
            </a>
            <a href="#para-veterinarios" className="text-sm text-ink-muted transition-colors hover:text-ink-body">
              Para veterinários
            </a>
            <a href="#precos" className="text-sm text-ink-muted transition-colors hover:text-ink-body">
              Preços
            </a>
          </nav>

          <div className="hidden items-center gap-3 lg:flex">
            <Link
              href="/login"
              className="text-sm font-medium mr-2 text-ink-muted transition-colors hover:text-ink-body"
            >
              Entrar
            </Link>
            <Link
              href="/signup/company"
              className="rounded-lg bg-primary px-5 py-2.5 text-sm font-semibold text-white transition-colors hover:bg-primary-hover"
            >
              Sou clínica
            </Link>
            <Link
              href="/signup/veterinary"
              className="rounded-lg border-[1.5px] border-primary px-5 py-2.5 text-sm font-semibold text-primary transition-colors hover:bg-primary-light"
            >
              Sou veterinário
            </Link>
          </div>

          <MobileMenuToggle />
        </div>
      </header>

      <section className="bg-surface">
        <div className="mx-auto max-w-[1440px] px-5 py-10 lg:flex lg:items-center lg:gap-16 lg:px-20 lg:py-20">
          <div className="flex flex-col gap-4 lg:flex-1 lg:gap-7">
            <div className="flex w-fit items-center gap-2 rounded-full bg-primary-light px-3.5 py-1.5">
              <CircleCheckIcon className="h-4 w-4 text-primary" />
              <span className="text-xs font-semibold text-primary lg:text-[13px]">
                Conectando clínicas e profissionais
              </span>
            </div>

            <h1 className="text-[28px] font-extrabold leading-[1.2] text-ink-body lg:text-[52px] lg:leading-[1.15]">
              Encontre o veterinário plantonista ideal para sua clínica
            </h1>

            <p className="text-[15px] leading-normal text-ink-muted lg:text-lg lg:leading-[1.6]">
              O VetPlant conecta clínicas veterinárias com profissionais plantonistas qualificados de forma rápida, segura e confiável.
            </p>

            <div className="relative h-[220px] w-full overflow-hidden rounded-xl lg:hidden">
              <img
                src="/images/generated-1775040252249.png"
                alt="Veterinária plantonista"
                className="h-full w-full object-cover"
              />
            </div>

            <div className="flex flex-col gap-3 lg:flex-row lg:gap-4">
              <Link
                href="/signup/company"
                className="flex items-center justify-center rounded-lg bg-primary px-8 py-[14px] text-[15px] font-semibold text-white transition-colors hover:bg-primary-hover lg:py-4 lg:text-base"
              >
                Cadastrar minha clínica
              </Link>
              <Link
                href="/signup/veterinary"
                className="flex items-center justify-center rounded-lg border-2 border-primary px-8 py-[14px] text-[15px] font-semibold text-primary transition-colors hover:bg-primary-light lg:py-4 lg:text-base"
              >
                Quero ser plantonista
              </Link>
            </div>

            <div className="flex flex-col gap-1.5 lg:flex-row lg:items-center lg:gap-6">
              <span className="text-xs text-success lg:text-[13px]">✓ Verificação de documentos</span>
              <span className="text-xs text-success lg:text-[13px]">✓ Pagamento seguro</span>
              <span className="text-xs text-success lg:text-[13px]">✓ Avaliações reais</span>
            </div>
          </div>

          <div className="hidden shrink-0 lg:block">
            <div className="h-[500px] w-[560px] overflow-hidden rounded-3xl">
              <img
                src="/images/generated-1775040612254.png"
                alt="Veterinária plantonista"
                className="h-full w-full object-cover"
              />
            </div>
          </div>
        </div>
      </section>

      <section id="como-funciona" className="bg-page">
        <div className="mx-auto max-w-[1440px] px-5 py-10 lg:px-20 lg:py-20">
          <div className="mb-8 flex flex-col items-center gap-3 text-center lg:mb-12 lg:gap-4">
            <p className="text-[11px] font-bold uppercase tracking-[1.5px] text-primary lg:hidden">
              COMO FUNCIONA
            </p>
            <h2 className="hidden text-[40px] font-extrabold text-ink-body lg:block">
              Como funciona
            </h2>
            <h2 className="text-[22px] font-bold text-ink-body lg:hidden">
              Simples, rápido e seguro
            </h2>
            <p className="max-w-xl text-sm leading-normal text-ink-muted lg:text-lg">
              <span className="hidden lg:inline">Em apenas três passos simples, conecte clínicas e veterinários plantonistas.</span>
              <span className="lg:hidden">Em poucos passos você conecta sua clínica ao profissional ideal</span>
            </p>
          </div>

          <div className="flex flex-col gap-4 lg:flex-row lg:gap-6">
            <div className="flex flex-1 flex-col gap-3 rounded-xl bg-surface p-5 shadow-[0_2px_8px_rgba(0,0,0,0.08)] lg:gap-4 lg:rounded-2xl lg:p-8">
              <div className="flex items-center gap-3 lg:flex-col lg:items-start lg:gap-4">
                <span className="text-[28px] font-extrabold text-primary lg:text-4xl">01</span>
                <div className="rounded-xl bg-primary-light p-3 lg:rounded-xl">
                  <UserPlusIcon className="h-[26px] w-[26px] text-primary" />
                </div>
                <h3 className="text-[15px] font-bold text-ink-body lg:hidden">Crie seu perfil</h3>
              </div>
              <h3 className="hidden text-xl font-bold text-ink-body lg:block">Crie seu perfil</h3>
              <p className="text-[13px] leading-normal text-ink-muted lg:text-[15px] lg:leading-[1.6]">
                Cadastre sua clínica ou perfil de veterinário em minutos. Verificamos suas credenciais para garantir segurança.
              </p>
            </div>

            <div className="flex flex-1 flex-col gap-3 rounded-xl bg-surface p-5 shadow-[0_2px_8px_rgba(0,0,0,0.08)] lg:gap-4 lg:rounded-2xl lg:p-8">
              <div className="flex items-center gap-3 lg:flex-col lg:items-start lg:gap-4">
                <span className="text-[28px] font-extrabold text-primary lg:text-4xl">02</span>
                <div className="rounded-xl bg-primary-light p-3">
                  <SearchIcon className="h-[26px] w-[26px] text-primary" />
                </div>
                <h3 className="text-[15px] font-bold text-ink-body lg:hidden">Encontre o match perfeito</h3>
              </div>
              <h3 className="hidden text-xl font-bold text-ink-body lg:block">Encontre o match perfeito</h3>
              <p className="text-[13px] leading-normal text-ink-muted lg:text-[15px] lg:leading-[1.6]">
                Clínicas publicam plantões disponíveis. Veterinários filtram por especialidade, região e horário.
              </p>
            </div>

            <div className="flex flex-1 flex-col gap-3 rounded-xl bg-surface p-5 shadow-[0_2px_8px_rgba(0,0,0,0.08)] lg:gap-4 lg:rounded-2xl lg:p-8">
              <div className="flex items-center gap-3 lg:flex-col lg:items-start lg:gap-4">
                <span className="text-[28px] font-extrabold text-primary lg:text-4xl">03</span>
                <div className="rounded-xl bg-primary-light p-3">
                  <CalendarCheckIcon className="h-[26px] w-[26px] text-primary" />
                </div>
                <h3 className="text-[15px] font-bold text-ink-body lg:hidden">Confirme e trabalhe</h3>
              </div>
              <h3 className="hidden text-xl font-bold text-ink-body lg:block">Confirme e trabalhe</h3>
              <p className="text-[13px] leading-normal text-ink-muted lg:text-[15px] lg:leading-[1.6]">
                Confirme o plantão, assine o contrato digital e realize o atendimento. Pagamento processado com segurança.
              </p>
            </div>
          </div>
        </div>
      </section>

      <section id="para-clinicas" className="bg-primary">
        <div className="mx-auto max-w-[1440px] px-5 py-10 lg:flex lg:items-center lg:gap-16 lg:px-20 lg:py-20">
          <div className="hidden shrink-0 lg:block">
            <div className="h-[440px] w-[480px] overflow-hidden rounded-3xl">
              <img
                src="/images/generated-1775040656003.png"
                alt="Clínica veterinária"
                className="h-full w-full object-cover"
              />
            </div>
          </div>

          <div className="flex flex-col gap-5 lg:flex-1 lg:gap-7">
            <div className="w-fit rounded-full bg-white/15 px-3.5 py-1.5">
              <span className="text-[11px] font-bold uppercase tracking-[1.5px] text-white/80 lg:text-[13px] lg:font-semibold lg:normal-case lg:tracking-normal">
                Para clínicas
              </span>
            </div>

            <h2 className="text-[22px] font-bold leading-[1.3] text-white lg:text-4xl lg:font-extrabold">
              Gerencie seus plantões com facilidade
            </h2>

            <div className="flex flex-col gap-4 lg:gap-5">
              <div className="flex items-start gap-3.5">
                <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-white/15">
                  <CalendarCheckIcon className="h-5 w-5 text-white" />
                </div>
                <div className="flex flex-col gap-0.5">
                  <span className="text-sm font-semibold text-white lg:text-base">
                    Agendamento simplificado
                  </span>
                  <span className="text-xs leading-[1.4] text-white/70 lg:text-sm lg:leading-[1.6]">
                    Publique vagas e receba candidatos qualificados em minutos.
                  </span>
                </div>
              </div>

              <div className="flex items-start gap-3.5">
                <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-white/15">
                  <ShieldCheckIcon className="h-5 w-5 text-white" />
                </div>
                <div className="flex flex-col gap-0.5">
                  <span className="text-sm font-semibold text-white lg:text-base">
                    Profissionais verificados
                  </span>
                  <span className="text-xs leading-[1.4] text-white/70 lg:text-sm lg:leading-[1.6]">
                    Todos os veterinários têm CRMV verificado e histórico de avaliações.
                  </span>
                </div>
              </div>

              <div className="flex items-start gap-3.5">
                <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-white/15">
                  <BanknoteIcon className="h-5 w-5 text-white" />
                </div>
                <div className="flex flex-col gap-0.5">
                  <span className="text-sm font-semibold text-white lg:text-base">
                    Pagamento simplificado
                  </span>
                  <span className="text-xs leading-[1.4] text-white/70 lg:text-sm lg:leading-[1.6]">
                    Contratos e pagamentos 100% digitais, sem burocracia.
                  </span>
                </div>
              </div>
            </div>

            <div className="pt-2">
              <Link
                href="/signup/company"
                className="inline-flex items-center gap-2 rounded-lg bg-white px-6 py-3.5 text-sm font-semibold text-primary transition-colors hover:bg-primary-light lg:text-base"
              >
                Cadastrar minha clínica
                <ArrowRightIcon className="h-4 w-4" />
              </Link>
            </div>
          </div>
        </div>
      </section>

      <section id="para-veterinarios" className="bg-surface">
        <div className="mx-auto max-w-[1440px] px-5 py-10 lg:flex lg:items-center lg:gap-16 lg:px-20 lg:py-20">
          <div className="flex flex-col gap-5 lg:flex-1 lg:gap-7">
            <div className="w-fit rounded-full bg-primary-light px-3.5 py-1.5">
              <span className="text-[11px] font-bold uppercase tracking-[1.5px] text-primary lg:text-[13px] lg:font-semibold lg:normal-case lg:tracking-normal">
                Para veterinários
              </span>
            </div>

            <h2 className="text-[22px] font-bold leading-[1.3] text-ink-body lg:text-4xl lg:font-extrabold">
              Trabalhe quando e onde quiser
            </h2>

            <div className="flex flex-col gap-4 lg:gap-5">
              <div className="flex items-start gap-3.5">
                <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary-light">
                  <TimerIcon className="h-5 w-5 text-primary" />
                </div>
                <div className="flex flex-col gap-0.5">
                  <span className="text-sm font-semibold text-ink-body lg:text-base">
                    Flexibilidade total
                  </span>
                  <span className="text-xs leading-[1.4] text-ink-muted lg:text-sm lg:leading-[1.6]">
                    Escolha os plantões que se encaixam na sua agenda, sem compromissos fixos.
                  </span>
                </div>
              </div>

              <div className="flex items-start gap-3.5">
                <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary-light">
                  <TrendingUpIcon className="h-5 w-5 text-primary" />
                </div>
                <div className="flex flex-col gap-0.5">
                  <span className="text-sm font-semibold text-ink-body lg:text-base">
                    Renda extra garantida
                  </span>
                  <span className="text-xs leading-[1.4] text-ink-muted lg:text-sm lg:leading-[1.6]">
                    Defina sua tabela de honorários e receba pagamentos pontuais e seguros.
                  </span>
                </div>
              </div>

              <div className="flex items-start gap-3.5">
                <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-xl bg-primary-light">
                  <StarIcon className="h-5 w-5 text-primary" />
                </div>
                <div className="flex flex-col gap-0.5">
                  <span className="text-sm font-semibold text-ink-body lg:text-base">
                    Construa sua reputação
                  </span>
                  <span className="text-xs leading-[1.4] text-ink-muted lg:text-sm lg:leading-[1.6]">
                    Receba avaliações e construa um portfólio reconhecido no mercado.
                  </span>
                </div>
              </div>
            </div>

            <div className="pt-2">
              <Link
                href="/signup/veterinary"
                className="inline-flex items-center gap-2 rounded-lg bg-primary px-6 py-3.5 text-sm font-semibold text-white transition-colors hover:bg-primary-hover lg:text-base"
              >
                Quero ser plantonista
                <ArrowRightIcon className="h-4 w-4" />
              </Link>
            </div>
          </div>

          <div className="hidden shrink-0 lg:block">
            <div className="h-[440px] w-[480px] overflow-hidden rounded-3xl">
              <img
                src="/images/generated-1775040684994.png"
                alt="Veterinário plantonista"
                className="h-full w-full object-cover"
              />
            </div>
          </div>
        </div>
      </section>

      <section aria-hidden="true" className="hidden">
        <div className="mx-auto max-w-[1440px] px-5 py-10 lg:px-20 lg:py-20">
          <div className="mb-8 flex flex-col items-center gap-3 text-center lg:mb-12 lg:gap-4">
            <p className="text-[11px] font-bold uppercase tracking-[1.5px] text-primary">DEPOIMENTOS</p>
            <h2 className="text-[20px] font-bold leading-[1.3] text-ink-body lg:text-[40px] lg:font-extrabold">
              O que nossos usuários dizem
            </h2>
            <p className="text-sm text-ink-muted lg:text-lg">
              Centenas de clínicas e profissionais já confiam no VetPlant.
            </p>
          </div>

          <div className="flex flex-col gap-3 lg:flex-row lg:gap-6">
            <div className="flex flex-1 flex-col gap-3 rounded-xl bg-surface p-5 shadow-[0_2px_8px_rgba(0,0,0,0.06)] lg:rounded-2xl lg:p-8">
              <div className="flex gap-1">
                {Array.from({ length: 5 }).map((_, i) => (
                  <StarIcon key={i} className="h-3.5 w-3.5 fill-warning text-warning" />
                ))}
              </div>
              <p className="text-[13px] leading-normal text-ink-body">
                &ldquo;A VetPlant transformou como gerenciamos os plantões. Antes era uma dor de cabeça, agora encontramos profissionais qualificados em minutos.&rdquo;
              </p>
              <div className="flex items-center gap-2.5">
                <div className="flex h-9 w-9 items-center justify-center rounded-full bg-primary-light">
                  <span className="text-[13px] font-bold text-primary">AM</span>
                </div>
                <div>
                  <p className="text-[13px] font-semibold text-ink-body">Ana Martins</p>
                  <p className="text-[11px] text-ink-muted">Diretora, Clínica Pet Life • São Paulo</p>
                </div>
              </div>
            </div>

            <div className="flex flex-1 flex-col gap-3 rounded-xl bg-surface p-5 shadow-[0_2px_8px_rgba(0,0,0,0.06)] lg:rounded-2xl lg:p-8">
              <div className="flex gap-1">
                {Array.from({ length: 5 }).map((_, i) => (
                  <StarIcon key={i} className="h-3.5 w-3.5 fill-warning text-warning" />
                ))}
              </div>
              <p className="text-[13px] leading-normal text-ink-body">
                &ldquo;Faço plantões pelo VetPlant há 8 meses. A flexibilidade de agenda e os pagamentos sempre em dia mudaram minha vida profissional.&rdquo;
              </p>
              <div className="flex items-center gap-2.5">
                <div className="flex h-9 w-9 items-center justify-center rounded-full bg-[#E6F0FF]">
                  <span className="text-[13px] font-bold text-[#3B5FCC]">CR</span>
                </div>
                <div>
                  <p className="text-[13px] font-semibold text-ink-body">Carlos Rodrigues</p>
                  <p className="text-[11px] text-ink-muted">Veterinário Plantonista • Rio de Janeiro</p>
                </div>
              </div>
            </div>

            <div className="flex flex-1 flex-col gap-3 rounded-xl bg-surface p-5 shadow-[0_2px_8px_rgba(0,0,0,0.06)] lg:rounded-2xl lg:p-8">
              <div className="flex gap-1">
                {Array.from({ length: 5 }).map((_, i) => (
                  <StarIcon key={i} className="h-3.5 w-3.5 fill-warning text-warning" />
                ))}
              </div>
              <p className="text-[13px] leading-normal text-ink-body">
                &ldquo;Processo de cadastro rápido e suporte excelente. Já indicamos para outras clínicas da rede. Plataforma indispensável!&rdquo;
              </p>
              <div className="flex items-center gap-2.5">
                <div className="flex h-9 w-9 items-center justify-center rounded-full bg-warning-subtle">
                  <span className="text-[13px] font-bold text-warning-ink">FS</span>
                </div>
                <div>
                  <p className="text-[13px] font-semibold text-ink-body">Fernanda Silva</p>
                  <p className="text-[11px] text-ink-muted">Gerente, Hospital Veterinário Central • BH</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section id="precos" className="bg-primary-light">
        <div className="mx-auto flex max-w-[1440px] flex-col items-center gap-4 px-5 py-10 text-center lg:gap-8 lg:px-20 lg:py-20">
          <h2 className="text-[22px] font-bold leading-[1.3] text-ink-body lg:text-5xl lg:font-extrabold">
            Pronto para começar?
          </h2>
          <p className="max-w-lg text-sm leading-normal text-ink-muted lg:max-w-[640px] lg:text-lg">
            Junte-se a centenas de clínicas e veterinários que já encontraram o match perfeito no VetPlant.
          </p>
          <div className="flex w-full flex-col gap-3 lg:w-auto lg:flex-row lg:gap-4">
            <Link
              href="/signup/company"
              className="flex items-center justify-center rounded-lg bg-primary px-12 py-3 text-[15px] font-semibold text-white transition-colors hover:bg-primary-hover lg:py-[18px] lg:text-lg"
            >
              Cadastrar minha clínica
            </Link>
            <Link
              href="/signup/veterinary"
              className="flex items-center justify-center rounded-lg border-[1.5px] border-primary bg-surface px-12 py-3 text-[15px] font-semibold text-primary transition-colors hover:bg-primary-light lg:py-[18px] lg:text-lg"
            >
              Quero ser plantonista
            </Link>
          </div>
        </div>
      </section>

      <footer className="bg-[#1A2A2A]">
        <div className="hidden lg:block">
          <div className="mx-auto max-w-[1440px] px-20 pt-12 pb-6">
            <div className="flex items-start gap-16 pb-10">
              <div className="flex items-center gap-2.5">
                <PawPrintIcon className="h-6 w-6 text-primary" />
                <span className="text-lg font-bold text-white">{APP_NAME}</span>
              </div>

              <div className="ml-auto flex gap-16">
                <div className="flex flex-col gap-3">
                  <p className="text-sm font-bold text-white">Links úteis</p>
                  <a href="#como-funciona" className="text-sm text-white/60 transition-colors hover:text-white/90">Como funciona</a>
                  <a href="#para-clinicas" className="text-sm text-white/60 transition-colors hover:text-white/90">Para clínicas</a>
                  <a href="#para-veterinarios" className="text-sm text-white/60 transition-colors hover:text-white/90">Para veterinários</a>
                  <a href="#precos" className="text-sm text-white/60 transition-colors hover:text-white/90">Preços</a>
                </div>
                <div className="flex flex-col gap-3">
                  <p className="text-sm font-bold text-white">Suporte</p>
                  <a href="#" className="text-sm text-white/60 transition-colors hover:text-white/90">Central de ajuda</a>
                  <a href="#" className="text-sm text-white/60 transition-colors hover:text-white/90">Contato</a>
                  <Link href="/privacy" className="text-sm text-white/60 transition-colors hover:text-white/90">Política de privacidade</Link>
                  <Link href="/terms" className="text-sm text-white/60 transition-colors hover:text-white/90">Termos de uso</Link>
                </div>
              </div>
            </div>

            <div className="flex items-center justify-between border-t border-white/10 pt-6">
              <div className="flex items-center gap-4">
                <a href="#" aria-label="Instagram" className="text-white/60 transition-colors hover:text-white">
                  <InstagramIcon className="h-5 w-5" />
                </a>
                <a href="#" aria-label="LinkedIn" className="text-white/60 transition-colors hover:text-white">
                  <LinkedinIcon className="h-5 w-5" />
                </a>
                <a href="#" aria-label="X / Twitter" className="text-white/60 transition-colors hover:text-white">
                  <TwitterXIcon className="h-5 w-5" />
                </a>
              </div>
              <p className="text-[13px] text-white/40">
                © 2025 VetPlant. Todos os direitos reservados.
              </p>
            </div>
          </div>
        </div>

        <div className="lg:hidden">
          <div className="flex flex-col gap-6 px-5 pt-10 pb-8">
            <div className="flex items-center gap-2">
              <PawPrintIcon className="h-[22px] w-[22px] text-primary" />
              <span className="text-lg font-bold text-white">{APP_NAME}</span>
            </div>
            <p className="text-[13px] leading-[1.6] text-white/50">
              Conectando clínicas veterinárias e plantonistas em todo o Brasil.
            </p>

            <div className="flex flex-col gap-3.5">
              <p className="text-[13px] font-semibold text-white">Links úteis</p>
              <a href="#como-funciona" className="text-[13px] text-white/50 transition-colors hover:text-white/80">Como funciona</a>
              <a href="#para-clinicas" className="text-[13px] text-white/50 transition-colors hover:text-white/80">Para clínicas</a>
              <a href="#para-veterinarios" className="text-[13px] text-white/50 transition-colors hover:text-white/80">Para veterinários</a>
              <a href="#" className="text-[13px] text-white/50 transition-colors hover:text-white/80">Suporte</a>
            </div>

            <div className="h-px bg-white/10" />

            <p className="text-center text-[11px] text-white/30">
              © 2025 VetPlant. Todos os direitos reservados.
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}

# Requisitos do Projeto

**Nome:** a definir.

Documento de requisitos da aplicação web para **veterinários plantonistas** e **proprietários de clínicas veterinárias**, incluindo cadastro verificado, oferta de trabalho, avaliações e modelo de negócio com segurança.

---

## 1. Visão geral

- **Clínicas** e **plantonistas** se cadastram com dados necessários e comprovação de existência/legitimidade.
- Ambos podem **ofertar e buscar** trabalho (clínicas publicam plantões; plantonistas se candidatam ou são convidados).
- Existe um **sistema de avaliações** mútuas para reputação e confiança.
- A aplicação deve ser **monetizável** e ao mesmo tempo **segura** para todos os envolvidos.

---

## 2. Requisitos funcionais

### 2.1 Cadastro e verificação de identidade

| ID | Requisito | Descrição |
|----|-----------|-----------|
| RF-01 | Cadastro de clínica | Clínica informa: razão social, CNPJ, endereço, telefone, e-mail, responsável legal e documentos de comprovação (alvará, contrato social ou equivalente). |
| RF-02 | Cadastro de plantonista | Plantonista informa: nome, CPF, CRMV (Conselho Regional de Medicina Veterinária), especialidades, disponibilidade geral e documentos (diploma, registro no CRMV). |
| RF-03 | Verificação de existência | Validação de CNPJ (clínicas) e CPF/CRMV (plantonistas) contra fontes oficiais ou parceiros (Receita, conselhos), quando disponível. |
| RF-04 | Upload e análise de documentos | Envio de documentos (PDF/imagem), armazenamento seguro e, quando possível, checagem automática (ex.: CRMV ativo) ou análise manual por equipe. |
| RF-05 | Perfil completo vs. perfil básico | Perfil só pode ofertar ou aceitar trabalho após conclusão do cadastro e aprovação mínima (ex.: documentos aceitos). Perfil básico pode apenas navegar e se cadastrar. |
| RF-06 | Atualização cadastral | Ambos os perfis podem atualizar dados e reenviar documentos; mudanças sensíveis (ex.: CNPJ, CPF) podem exigir revalidação. |

### 2.2 Oferta e busca de trabalho

| ID | Requisito | Descrição |
|----|-----------|-----------|
| RF-07 | Publicação de plantão (clínica) | Clínica cria oferta com: data/hora, duração, tipo (emergência, consulta, cirurgia etc.), valor ofertado, requisitos (ex.: especialidade), descrição e local. |
| RF-08 | Busca e filtros (plantonista) | Plantonista filtra por data, localização, tipo de serviço, valor, especialidade e status (aberto, em andamento, concluído). |
| RF-09 | Candidatura e convite | Plantonista se candidata a plantões; clínica pode convidar plantonistas específicos. Fluxo de aceite/recusa com prazos. |
| RF-10 | Calendário e disponibilidade | Plantonista informa blocos de disponibilidade; clínicas veem compatibilidade ao publicar ou convidar. Opção de sincronização com calendário externo (futuro). |
| RF-11 | Match e confirmação | Após aceite mútuo, plantão fica "confirmado"; ambas as partes recebem confirmação e detalhes (local, contato, valor acordado). |
| RF-12 | Cancelamento e políticas | Regras claras de cancelamento (prazos, multas ou penalidades na reputação), tanto para clínica quanto para plantonista. |

### 2.3 Avaliações e reputação

| ID | Requisito | Descrição |
|----|-----------|-----------|
| RF-13 | Avaliação pós-plantão | Após conclusão do plantão, clínica avalia plantonista e plantonista avalia clínica (nota e comentário opcional). |
| RF-14 | Exibição de reputação | Nota média, quantidade de avaliações e (opcional) comentários recentes visíveis no perfil, respeitando LGPD (sem dados pessoais desnecessários). |
| RF-15 | Indisponibilidade de edição | Avaliações não editáveis após envio; possível política de resposta única da parte avaliada (replica) para contexto. |
| RF-16 | Moderação | Denúncias de avaliações falsas ou ofensivas; equipe pode ocultar ou remover em caso de violação. |

### 2.4 Pagamentos e financeiro (segurança e monetização)

| ID | Requisito | Descrição |
|----|-----------|-----------|
| RF-17 | Pagamento seguro (clínica → plantonista) | Opção de pagamento via plataforma (ex.: intermediado/escrow): clínica deposita valor; liberado ao plantonista após confirmação de conclusão ou prazo. Reduz risco de calote. |
| RF-18 | Pagamento fora da plataforma | Permitir acordos "combinação direta" (ex.: PIX fora do app), com clara indicação de que a plataforma não garante o repasse. |
| RF-19 | Taxa da plataforma | Cobrança de taxa sobre valor do plantão (ex.: % sobre o valor intermediado) ou valor fixo por plantão, conforme plano. |
| RF-20 | Faturamento e comprovantes | Emissão de recibo/comprovante para clínica e plantonista; relatórios básicos para declaração (futuro: integração contábil). |
| RF-21 | Conflitos de pagamento | Canal para disputas (ex.: "não recebi", "serviço não realizado"); medição por equipe e políticas de estorno quando aplicável. |

### 2.5 Comunicação e notificações

| ID | Requisito | Descrição |
|----|-----------|-----------|
| RF-22 | Notificações in-app | Avisos de novas ofertas, candidaturas, convites, confirmações, lembretes de plantão e mensagens importantes. |
| RF-23 | Notificações por e-mail | Resumos e alertas críticos por e-mail (confirmação de cadastro, confirmação de plantão, avaliação pendente). |
| RF-24 | Mensagens entre partes | Canal de mensagens entre clínica e plantonista (por plantão ou geral), com histórico e possível moderação. Evita expor telefone/e-mail antes da confiança. |

### 2.6 Segurança e conformidade

| ID | Requisito | Descrição |
|----|-----------|-----------|
| RF-25 | Autenticação e sessão | Login seguro (e-mail/senha, com possibilidade futura de 2FA); sessões com expiração e logout em todos os dispositivos. |
| RF-26 | Autorização por perfil | Apenas clínicas aprovadas publicam plantões; apenas plantonistas aprovados se candidatam. Dados sensíveis (documentos) acessíveis só ao dono e à administração. |
| RF-27 | LGPD | Consentimento para uso de dados, finalidade clara, direito de acesso, correção e exclusão; política de privacidade e termos de uso. |
| RF-28 | Auditoria | Registro de ações sensíveis (login, alteração de documento, pagamento, disputa) para suporte e eventual auditoria. |

---

## 3. Requisitos de monetização

| ID | Requisito | Descrição |
|----|-----------|-----------|
| RM-01 | Taxa por transação | Percentual ou valor fixo sobre cada plantão pago pela plataforma (quem paga: clínica, plantonista ou divisão). |
| RM-02 | Planos de assinatura (clínicas) | Planos mensais/anuais: básico (X publicações/mês), profissional (ilimitado + destaque), enterprise (múltiplas unidades, relatórios). |
| RM-03 | Planos para plantonistas | Free (buscar e se candidatar com limite) e premium (candidaturas ilimitadas, perfil em destaque, ver quem visualizou). |
| RM-04 | Destaque e visibilidade | Pagamento por destaque em listagens ("plantão em destaque", "perfil em destaque") para maior visibilidade. |
| RM-05 | Selo "Verificado" | Opção paga ou gratuita após verificação reforçada (documentos + possível checagem manual), aumentando confiança e conversão. |
| RM-06 | Relatórios e analytics (premium) | Para clínicas: relatório de plantões, custos, plantonistas mais contratados. Para plantonistas: rendimentos, clínicas que mais contratam. |

---

## 4. Requisitos não funcionais

### 4.1 Segurança

- Dados de documentos e pagamento criptografados em repouso e em trânsito (HTTPS, boas práticas de armazenamento).
- Integração com gateway de pagamento homologado (ex.: Stripe, PagSeguro, Mercado Pago) para não armazenar dados completos de cartão.
- Prevenção a fraudes: limite de cadastros por CPF/CNPJ, detecção de comportamento suspeito (muitos cancelamentos, avaliações extremas em sequência).

### 4.2 Usabilidade e desempenho

- Interface responsiva (mobile-first), considerando uso em celular por plantonistas em deslocamento.
- Tempos de carregamento aceitáveis para listagens e buscas (ex.: < 3s para resultados).
- Acessibilidade básica (contraste, labels, navegação por teclado).

### 4.3 Disponibilidade e suporte

- SLA desejado para disponibilidade do sistema (ex.: 99% uptime em ambiente de produção).
- Canal de suporte (e-mail/chat) e FAQ para dúvidas de cadastro, pagamento e uso.

---

## 5. Requisitos futuros (backlog)

- Integração com calendário (Google Calendar, Outlook).
- App mobile nativo ou PWA completa.
- Contratos digitais (assinatura eletrônica) para cada plantão.
- Integração com sistemas de prontuário das clínicas (apenas leitura para contexto).
- Programa de fidelidade ou benefícios para usuários frequentes.
- Múltiplas unidades por rede de clínicas (gestão centralizada).
- Categorização por espécie/especialidade (pequenos, grandes, exóticos, emergência 24h).

---

## 6. Resumo de prioridades sugeridas

1. **Fase 1 (MVP):** Cadastro e verificação (RF-01 a RF-06), oferta e candidatura básica (RF-07 a RF-11), avaliações (RF-13, RF-14), autenticação e LGPD (RF-25 a RF-27).
2. **Fase 2:** Pagamento intermediado (RF-17, RF-19), notificações (RF-22, RF-23), mensagens (RF-24), políticas de cancelamento (RF-12).
3. **Fase 3:** Monetização (RM-01 a RM-05), disputas (RF-21), relatórios (RM-06), melhorias de segurança e desempenho.

---
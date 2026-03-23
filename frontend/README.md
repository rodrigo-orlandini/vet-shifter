# Frontend (Next.js)

Web app for Vet Shifter (clinics and shift veterinarians), built with **Next.js** and **React**.

## Prerequisites

- **Node.js** 20+ (aligned with `@types/node` in this project)
- **npm** (or pnpm/yarn, if you prefer)

## Initial setup

From the `frontend` directory:

```powershell
cd frontend
npm install
```

## Environment variables

| Variable | Description |
|----------|-------------|
| `NEXT_PUBLIC_API_URL` | Base URL of the backend API. Default in code: `http://localhost:8000`. |

Create a `.env.local` file if you need a non-default API URL:

```env
NEXT_PUBLIC_API_URL=http://localhost:8000
```

The Axios client uses `withCredentials: true` for cookie-based auth.

## Running in development

1. Start the **backend** API (see `../backend/README.md`) so the app can call it.
2. Start the Next.js dev server:

```powershell
cd frontend
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) (default Next.js port).

## API client (Orval)

Types and clients are generated from the backend Swagger file:

- **Input:** `../backend/cmd/api/docs/swagger.json`
- **Output:** `src/api/generated/api.ts`
- **Config:** `orval.config.ts`

After you change backend routes or regenerate Swagger, refresh the client:

```powershell
cd frontend
npm run gen:api
```

You can also use the repository root script `gen-api.ps1`, which runs `swag init` on the backend and then Orval on the frontend.

## Other scripts

| Command | Purpose |
|---------|---------|
| `npm run build` | Production build |
| `npm run start` | Run production server (after `build`) |
| `npm run lint` | ESLint |

# Vet Shifter — Claude Guidelines

## Server vs Client Components

**Default to Server Components.** Only add `"use client"` when strictly necessary.

A component needs `"use client"` if it uses:
- React hooks (`useState`, `useEffect`, `useRef`, `useCallback`, `useContext`, etc.)
- Browser-only APIs (`window`, `document`, `localStorage`, etc.)
- Next.js navigation hooks that require client (`useRouter`, `usePathname`, `useSearchParams`)

A component does **not** need `"use client"` just because:
- It receives `onClick`/`onChange` props (those come from the client tree above it)
- It uses `useId` — this works in RSC
- It renders interactive HTML elements (`<button>`, `<input>`) — those are fine in RSC when event handlers are not defined inline

### Split pattern for large client components

When a component is mostly static but has one small interactive part (e.g. a nav link that depends on the current path), **split it**:

```
AuthTopNav.tsx       (server) — renders the shell: logo, outer layout
AuthNavLink.tsx      ("use client") — only the conditional link using usePathname()
```

This keeps the majority of the component server-rendered and limits hydration to the smallest possible slice.

### Current server/client breakdown

| Component | Type | Reason |
|---|---|---|
| `AuthLayout` | Server | No hooks, static shell |
| `AuthTopNav` | Server | Delegates interactive part to AuthNavLink |
| `AuthNavLink` | Client | usePathname() |
| `AuthCard` | Server | Pure presentational |
| `AuthFooterLinks` | Server | Pure links, no interactivity |
| `StepIndicator` | Server | No hooks, pure JSX |
| `Button` | Server | No hooks; onClick comes from parent |
| `FieldWithError` | Server | useId works in RSC; event handlers come from client tree |
| `FormField`, `Label`, `Input`, `Select`, `Checkbox`, `Badge` | Server | Pure presentational primitives |
| `PasswordFields` | Client | useState (show/hide password) |
| `DocumentUploadSlot` | Client | useRef, useId with side effects |
| `StepLayout` | Client | useCallback, useEffect, useRef |
| `ToastProvider` | Client | Context + state |
| `page.tsx` (all auth pages) | Client | useState, useRouter, form submission |
| Step forms (`Step1Form`, etc.) | Client | Function props from client page; interactivity |

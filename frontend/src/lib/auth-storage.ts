const AUTH_TOKEN_KEY = "auth_token";
const AUTH_USER_KEY = "auth_user";
const AUTH_EXPIRES_AT_KEY = "auth_expires_at";

export interface AuthUser {
  id: string;
  email: string;
  type: string;
}

function isNativePlatform(): boolean {
  if (typeof window === "undefined") return false;
  const cap = (window as unknown as { Capacitor?: { isNativePlatform?: () => boolean } }).Capacitor;
  return cap?.isNativePlatform?.() ?? false;
}

export async function getToken(): Promise<string | null> {
  if (typeof window === "undefined") return null;
  if (isNativePlatform()) {
    const { Preferences } = await import("@capacitor/preferences");
    const { value } = await Preferences.get({ key: AUTH_TOKEN_KEY });
    return value;
  }
  return sessionStorage.getItem(AUTH_TOKEN_KEY) ?? localStorage.getItem(AUTH_TOKEN_KEY);
}

export async function setToken(
  token: string,
  user: AuthUser,
  expiresAt: string,
  rememberMe: boolean
): Promise<void> {
  if (typeof window === "undefined") return;
  const userJson = JSON.stringify(user);
  if (isNativePlatform()) {
    const { Preferences } = await import("@capacitor/preferences");
    await Preferences.set({ key: AUTH_TOKEN_KEY, value: token });
    await Preferences.set({ key: AUTH_USER_KEY, value: userJson });
    await Preferences.set({ key: AUTH_EXPIRES_AT_KEY, value: expiresAt });
    return;
  }
  const storage = rememberMe ? localStorage : sessionStorage;
  storage.setItem(AUTH_TOKEN_KEY, token);
  storage.setItem(AUTH_USER_KEY, userJson);
  storage.setItem(AUTH_EXPIRES_AT_KEY, expiresAt);
}

export async function clearAuth(): Promise<void> {
  if (typeof window === "undefined") return;
  if (isNativePlatform()) {
    const { Preferences } = await import("@capacitor/preferences");
    await Preferences.remove({ key: AUTH_TOKEN_KEY });
    await Preferences.remove({ key: AUTH_USER_KEY });
    await Preferences.remove({ key: AUTH_EXPIRES_AT_KEY });
    return;
  }
  sessionStorage.removeItem(AUTH_TOKEN_KEY);
  sessionStorage.removeItem(AUTH_USER_KEY);
  sessionStorage.removeItem(AUTH_EXPIRES_AT_KEY);
  localStorage.removeItem(AUTH_TOKEN_KEY);
  localStorage.removeItem(AUTH_USER_KEY);
  localStorage.removeItem(AUTH_EXPIRES_AT_KEY);
}

export async function getUser(): Promise<AuthUser | null> {
  if (typeof window === "undefined") return null;
  let raw: string | null;
  if (isNativePlatform()) {
    const { Preferences } = await import("@capacitor/preferences");
    const { value } = await Preferences.get({ key: AUTH_USER_KEY });
    raw = value;
  } else {
    raw = sessionStorage.getItem(AUTH_USER_KEY) ?? localStorage.getItem(AUTH_USER_KEY);
  }
  if (!raw) return null;
  try {
    return JSON.parse(raw) as AuthUser;
  } catch {
    return null;
  }
}

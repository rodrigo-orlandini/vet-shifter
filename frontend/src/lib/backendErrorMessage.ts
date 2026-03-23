import axios from "axios";

const GENERIC_INTERNAL_MESSAGE = "Algo deu errado. Tente novamente.";

type BackendErrorShape = {
  code?: string;
  error?: string;
};

export function getBackendErrorMessage(err: unknown): string {
  if (!axios.isAxiosError(err)) {
    return err instanceof Error ? err.message : GENERIC_INTERNAL_MESSAGE;
  }

  const status = err.response?.status;
  const data = err.response?.data as BackendErrorShape | undefined;

  if (typeof status === "number" && status >= 500) {
    return GENERIC_INTERNAL_MESSAGE;
  }

  const backendMessage = data?.error;
  if (backendMessage && backendMessage.trim().length > 0) {
    return backendMessage;
  }

  return GENERIC_INTERNAL_MESSAGE;
}


import Axios, { type AxiosRequestConfig } from "axios";
import { clearAuth, getToken } from "@/lib/auth-storage";

const baseURL =
  typeof process !== "undefined" && process.env?.NEXT_PUBLIC_API_URL
    ? process.env.NEXT_PUBLIC_API_URL
    : "http://localhost:8080";

export const axiosInstance = Axios.create({
  baseURL,
  headers: {
    "Content-Type": "application/json",
  },
});

axiosInstance.interceptors.request.use(async (config) => {
  const token = await getToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

axiosInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401 && typeof window !== "undefined") {
      await clearAuth();
      window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);

export const customInstance = <T>(
  config: AxiosRequestConfig,
  options?: AxiosRequestConfig
): Promise<T> => {
  return axiosInstance({
    ...config,
    ...options,
  }).then(({ data }) => data as T);
};

export type ErrorType = unknown;
export type BodyType<BodyData> = BodyData;

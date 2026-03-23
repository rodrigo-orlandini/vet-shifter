import Axios, { type AxiosRequestConfig } from "axios";

const baseURL =
  typeof process !== "undefined" && process.env?.NEXT_PUBLIC_API_URL
    ? process.env.NEXT_PUBLIC_API_URL
    : "http://localhost:8000";

export const axiosInstance = Axios.create({
  baseURL,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

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

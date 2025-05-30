import { env } from "@/env";
import { ZodSchema } from "zod";

export const fetchUtil = async <T>(url: `/${string}`, schema: ZodSchema<T>) => {
  const res = await fetch(env.VITE_SERVER_URL + url, {
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
  });

  if (!res.ok) {
    throw new ApiError(res.statusText, res.status);
  }

  // Parse and return the data as the inferred type
  return schema.parse(await res.json());
};

export class ApiError extends Error {
  public isUnauthorized: boolean;

  constructor(
    message: string,
    public status: number,
  ) {
    super(message);

    this.name = "ApiError";
    this.isUnauthorized = status === 401;
  }
}

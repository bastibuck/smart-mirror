import { env } from "@/env";
import { ZodSchema } from "zod";

export const fetchUtil = async <T>(
  url: `/${string}`,
  schema: ZodSchema<T>,
): Promise<T> => {
  const res = await fetch(env.VITE_SERVER_URL + url, {
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
  });

  if (!res.ok) {
    throw new Error(res.statusText);
  }

  // Parse and return the data as the inferred type
  return schema.parse(await res.json());
};

import { z } from "zod";

export enum AppType {
  SIMPLE_TEXT = "simple_text",
}

export const UpsertAppSchema = z.object({
  name: z.string().min(1),
  type: z.nativeEnum(AppType),
});

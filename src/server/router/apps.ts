import { createRouter } from "./context";
import { z } from "zod";
import { AppType, UpsertAppSchema } from "../../types/shared";

export const appsRouter = createRouter()
  .query("getAll", {
    async resolve({ ctx }) {
      const apps = await ctx.prisma.app.findMany();

      return apps.map((app) => ({ ...app, type: app.type as AppType }));
    },
  })
  .mutation("upsertApp", {
    // validate input with Zod
    input: UpsertAppSchema,
    async resolve(params) {
      await params.ctx.prisma.app.create({
        data: { name: params.input.name, type: params.input.type },
      });

      return null;
    },
  });

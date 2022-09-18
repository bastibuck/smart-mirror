import { createRouter } from "./context";
import { z } from "zod";

enum AppType {
  SIMPLE_TEXT = "simple_text",
}

export const appsRouter = createRouter().query("getAll", {
  async resolve({ ctx }) {
    const apps = await ctx.prisma.app.findMany();

    return apps.map((app) => ({ ...app, type: app.type as AppType }));
  },
});

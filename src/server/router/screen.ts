import { createRouter } from "./context";
import { z } from "zod";

export const screenRouter = createRouter().query("getAll", {
  async resolve({ ctx }) {
    return await ctx.prisma.screen.findMany();
  },
});

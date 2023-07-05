import { z } from "zod";
import { createTRPCRouter, publicProcedure } from "~/server/api/trpc";

export const exampleRouter = createTRPCRouter({
  hello: publicProcedure
    .input(z.object({ text: z.string() }))
    .query(({ input }) => {
      return {
        greeting: `Hello ${input.text}`,
      };
    }),
  getAll: publicProcedure.query(({ ctx }) => {
    return ctx.prisma.example.findMany();
  }),

  add: publicProcedure.mutation(({ ctx }) => {
    return ctx.prisma.example.create({ data: {} });
  }),

  delete: publicProcedure
    .input(z.string().cuid())
    .mutation(({ input, ctx }) => {
      return ctx.prisma.example.delete({ where: { id: input } });
    }),
});

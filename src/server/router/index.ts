// src/server/router/index.ts
import { createRouter } from "./context";
import superjson from "superjson";

import { screenRouter } from "./screen";

export const appRouter = createRouter()
  .transformer(superjson)
  .merge("screen.", screenRouter);

// export type definition of API
export type AppRouter = typeof appRouter;

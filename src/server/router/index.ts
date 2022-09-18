// src/server/router/index.ts
import { createRouter } from "./context";
import superjson from "superjson";

import { screenRouter } from "./screen";
import { appsRouter } from "./apps";

export const appRouter = createRouter()
  .transformer(superjson)
  .merge("screen.", screenRouter)
  .merge("apps.", appsRouter);

// export type definition of API
export type AppRouter = typeof appRouter;

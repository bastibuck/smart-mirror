// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-nocheck

import { createEnv } from "@t3-oss/env-core";

import { z } from "zod";

const env = createEnv({
  server: {},

  /**
   * The prefix that client-side variables must have. This is enforced both at
   * a type-level and at runtime.
   */
  clientPrefix: "VITE_",

  client: {
    VITE_SERVER_URL: z.string().url(),
    VITE_IS_PROD: z.boolean().default(false),
    VITE_VERSION_HASH: z.string().default("notset"),
    VITE_STRAVA_CLIENT_ID: z.number({ coerce: true }),
  },

  /**
   * What object holds the environment variables at runtime. This is usually
   * `process.env` or `import.meta.env`.
   */
  runtimeEnv: import.meta.env,
  emptyStringAsUndefined: true,

  runtimeEnvStrict: {
    VITE_SERVER_URL: import.meta.env.VITE_SERVER_URL,
    VITE_IS_PROD: import.meta.env.PROD,
    VITE_VERSION_HASH: import.meta.env.VITE_VERSION_HASH,
    VITE_STRAVA_CLIENT_ID: import.meta.env.VITE_STRAVA_CLIENT_ID,
  },
});

export { env };

import * as Sentry from "@sentry/react";
import { env } from "@/env";

if (env.VITE_IS_PROD && env.VITE_SENTRY_DSN !== "notset") {
  Sentry.init({
    dsn: env.VITE_SENTRY_DSN,
    // Setting this option to true will send default PII data to Sentry.
    // For example, automatic IP address collection on events
    sendDefaultPii: true,
  });
}

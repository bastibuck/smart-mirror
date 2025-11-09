// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-nocheck

import path from "path";
import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import tanstackRouter from "@tanstack/router-plugin/vite";
import tailwindcss from "@tailwindcss/vite";
import { sentryVitePlugin } from "@sentry/vite-plugin";

// https://vitejs.dev/config/
export default defineConfig(async ({ mode }) => {
  import.meta.env = loadEnv(mode, process.cwd(), ["VITE_", "SENTRY_"]);
  await import("./src/env");

  return {
    build: {
      sourcemap: true,
    },

    plugins: [
      tanstackRouter({
        target: "react",
      }),
      react(),
      tailwindcss(),
      sentryVitePlugin({
        org: "bastibuck-org",
        project: "smartmirror",
        authToken: import.meta.env.SENTRY_AUTH_TOKEN,
        disable: !import.meta.env.PROD || !import.meta.env.SENTRY_AUTH_TOKEN,
      }),
    ],

    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
  };
});

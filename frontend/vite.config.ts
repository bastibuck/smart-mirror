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
  import.meta.env = loadEnv(mode, process.cwd());
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
        org: "bastibuck-org", // Your Sentry organization slug
        project: "smartmirror", // Your Sentry project name
        authToken:
          "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",
      }),
    ],

    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
  };
});

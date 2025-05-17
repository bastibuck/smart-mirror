// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-nocheck

import path from "path";
import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import { TanStackRouterVite } from "@tanstack/router-plugin/vite";
import tailwindcss from "@tailwindcss/vite";

// https://vitejs.dev/config/
export default defineConfig(async ({ mode }) => {
  import.meta.env = loadEnv(mode, process.cwd());
  await import("./src/env");

  return {
    plugins: [
      TanStackRouterVite({
        target: "react",
      }),
      react(),
      tailwindcss(),
    ],

    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
  };
});

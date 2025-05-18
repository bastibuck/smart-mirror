// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-nocheck

import path from "path";
import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import { TanStackRouterVite } from "@tanstack/router-plugin/vite";
import tailwindcss from "@tailwindcss/vite";

import * as child from "child_process";

const commitHash = child
  .execSync("git rev-parse --short HEAD")
  .toString()
  .trim();

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

    define: {
      "import.meta.env.VITE_VERSION_HASH": JSON.stringify(commitHash),
    },
  };
});

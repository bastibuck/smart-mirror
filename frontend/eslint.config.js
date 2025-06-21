import js from "@eslint/js";
import globals from "globals";
import tseslint from "typescript-eslint";
import pluginReact from "eslint-plugin-react";
import pluginRouter from "@tanstack/eslint-plugin-router";
import pluginQuery from "@tanstack/eslint-plugin-query";
import eslintConfigPrettier from "eslint-config-prettier/flat";
import { defineConfig } from "eslint/config";

const isStrict =
  typeof process !== "undefined" &&
  // eslint-disable-next-line no-undef
  process.env.ESLINT_STRICT === "true";

export default defineConfig([
  {
    files: ["**/*.{js,mjs,cjs,ts,jsx,tsx}"],
    plugins: { js },
    extends: ["js/recommended"],
  },
  {
    files: ["**/*.{js,mjs,cjs,ts,jsx,tsx}"],
    languageOptions: { globals: globals.browser },
  },
  {
    ignores: ["dist/"],
  },

  tseslint.configs.recommended,
  pluginReact.configs.flat.recommended,
  pluginReact.configs.flat["jsx-runtime"],

  ...pluginRouter.configs["flat/recommended"],
  ...pluginQuery.configs["flat/recommended"],

  eslintConfigPrettier,

  // Overrides at the end to ensure they are not overwritten
  {
    files: ["**/*.{js,mjs,cjs,ts,jsx,tsx}"],
    settings: {
      react: {
        version: "detect",
      },
    },
    rules: {
      "react/prop-types": "off",
      "no-console": isStrict ? "error" : "warn",
      "no-restricted-imports": [
        "error",
        {
          name: "clsx",
          message:
            "We have a convenience wrapper for clsx called 'cn' that uses tailwind-merge to avoid conflicts.",
        },
        {
          name: "zod",
          message:
            "Zod provides zod/v4 version that is already available before the official release. Do not use v3.",
        },
      ],
    },
  },
]);

import React from "react";

import { fetchUtil } from "@/lib/api";
import { useQuery } from "@tanstack/react-query";
import { z } from "zod/v4";
import { env } from "@/env";

const AutoReloader: React.FC = () => {
  useQuery({
    queryKey: ["version-hash"],
    queryFn: async () => {
      const result = await fetchUtil(
        "/version-hash",
        z.object({
          versionHash: z.string(),
        }),
      );

      if (env.VITE_VERSION_HASH !== result.versionHash) {
        // eslint-disable-next-line no-console
        console.warn("Client-Server version mismatch found", {
          client: env.VITE_VERSION_HASH,
          server: result.versionHash,
        });

        window.location.href = "/";
      }

      return result.versionHash;
    },
  });

  return null;
};

export default AutoReloader;

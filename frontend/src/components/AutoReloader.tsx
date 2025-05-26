import React from "react";

import { fetchUtil } from "@/lib/api";
import { useQuery } from "@tanstack/react-query";
import { z } from "zod";
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
        window.location.reload();
      }

      return result.versionHash;
    },
    refetchInterval: 1000 * 60 * 1, // 5 minutes
  });

  return null;
};

export default AutoReloader;

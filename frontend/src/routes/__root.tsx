import {
  QueryCache,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import { Outlet, createRootRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { cn } from "@/lib/utils";
import { env } from "@/env";
import AutoReloader from "@/components/AutoReloader";
import { ApiError } from "@/lib/api";
import { useKeyPressEvent } from "react-use";
import { Toaster } from "@/components/ui/sonner";
import { useAutoRotateRoutes, useDirectionalNavigate } from "@/lib/navigation";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      // global refetch interval for all queries, server is in charge of caching data if needed
      refetchInterval: 1000 * 60 * 5, // 5 minutes

      // do not retry on 401s
      retry(failureCount, error) {
        if (error instanceof ApiError && error.isUnauthorized) {
          return false;
        }

        return failureCount < 3; // Retry up to 3 times
      },
    },
  },

  queryCache: new QueryCache({
    onError: (error) => {
      // eslint-disable-next-line no-console
      console.error("global error in query: ", error);
    },
  }),
});

export const Route = createRootRoute({
  component: RootComponent,
});

function RootComponent() {
  useAutoRotateRoutes();

  const { directionalNavigate } = useDirectionalNavigate();

  useKeyPressEvent("ArrowLeft", () => {
    directionalNavigate("previous");
  });
  useKeyPressEvent("ArrowRight", () => {
    directionalNavigate("next");
  });

  return (
    <>
      <QueryClientProvider client={queryClient}>
        <div
          className={cn("widget-grid [view-transition-name:main-content]", {
            "cursor-none": env.VITE_IS_PROD,
          })}
        >
          <Outlet />
        </div>

        <ReactQueryDevtools initialIsOpen={false} />

        <AutoReloader />
      </QueryClientProvider>

      <TanStackRouterDevtools position="bottom-left" />

      <Toaster />
    </>
  );
}

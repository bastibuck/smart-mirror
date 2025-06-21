import {
  QueryCache,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import {
  Outlet,
  createRootRoute,
  useNavigate,
  useRouter,
  useRouterState,
} from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { cn } from "@/lib/utils";
import { env } from "@/env";
import AutoReloader from "@/components/AutoReloader";
import { ApiError } from "@/lib/api";
import { useKeyPressEvent } from "react-use";
import { Toaster } from "@/components/ui/sonner";
import { toast } from "sonner";
import { useEffect } from "react";

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
  const navigate = useNavigate();

  const { location } = useRouterState();
  const { routesByPath } = useRouter();

  const navigateInDirection = (event: KeyboardEvent) => {
    const direction = event.key === "ArrowRight" ? "next" : "previous";

    const currentPath = location.pathname;

    const currentPageNumber = parseInt(currentPath.split("/").pop()!, 10);
    if (typeof currentPageNumber !== "number" || isNaN(currentPageNumber)) {
      return;
    }

    const targetPageNumber =
      direction === "next" ? currentPageNumber + 1 : currentPageNumber - 1;

    const to = `/${targetPageNumber}`;

    const allRoutes = new Set(Object.keys(routesByPath));
    if (!allRoutes.has(to)) {
      toast.error(`No ${direction} page found`, {
        className: "!bg-destructive !text-foreground",
      });
      return;
    }

    navigate({
      to,
      viewTransition: {
        types: [direction === "next" ? "slide-left" : "slide-right"],
      },
    });
  };

  useKeyPressEvent("ArrowLeft", navigateInDirection);
  useKeyPressEvent("ArrowRight", navigateInDirection);

  useEffect(() => {
    const intervalId = setInterval(
      () => {
        const currentPath = location.pathname;

        const currentPageNumber = parseInt(currentPath.split("/").pop()!, 10);
        if (typeof currentPageNumber !== "number" || isNaN(currentPageNumber)) {
          return;
        }

        const direction = currentPageNumber === 1 ? "next" : "previous";

        const targetPageNumber =
          direction === "next" ? currentPageNumber + 1 : currentPageNumber - 1;

        const to = `/${targetPageNumber}`;
        navigate({
          to,
          viewTransition: {
            types: [direction === "next" ? "slide-left" : "slide-right"],
          },
        });
      },
      1000 * 30, // 30 seconds,
    );

    return () => {
      clearInterval(intervalId);
    };
  }, [location.pathname, navigate]);

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

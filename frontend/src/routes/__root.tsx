import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Outlet, createRootRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      // global refetch interval for all queries, server is in charge of caching data if needed
      refetchInterval: 1000 * 60 * 5, // 5 minutes
    },
  },
});

export const Route = createRootRoute({
  component: RootComponent,
});

function RootComponent() {
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <div className="widget-grid">
          <Outlet />
        </div>

        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>

      <TanStackRouterDevtools position="bottom-left" />
    </>
  );
}

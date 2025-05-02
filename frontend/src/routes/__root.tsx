import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Outlet, createRootRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      // global refetch interval for all queries, server is in charge of caching data if needed
      refetchInterval: 1000 * 60 * 0.5, // 5 minutes
    },
  },
});

export const Route = createRootRoute({
  component: RootComponent,
});

function RootComponent() {
  return (
    <div className="dark">
      <QueryClientProvider client={queryClient}>
        <div className="widget-grid h-screen w-screen bg-black p-4 text-white">
          <Outlet />
        </div>
      </QueryClientProvider>

      <TanStackRouterDevtools position="bottom-right" />
    </div>
  );
}

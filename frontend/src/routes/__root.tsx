import { Outlet, createRootRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

export const Route = createRootRoute({
  component: RootComponent,
});

function RootComponent() {
  return (
    <>
      <div className="widget-grid h-screen w-screen bg-black px-2 py-4 text-white">
        <Outlet />
      </div>

      <TanStackRouterDevtools position="bottom-right" />
    </>
  );
}

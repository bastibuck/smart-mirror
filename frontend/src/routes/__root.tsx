import { Outlet, createRootRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

export const Route = createRootRoute({
  component: RootComponent,
});

function RootComponent() {
  return (
    <>
      <div className="flex h-screen w-screen flex-col items-center justify-center bg-gray-800 text-white">
        <Outlet />
      </div>

      <TanStackRouterDevtools position="bottom-right" />
    </>
  );
}

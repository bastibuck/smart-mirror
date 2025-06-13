import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/(redirect_urls)/strava/token-failure")({
  component: RouteComponent,
});

function RouteComponent() {
  return <div>Something went wront signing you in</div>;
}

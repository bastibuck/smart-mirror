import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/(sports)/strava/token-success")({
  component: RouteComponent,
});

function RouteComponent() {
  // TODO! trigger invalidate of strava stats query

  return (
    <div>
      oAuth successful. You can close this now. Mirror will refetch
      automatically in maximum 5 minutes
    </div>
  );
}

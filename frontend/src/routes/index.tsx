import { createFileRoute } from "@tanstack/react-router";
import StravaProgress from "../widgets/strava/Progress";

export const Route = createFileRoute("/")({
  component: HomeComponent,
});

function HomeComponent() {
  return (
    <>
      <StravaProgress position="center" />
    </>
  );
}

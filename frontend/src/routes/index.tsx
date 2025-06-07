import { createFileRoute } from "@tanstack/react-router";
import AnnualStats from "@/widgets/strava/AnnualStats";
import Clock from "@/widgets/clock/Clock";
import LastActivity from "@/widgets/strava/LastActivity";

export const Route = createFileRoute("/")({
  component: HomeComponent,
});

function HomeComponent() {
  return (
    <>
      <Clock position="top-left" />

      <LastActivity position="bottom-left" />

      <AnnualStats position="bottom-right" />
    </>
  );
}

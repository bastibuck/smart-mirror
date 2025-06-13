import { createFileRoute } from "@tanstack/react-router";
import AnnualStats from "@/widgets/strava/AnnualStats";
import Clock from "@/widgets/clock/Clock";
import LastActivity from "@/widgets/strava/LastActivity";

export const Route = createFileRoute("/(sports)/1")({
  component: SportsPage,
});

function SportsPage() {
  return (
    <>
      <Clock position="top-left" />
      <LastActivity position="bottom-left" />
      <AnnualStats position="bottom-right" />
    </>
  );
}

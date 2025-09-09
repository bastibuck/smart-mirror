import { createFileRoute } from "@tanstack/react-router";
import AnnualStats from "@/widgets/strava/AnnualStats";
import LastActivity from "@/widgets/strava/LastActivity";
import Clock from "@/widgets/clock/Clock";

export const Route = createFileRoute("/1/(sports)/")({
  component: SportsPage,
});

function SportsPage() {
  return (
    <>
      <Clock position="top-left" />
      {/* <StepsOfWeek position="top-right" /> */}
      <LastActivity position="bottom-left" />
      <AnnualStats position="bottom-right" />
    </>
  );
}

import { createFileRoute } from "@tanstack/react-router";
import AnnualStats from "@/widgets/strava/AnnualStats";
import LastActivity from "@/widgets/strava/LastActivity";
import NextDepartures from "@/widgets/kvg/NextDepartures";
import StepsOfWeek from "@/widgets/garmin/StepsOfWeek";

export const Route = createFileRoute("/1/(sports)/")({
  component: SportsPage,
});

function SportsPage() {
  return (
    <>
      <NextDepartures position="top-left" />
      <StepsOfWeek position="top-right" />
      <LastActivity position="bottom-left" />
      <AnnualStats position="bottom-right" />
    </>
  );
}

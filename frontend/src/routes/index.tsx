import { createFileRoute } from "@tanstack/react-router";
import AnnualStats from "@/widgets/strava/AnnualStats";
import Clock from "@/widgets/clock/Clock";

export const Route = createFileRoute("/")({
  component: HomeComponent,
});

function HomeComponent() {
  return (
    <>
      <Clock position="top-right" />

      <AnnualStats position="bottom-right" />
    </>
  );
}

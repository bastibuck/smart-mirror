import { createFileRoute } from "@tanstack/react-router";
import AnnualStats from "@/widgets/strava/AnnualStats";
import Clock from "@/widgets/clock/Clock";
import Map from "@/widgets/Map";

export const Route = createFileRoute("/")({
  component: HomeComponent,
});

function HomeComponent() {
  return (
    <>
      <Clock position="top-right" />

      <Map position="bottom-left" />

      <AnnualStats position="bottom-right" />
    </>
  );
}

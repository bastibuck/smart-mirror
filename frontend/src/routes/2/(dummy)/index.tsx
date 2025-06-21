import Clock from "@/widgets/clock/Clock";
import DailyRecipes from "@/widgets/kptncook/DailyRecipes";
import NextDepartures from "@/widgets/kvg/NextDepartures";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/2/(dummy)/")({
  component: DummyPage,
});

function DummyPage() {
  return (
    <>
      <Clock position="top-left" />
      <DailyRecipes position="top-right" />
      <NextDepartures position="bottom-left" />
    </>
  );
}

import Clock from "@/widgets/clock/Clock";
import NextDepartures from "@/widgets/kvg/NextDepartures";
import SpeedtestResults from "@/widgets/speedtest/SpeedtestResults";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/2/(dummy)/")({
  component: DummyPage,
});

function DummyPage() {
  return (
    <>
      <Clock position="top-left" />
      <NextDepartures position="bottom-left" />
      <SpeedtestResults position="bottom-right" />
    </>
  );
}

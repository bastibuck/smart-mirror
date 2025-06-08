import Clock from "@/widgets/clock/Clock";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/2")({
  component: DummyPage,
});

function DummyPage() {
  return (
    <>
      <Clock position="top-left" />
    </>
  );
}

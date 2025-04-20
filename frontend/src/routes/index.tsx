import { createFileRoute } from "@tanstack/react-router";
import QuoteOfTheDay from "../widgets/quote-of-the-day/QuoteOfTheDay";

export const Route = createFileRoute("/")({
  component: HomeComponent,
});

function HomeComponent() {
  return (
    <>
      <QuoteOfTheDay position="top-right" size="large" />
      <QuoteOfTheDay position="bottom-left" size="large" />
    </>
  );
}

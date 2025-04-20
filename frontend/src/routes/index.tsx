import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: HomeComponent,
});

function HomeComponent() {
  return <h1 className="text-3xl font-bold">Hello world!</h1>;
}

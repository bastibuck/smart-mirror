import { useQuery } from "@tanstack/react-query";

import WidgetPositioner from "../_layout/WidgetPositioner";

const StravaProgress: React.FC<
  React.ComponentProps<typeof WidgetPositioner>
> = ({ ...widgetPositionerProps }) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["strava-progress"],
    queryFn: () => ({
      running: {
        time_s: 300,
        distance_m: 100,
      },
      cycling: {
        time_s: 300,
        distance_m: 100,
      },
      kiting: {
        time_s: 300,
        distance_m: 100,
      },
    }),
  });

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <Component />
    </WidgetPositioner>
  );
};

export default StravaProgress;

import { CartesianGrid, Line, LineChart } from "recharts";

import { ChartConfig, ChartContainer } from "@/components/ui/chart";

const chartData = [
  { running: 186, kiting: 100, cycling: 200 },
  { running: 305, kiting: 100, cycling: 200 },
  { running: 237, kiting: 100, cycling: 300 },
  { running: 73, kiting: 200, cycling: 100 },
  { running: 209, kiting: 100, cycling: 400 },
  { running: 214, kiting: 100, cycling: 200 },
];

const chartConfig = {
  running: {
    color: "var(--chart-1)",
  },
  cycling: {
    color: "var(--chart-2)",
  },
  kiting: {
    color: "var(--chart-3)",
  },
} satisfies ChartConfig;

export function Component() {
  return (
    <ChartContainer config={chartConfig}>
      <LineChart
        data={chartData}
        margin={{
          left: 8,
          right: 8,
          top: 8,
          bottom: 8,
        }}
      >
        <CartesianGrid />

        <Line
          dataKey="running"
          stroke="var(--color-running)"
          dot={{
            fill: "bg-background",
            r: 6,
          }}
          strokeWidth={3}
        />

        <Line
          dataKey="cycling"
          stroke="var(--color-cycling)"
          dot={{
            fill: "bg-background",
            r: 6,
          }}
          strokeWidth={3}
        />

        <Line
          dataKey="kiting"
          stroke="var(--color-kiting)"
          dot={{
            fill: "bg-background",
            r: 6,
          }}
          strokeWidth={3}
        />
      </LineChart>
    </ChartContainer>
  );
}

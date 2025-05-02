import { useQuery } from "@tanstack/react-query";
import { CartesianGrid, Line, LineChart } from "recharts";

import WidgetPositioner from "../_layout/WidgetPositioner";

import { ChartConfig, ChartContainer } from "@/components/ui/chart";

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

const StravaProgress: React.FC<
  React.ComponentProps<typeof WidgetPositioner>
> = ({ ...widgetPositionerProps }) => {
  const { data, isError, error } = useQuery({
    queryKey: ["strava-progress"],
    queryFn: () => [
      { running: 186, kiting: 100, cycling: 200 },
      { running: 305, kiting: 100, cycling: 200 },
      { running: 237, kiting: 100, cycling: 300 },
      { running: 73, kiting: 200, cycling: 100 },
      { running: 209, kiting: 100, cycling: 400 },
      { running: 214, kiting: 100, cycling: 200 },
    ],
    initialData: [],
  });

  if (isError) {
    return (
      <WidgetPositioner {...widgetPositionerProps}>
        TODO! Handle error case!
        {error instanceof Error ? error.message : "Unknown error"}
      </WidgetPositioner>
    );
  }

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <ChartContainer config={chartConfig}>
        <LineChart
          data={data}
          margin={{
            left: 8,
            right: 8,
            top: 8,
            bottom: 8,
          }}
        >
          <CartesianGrid />

          {SportLine({ dataKey: "cycling" })}
          {SportLine({ dataKey: "running" })}
          {SportLine({ dataKey: "kiting" })}
        </LineChart>
      </ChartContainer>
    </WidgetPositioner>
  );
};

export default StravaProgress;

/**
 * @example
 * {SportLine({ dataKey: "cycling" })}
 */
const SportLine: React.FC<{ dataKey: keyof typeof chartConfig }> = ({
  dataKey,
}) => {
  return (
    <Line
      dataKey={dataKey}
      stroke={`var(--color-${dataKey})`}
      dot={{
        fill: "bg-background",
        r: 6,
      }}
      strokeWidth={3}
    />
  );
};

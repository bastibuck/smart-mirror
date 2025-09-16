import React from "react";
import WidgetPositioner from "../_layout/WidgetPositioner";
import { useQuery } from "@tanstack/react-query";
import { fetchUtil } from "@/lib/api";
import { z } from "zod/v4";
import { ChartContainer } from "@/components/ui/chart";
import {
  CartesianGrid,
  LabelList,
  Line,
  LineChart,
  ReferenceLine,
  XAxis,
  YAxis,
} from "recharts";

const SpeedtestResultSchema = z
  .object({
    secondsAgo: z.number(),
    download: z.number(),
    upload: z.number(),
    ping: z.number(),
  })
  .array();

const SpeedtestResults: React.FC<
  React.ComponentProps<typeof WidgetPositioner>
> = ({ ...widgetPositionerProps }) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["speedtest", "results"],
    queryFn: () => fetchUtil("/speedtest", SpeedtestResultSchema),

    staleTime: 1000 * 60 * 1,
    refetchInterval: 1000 * 60 * 1,
  });

  if (isError) {
    return (
      <WidgetPositioner {...widgetPositionerProps}>
        <p>{error.message}</p>
      </WidgetPositioner>
    );
  }

  if (isPending) {
    return <WidgetPositioner {...widgetPositionerProps} />;
  }

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <ChartContainer
        config={{
          download: {
            color: "var(--chart-1)",
          },
          upload: {
            color: "var(--chart-2)",
          },
        }}
        className="mb-3"
      >
        <LineChart
          data={data}
          margin={{ left: 30, right: 30, top: 30, bottom: 0 }}
        >
          <CartesianGrid vertical={false} syncWithTicks />

          <YAxis domain={[0, 100]} hide />

          <XAxis
            dataKey="secondsAgo"
            tickLine={false}
            tickMargin={10}
            axisLine={false}
            tickFormatter={(value) =>
              new Intl.DateTimeFormat("de-DE", {
                timeStyle: "short",
              }).format(new Date().getTime() - value * 1000)
            }
          />

          <ReferenceLine
            y={100}
            strokeDasharray="8 12"
            stroke="var(--color-download)"
          />

          <Line
            dataKey="download"
            stroke="var(--color-download)"
            fill="var(--color-download)"
          >
            <LabelList
              position="top"
              offset={12}
              className="fill-foreground"
              fontSize={12}
              formatter={(value: number) =>
                new Intl.NumberFormat("de-DE", {
                  maximumFractionDigits: 0,
                  unit: "megabit-per-second",
                  unitDisplay: "narrow",
                  style: "unit",
                }).format(value)
              }
            />
          </Line>

          <ReferenceLine
            y={50}
            strokeDasharray="8 12"
            stroke="var(--color-upload)"
          />

          <Line
            dataKey="upload"
            stroke="var(--color-upload)"
            fill="var(--color-upload)"
          >
            <LabelList
              position="top"
              offset={12}
              className="fill-foreground"
              fontSize={12}
              formatter={(value: number) =>
                new Intl.NumberFormat("de-DE", {
                  maximumFractionDigits: 0,
                  unit: "megabit-per-second",
                  unitDisplay: "narrow",
                  style: "unit",
                }).format(value)
              }
            />
          </Line>
        </LineChart>
      </ChartContainer>
    </WidgetPositioner>
  );
};

export default SpeedtestResults;

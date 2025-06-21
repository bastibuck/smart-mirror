import React from "react";
import WidgetPositioner from "../_layout/WidgetPositioner";
import { useQuery } from "@tanstack/react-query";
import { fetchUtil } from "@/lib/api";
import { z } from "zod/v4";
import {
  Bar,
  BarChart,
  CartesianGrid,
  LabelList,
  ReferenceLine,
  XAxis,
  YAxis,
} from "recharts";
import { ChartContainer } from "@/components/ui/chart";

const StepsOfWeekSchema = z.object({
  total: z.number().nonnegative().int(),
  average: z.number().nonnegative().int(),
  days: z
    .object({
      steps: z.number().nonnegative().int(),
      date: z.coerce.date(),
    })
    .array(),
});

const StepsOfWeek: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["steps", "week"],
    queryFn: () => fetchUtil("/steps/weekly", StepsOfWeekSchema),
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

  const maxValue = Math.max(...data.days.map((d) => d.steps));
  const nearestTwoPointFiveKTickValue = Math.ceil(maxValue / 2500) * 2500;

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <ChartContainer
        config={{
          steps: {
            color: "var(--chart-1)",
          },
        }}
        className="mb-3"
      >
        <BarChart data={data.days} margin={{ top: 30 }}>
          <CartesianGrid vertical={false} syncWithTicks />

          <YAxis domain={[0, nearestTwoPointFiveKTickValue]} hide />

          <XAxis
            dataKey="date"
            tickLine={false}
            tickMargin={10}
            axisLine={false}
            tickFormatter={(value) =>
              new Intl.DateTimeFormat("de-DE", { weekday: "short" }).format(
                value,
              )
            }
          />

          <ReferenceLine y={10_000} strokeDasharray="4" stroke="white" />

          <Bar dataKey="steps" fill="var(--color-steps)" radius={8}>
            <LabelList
              position="top"
              offset={12}
              className="fill-foreground"
              fontSize={12}
              formatter={(value: number) =>
                new Intl.NumberFormat("de-DE").format(value)
              }
            />
          </Bar>
        </BarChart>
      </ChartContainer>

      <div className="mx-3 inline-block space-y-1">
        <div className="text-3xl font-semibold">
          {new Intl.NumberFormat("de-DE").format(data.total)}
        </div>
        <div className="text-muted-foreground text-base leading-2">Week</div>
      </div>

      <div className="mx-3 inline-block space-y-1">
        <div className="text-3xl font-semibold">
          {new Intl.NumberFormat("de-DE").format(data.days.at(-1)?.steps ?? 0)}
        </div>
        <div className="text-muted-foreground text-base leading-2">Today</div>
      </div>
    </WidgetPositioner>
  );
};

export default StepsOfWeek;

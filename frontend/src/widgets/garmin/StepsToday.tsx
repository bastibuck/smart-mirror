import React from "react";
import WidgetPositioner from "../_layout/WidgetPositioner";
import { useQuery } from "@tanstack/react-query";
import { fetchUtil } from "@/lib/api";
import { z } from "zod/v4";

const StepsTodaySchema = z.object({
  steps: z.number().nonnegative().int(),
});

const StepsToday: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["steps", "today"],
    queryFn: () => fetchUtil("/steps/today", StepsTodaySchema),
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
      <div className="space-y-1">
        <div className="text-3xl font-semibold">{data.steps}</div>
        <div className="text-muted-foreground text-base leading-2">today</div>
      </div>
    </WidgetPositioner>
  );
};

export default StepsToday;

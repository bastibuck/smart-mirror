import React from "react";
import WidgetPositioner from "../_layout/WidgetPositioner";
import { fetchUtil } from "@/lib/api";
import { useQuery } from "@tanstack/react-query";
import z from "zod/v4";
import Compass from "./Compass";

const WindspeedSchema = z.object({
  windSpeedKn: z.number().nonnegative(),
  gustSpeedKn: z.number().nonnegative(),
  windDirectionDeg: z.number().int().min(0).max(360),
});

const Windspeed: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["windspeed"],
    queryFn: () => fetchUtil("/windspeed", WindspeedSchema),

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
      <Compass
        direction={data.windDirectionDeg}
        speed={data.windSpeedKn}
        gusts={data.gustSpeedKn}
        size={300}
      />
    </WidgetPositioner>
  );
};

export default Windspeed;

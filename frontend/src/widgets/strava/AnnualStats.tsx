import React from "react";
import { useQuery } from "@tanstack/react-query";

import WidgetPositioner from "../_layout/WidgetPositioner";
import { z } from "zod/v4";
import { ApiError, fetchUtil } from "@/lib/api";
import TypeIcon from "./components/TypeIcon";
import { StatCategory, StatValue } from "./components/Stats";
import { formatDuration } from "./utils/date";
import Login from "./components/Login";

const SportsStatsSchema = z.object({
  count: z.number(),
  distance_m: z.number(),
  moving_time_s: z.number(),
});

const AnnualStatsSchema = z.object({
  running: SportsStatsSchema,
  cycling: SportsStatsSchema,
  hiking: SportsStatsSchema,
  kiting: SportsStatsSchema,
});

const AnnualStats: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["strava", "annual"],
    queryFn: () => fetchUtil("/strava/annual", AnnualStatsSchema),
  });

  if (isError) {
    if (error instanceof ApiError && error.isUnauthorized) {
      return (
        <WidgetPositioner {...widgetPositionerProps}>
          <Login />
        </WidgetPositioner>
      );
    }

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
      <StatCategory name={<TypeIcon type="Run" />}>
        <StatValue label="#" value={data.running.count.toString()} />
        <StatValue
          label="km"
          value={Math.floor(data.running.distance_m / 1000).toString()}
        />
        <StatValue
          label="hh:mm"
          value={formatDuration(data.running.moving_time_s, {
            showSeconds: false,
          })}
        />
      </StatCategory>

      <StatCategory name={<TypeIcon type="Ride" />}>
        <StatValue label="#" value={data.cycling.count.toString()} />
        <StatValue
          label="km"
          value={Math.floor(data.cycling.distance_m / 1000).toString()}
        />
        <StatValue
          label="hh:mm"
          value={formatDuration(data.cycling.moving_time_s, {
            showSeconds: false,
          })}
        />
      </StatCategory>

      <StatCategory name={<TypeIcon type="Hike" />}>
        <StatValue label="#" value={data.hiking.count.toString()} />
        <StatValue
          label="km"
          value={Math.floor(data.hiking.distance_m / 1000).toString()}
        />
        <StatValue
          label="hh:mm"
          value={formatDuration(data.hiking.moving_time_s, {
            showSeconds: false,
          })}
        />
      </StatCategory>

      <StatCategory name={<TypeIcon type="Kite" />}>
        <StatValue label="#" value={data.kiting.count.toString()} />
        <StatValue
          label="km"
          value={Math.floor(data.kiting.distance_m / 1000).toString()}
        />
        <StatValue
          label="hh:mm"
          value={formatDuration(data.kiting.moving_time_s, {
            showSeconds: false,
          })}
        />
      </StatCategory>
    </WidgetPositioner>
  );
};

export default AnnualStats;

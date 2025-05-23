import React from "react";
import { useQuery } from "@tanstack/react-query";

import WidgetPositioner from "../_layout/WidgetPositioner";
import { z } from "zod";
import { fetchUtil } from "@/lib/api";
import { Bike, Turtle } from "lucide-react";

const SportsStatsSchema = z.object({
  count: z.number(),
  distance_m: z.number(),
  moving_time_s: z.number(),
});

const AnnualStatsSchema = z.object({
  running: SportsStatsSchema,
  cycling: SportsStatsSchema,
});

const AnnualStats: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["strava-annual-stats"],
    queryFn: () => fetchUtil("/strava-stats", AnnualStatsSchema),
  });

  if (isError) {
    return (
      <WidgetPositioner {...widgetPositionerProps}>
        <p>TODO! Handle error case!</p>
        <p>{error.message}</p>
      </WidgetPositioner>
    );
  }

  if (isPending) {
    return (
      <WidgetPositioner {...widgetPositionerProps}>
        <p>TODO! Handle loading!</p>
      </WidgetPositioner>
    );
  }

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <StatCategory name={<Bike size={50} />}>
        <StatValue label="#" value={data.cycling.count.toString()} />
        <StatValue
          label="km"
          value={Math.floor(data.cycling.distance_m / 1000).toString()}
        />
        <StatValue
          label="hh:mm"
          value={formatTime(data.cycling.moving_time_s)}
        />
      </StatCategory>

      <StatCategory name={<Turtle size={50} />}>
        <StatValue label="#" value={data.running.count.toString()} />
        <StatValue
          label="km"
          value={Math.floor(data.running.distance_m / 1000).toString()}
        />
        <StatValue
          label="hh:mm"
          value={formatTime(data.running.moving_time_s)}
        />
      </StatCategory>
    </WidgetPositioner>
  );
};

export default AnnualStats;

const StatCategory: React.FC<
  React.PropsWithChildren<{ name: React.ReactElement }>
> = ({ name, children }) => {
  return (
    <div className="mb-9 grid grid-cols-3 gap-x-6">
      <div className="text-muted-foreground col-span-3 flex justify-end text-3xl">
        {name}
      </div>
      {children}
    </div>
  );
};

const StatValue: React.FC<{ label: string; value: string }> = ({
  label,
  value,
}) => {
  return (
    <div className="space-y-1">
      <div className="text-3xl font-semibold">{value}</div>
      <div className="text-muted-foreground text-base leading-2">{label}</div>
    </div>
  );
};

const formatTime = (seconds: number): string => {
  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  return `${hours.toString().padStart(2, "0")}:${minutes.toString().padStart(2, "0")}`;
};

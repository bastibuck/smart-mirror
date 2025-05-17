import { useQuery } from "@tanstack/react-query";

import WidgetPositioner from "../_layout/WidgetPositioner";
import { z } from "zod";
import { fetchUtil } from "@/lib/api";

const SportsStatsSchema = z.object({
  count: z.number(),
  distance: z.number(),
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
      <div className="mb-7 grid grid-cols-3 gap-x-12 gap-y-1">
        <h2 className="col-span-3 text-3xl font-bold text-pretty">Cycling</h2>

        <div className="w-full">
          <div className="mb-2 text-4xl font-semibold">
            {data.cycling.count}
          </div>
          <div className="text-muted-foreground text-base leading-6 lg:text-lg">
            Count
          </div>
        </div>

        <div className="w-full">
          <div className="mb-2 text-4xl font-semibold">
            {Math.floor(data.cycling.distance / 1000)}
          </div>
          <div className="text-muted-foreground text-base leading-6 lg:text-lg">
            Distance (km)
          </div>
        </div>

        <div className="w-full">
          <div className="mb-2 text-4xl font-semibold">
            {formatTime(data.cycling.moving_time_s)}
          </div>
          <div className="text-muted-foreground text-base leading-6 lg:text-lg">
            Hours
          </div>
        </div>
      </div>

      <div className="grid grid-cols-3 gap-x-12 gap-y-1">
        <h2 className="col-span-3 text-3xl font-bold text-pretty">Running</h2>

        <div className="w-full">
          <div className="mb-2 text-4xl font-semibold">
            {data.running.count}
          </div>
          <div className="text-muted-foreground text-base leading-6">Count</div>
        </div>

        <div className="w-full">
          <div className="mb-2 text-4xl font-semibold">
            {Math.floor(data.running.distance / 1000)}
          </div>
          <div className="text-muted-foreground text-base leading-6">
            Distance (km)
          </div>
        </div>

        <div className="w-full">
          <div className="mb-2 text-4xl font-semibold">
            {formatTime(data.running.moving_time_s)}
          </div>
          <div className="text-muted-foreground text-base leading-6">Hours</div>
        </div>
      </div>
    </WidgetPositioner>
  );
};

export default AnnualStats;

const formatTime = (seconds: number): string => {
  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  return `${hours.toString().padStart(2, "0")}:${minutes.toString().padStart(2, "0")}`;
};

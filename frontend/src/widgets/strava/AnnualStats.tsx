import React from "react";
import { useQuery } from "@tanstack/react-query";
import { QRCodeSVG } from "qrcode.react";

import WidgetPositioner from "../_layout/WidgetPositioner";
import { z } from "zod";
import { ApiError, fetchUtil } from "@/lib/api";
import { Bike, Turtle, Wind, Mountain } from "lucide-react";
import { env } from "@/env";

const STRAVA_LOGIN_URL = `http://www.strava.com/oauth/authorize?client_id=${env.VITE_STRAVA_CLIENT_ID}&response_type=code&redirect_uri=${env.VITE_SERVER_URL}/strava/exchange-token&scope=profile:read_all,activity:read_all`;

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
    queryKey: ["strava-annual-stats"],
    queryFn: () => fetchUtil("/strava/stats", AnnualStatsSchema),
  });

  if (isError) {
    if (error instanceof ApiError && error.isUnauthorized) {
      return (
        <WidgetPositioner {...widgetPositionerProps}>
          <div className="space-y-4">
            <p className="text-xl">Please log in to see your Strava stats.</p>

            {env.VITE_IS_PROD === false ? (
              <a
                href={STRAVA_LOGIN_URL}
                target="_blank"
                rel="noopener noreferrer"
                className="mb-32 block"
              >
                Login
              </a>
            ) : (
              <QRCodeSVG
                value={STRAVA_LOGIN_URL}
                size={280}
                className="inline"
              />
            )}
          </div>
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

      <StatCategory name={<Mountain size={50} />}>
        <StatValue label="#" value={data.hiking.count.toString()} />
        <StatValue
          label="km"
          value={Math.floor(data.hiking.distance_m / 1000).toString()}
        />
        <StatValue
          label="hh:mm"
          value={formatTime(data.hiking.moving_time_s)}
        />
      </StatCategory>

      <StatCategory name={<Wind size={50} />}>
        <StatValue label="#" value={data.kiting.count.toString()} />
        <StatValue
          label="km"
          value={Math.floor(data.kiting.distance_m / 1000).toString()}
        />
        <StatValue
          label="hh:mm"
          value={formatTime(data.kiting.moving_time_s)}
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

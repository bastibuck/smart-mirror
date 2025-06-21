import React from "react";
import { useQuery } from "@tanstack/react-query";
import { ApiError, fetchUtil } from "@/lib/api";
import { z } from "zod/v4";
import MapLibre, { Layer, Source } from "react-map-gl/maplibre";
import "maplibre-gl/dist/maplibre-gl.css";
import WidgetPositioner from "../_layout/WidgetPositioner";
import TypeIcon from "./components/TypeIcon";
import { formatDuration } from "./utils/date";
import { StatValue } from "./components/Stats";
import Login from "./components/Login";

const LastActivitySchema = z
  .object({
    name: z.string(),
    date: z.coerce.date(),
    type: z.enum(["Run", "Ride", "Hike", "Kite"]),
    coordinates: z
      .tuple([z.number(), z.number()])
      .array()
      .describe("Coordinates of the last activity in [lng, lat] format"),
    distance_m: z.number(),
    moving_time_s: z.number(),
  })
  .nullable();

type LastActivityType = NonNullable<z.infer<typeof LastActivitySchema>>;

const LastActivity: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["strava", "last-activity"],
    queryFn: () => fetchUtil("/strava/last-activity", LastActivitySchema),
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

  if (data === null) {
    return <WidgetPositioner {...widgetPositionerProps} />;
  }

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <div className="relative h-full w-full">
        <MapLibre
          onData={(e) => {
            if (
              e.dataType === "source" &&
              (e.sourceDataChanged || e.isSourceLoaded) &&
              e.source.type === "geojson" &&
              typeof e.source.data !== "string" &&
              e.source.data.type === "Feature" &&
              e.source.data.geometry.type === "LineString"
            ) {
              e.target.fitBounds(
                getMaxBounds(e.source.data.geometry.coordinates),
                {
                  animate: false,
                  padding: 40,
                },
              );
            }
          }}
          interactive={false}
          attributionControl={{ compact: false }}
        >
          <Source
            type="geojson"
            data={{
              type: "Feature",
              properties: null,
              geometry: {
                type: "LineString",
                coordinates: data.coordinates,
              },
            }}
          >
            <Layer
              type="line"
              paint={{
                "line-color": "#555",
                "line-width": 3,
              }}
            />
          </Source>
        </MapLibre>

        <div className="absolute inset-0 grid place-items-center text-shadow-black text-shadow-lg">
          <div className="space-y-2">
            <SharedInfo activity={data} />
            <TypedInfo activity={data} />
          </div>
        </div>
      </div>
    </WidgetPositioner>
  );
};

export default LastActivity;

/**
 * calculates maximum bounds from a list of [lng, lat]
 *
 * @returns a list of max bounds with order [west, south, east, north]
 */
function getMaxBounds(coords: number[][]) {
  const lats = coords.map(([, lat]) => lat);
  const lngs = coords.map(([lng]) => lng);

  const minLat = Math.min(...lats);
  const maxLat = Math.max(...lats);
  const minLng = Math.min(...lngs);
  const maxLng = Math.max(...lngs);

  return [minLng, minLat, maxLng, maxLat] as [number, number, number, number];
}

const SharedInfo: React.FC<{ activity: LastActivityType }> = ({ activity }) => {
  return (
    <>
      <div className="flex justify-between gap-2">
        <StatValue
          label={new Intl.DateTimeFormat("de-DE", {
            dateStyle: "medium",
            timeStyle: "short",
            timeZone: "UTC",
          }).format(activity.date)}
          value={activity.name}
        />
        <TypeIcon type={activity.type} className="text-muted-foreground" />
      </div>

      <StatValue
        inline
        label="km"
        value={(activity.distance_m / 1000).toFixed(1)}
      />
    </>
  );
};

const TypedInfo: React.FC<{
  activity: LastActivityType;
}> = ({ activity }) => {
  switch (activity.type) {
    case "Run":
      return <Run activity={activity} />;

    case "Ride":
      return <Ride activity={activity} />;

    case "Hike":
      return <Hike activity={activity} />;

    case "Kite":
      return <Kite activity={activity} />;
  }
};

const Run: React.FC<{ activity: LastActivityType }> = ({ activity }) => {
  const isMoreThanOneHour = activity.moving_time_s > 3600;

  return (
    <>
      <StatValue
        inline
        label={isMoreThanOneHour ? "hh:mm:ss" : "mm:ss"}
        value={formatDuration(activity.moving_time_s, {
          showHours: isMoreThanOneHour,
        })}
      />

      <StatValue
        inline
        label="/km"
        value={formatDuration(
          (activity.moving_time_s / 60 / (activity.distance_m / 1000)) * 60,
          {
            showHours: false,
          },
        )}
      />
    </>
  );
};

const Ride: React.FC<{ activity: LastActivityType }> = ({ activity }) => {
  return (
    <>
      <StatValue
        inline
        label="hh:mm:ss"
        value={formatDuration(activity.moving_time_s)}
      />

      <StatValue
        inline
        label="⌀ km/h"
        value={(
          activity.distance_m /
          1000 /
          (activity.moving_time_s / 60 / 60)
        ).toFixed(1)}
      />
    </>
  );
};

const Hike: React.FC<{ activity: LastActivityType }> = ({ activity }) => {
  return (
    <>
      <StatValue
        inline
        label="hh:mm"
        value={formatDuration(activity.moving_time_s, {
          showSeconds: false,
        })}
      />
    </>
  );
};

const Kite: React.FC<{ activity: LastActivityType }> = ({ activity }) => {
  return (
    <>
      <StatValue
        inline
        label="hh:mm"
        value={formatDuration(activity.moving_time_s, {
          showSeconds: false,
        })}
      />
    </>
  );
};

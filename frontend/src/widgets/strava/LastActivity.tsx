import React from "react";
import { useQuery } from "@tanstack/react-query";
import { ApiError, fetchUtil } from "@/lib/api";
import { z } from "zod";
import MapLibre, { Layer, Source } from "react-map-gl/maplibre";
import "maplibre-gl/dist/maplibre-gl.css";
import WidgetPositioner from "../_layout/WidgetPositioner";
import TypeIcon from "./components/TypeIcon";
import { formatDuration } from "./utils/date";
import { StatValue } from "./components/Stats";
import Login from "./components/Login";

const LastActivitySchema = z
  .object({
    coordinates: z
      .tuple([z.number(), z.number()])
      .array()
      .describe("Coordinates of the last activity in [lng, lat] format"),
    type: z.enum(["Run", "Ride", "Hike", "Kite"]),
    distance_m: z.number(),
    moving_time_s: z.number(),
  })
  .nullable();

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
                "line-color": "#333",
                "line-width": 3,
              }}
            />
          </Source>
        </MapLibre>

        <div className="absolute inset-0 grid place-items-center">
          <div className="space-y-2">
            <TypeIcon type={data.type} className="text-muted-foreground" />

            <StatValue
              inline
              label="km"
              value={(data.distance_m / 1000).toFixed(1)}
            />

            <StatValue
              inline
              label="mm:ss"
              value={formatDuration(data.moving_time_s, {
                mode: data.type === "Run" ? "minutes" : "hours",
              })}
            />

            {data.type === "Run" ? (
              <>
                <StatValue
                  inline
                  label="/km"
                  value={formatDuration(
                    (data.moving_time_s / 60 / (data.distance_m / 1000)) * 60,
                    {
                      mode: "minutes",
                    },
                  )}
                />
              </>
            ) : null}

            {data.type === "Ride" ? (
              <>
                <StatValue
                  inline
                  label="⌀ km/h"
                  value={(
                    data.distance_m /
                    1000 /
                    (data.moving_time_s / 60 / 60)
                  ).toFixed(1)}
                />
              </>
            ) : null}
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

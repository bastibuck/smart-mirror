import React from "react";
import { useQuery } from "@tanstack/react-query";
import { ApiError, fetchUtil } from "@/lib/api";
import { z } from "zod";
import { env } from "@/env";
import { QRCodeSVG } from "qrcode.react";
import MapLibre, { Layer, Source } from "react-map-gl/maplibre";
import "maplibre-gl/dist/maplibre-gl.css";
import WidgetPositioner from "../_layout/WidgetPositioner";

const STRAVA_LOGIN_URL = `http://www.strava.com/oauth/authorize?client_id=${env.VITE_STRAVA_CLIENT_ID}&response_type=code&redirect_uri=${env.VITE_SERVER_URL}/strava/exchange-token&scope=profile:read_all,activity:read_all`;

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

  if (data === null) {
    return <WidgetPositioner {...widgetPositionerProps} />;
  }

  const bounds = getMaxBounds(data.coordinates);

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <MapLibre
        onLoad={(e) => {
          e.target.fitBounds(bounds, {
            animate: false,
          });
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
              "line-color": "white",
            }}
          />
        </Source>
      </MapLibre>
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

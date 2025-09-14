import React from "react";
import WidgetPositioner from "../_layout/WidgetPositioner";
import { useQuery } from "@tanstack/react-query";
import { fetchUtil } from "@/lib/api";
import { z } from "zod/v4";
import { cn } from "@/lib/utils";
import { TriangleAlert } from "lucide-react";

const NextDeparturesSchema = z.object({
  stopName: z.string(),
  departures: z.array(
    z.object({
      line: z.string(),
      destination: z.string(),
      departureTime: z.string(),
      delayMinutes: z.number().optional(),
    }),
  ),
  alerts: z.array(z.string()),
});

const NextDepartures: React.FC<
  React.ComponentProps<typeof WidgetPositioner>
> = ({ ...widgetPositionerProps }) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["transportation", "departures"],
    queryFn: () =>
      fetchUtil("/transportation/departures?limit=7", NextDeparturesSchema),

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
      <div className="inline-block space-y-1">
        <div className="text-3xl font-semibold">{data.stopName}</div>

        {data.alerts.map((alert, index) => (
          <div
            className="mb-2 flex items-center gap-2 rounded border border-yellow-500 p-2 text-sm text-yellow-500"
            key={index}
          >
            <TriangleAlert />
            {alert}
          </div>
        ))}

        {data.departures.map((departure) => (
          <div
            className="flex gap-6"
            key={`${departure.line}-${departure.destination}-${departure.departureTime}`}
          >
            <div
              className={cn("text-left font-mono text-lg", {
                "font-black text-red-500 italic": departure.delayMinutes,
              })}
            >
              <span>{departure.departureTime} </span>
            </div>

            <div className="flex gap-2">
              <span className="font-mono text-lg">{departure.line}</span>
              <span className="text-lg">{departure.destination}</span>
            </div>
          </div>
        ))}
      </div>
    </WidgetPositioner>
  );
};

export default NextDepartures;

import React from "react";
import WidgetPositioner from "../_layout/WidgetPositioner";
import { useQuery } from "@tanstack/react-query";
import { fetchUtil } from "@/lib/api";
import { z } from "zod";
import { cn } from "@/lib/utils";

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

        {data.departures.map((departure) => (
          <>
            <div className="flex gap-6">
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
          </>
        ))}
      </div>
    </WidgetPositioner>
  );
};

export default NextDepartures;

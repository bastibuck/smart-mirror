import { useQuery } from "@tanstack/react-query";

import WidgetPositioner from "../_layout/WidgetPositioner";
import { Progress } from "@/components/ui/progress";

const StravaProgress: React.FC<
  React.ComponentProps<typeof WidgetPositioner>
> = ({ ...widgetPositionerProps }) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["strava-progress"],
    queryFn: () => ({
      running: {
        time_s: 300,
        distance_m: 100,
      },
      cycling: {
        time_s: 300,
        distance_m: 100,
      },
      kiting: {
        time_s: 300,
        distance_m: 100,
      },
    }),
  });

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <Progress value={33} />
    </WidgetPositioner>
  );
};

export default StravaProgress;

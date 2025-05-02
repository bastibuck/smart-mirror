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
        hours: 300,
        distance: 100,
      },
      cycling: {
        hours: 300,
        distance: 100,
      },
      kiting: {
        hours: 300,
        distance: 100,
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

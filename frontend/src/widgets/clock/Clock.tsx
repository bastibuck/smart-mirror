import { useEffect, useState } from "react";
import WidgetPositioner from "../_layout/WidgetPositioner";

const Clock: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const [now, setNow] = useState(new Date());

  const timeString = new Intl.DateTimeFormat("de-DE", {
    timeStyle: "short",
  }).format(now);
  const dateString = new Intl.DateTimeFormat("de-DE", {
    dateStyle: "short",
  }).format(now);
  const [hours, minutes] = timeString.split(":");

  useEffect(() => {
    const intervalId = setInterval(() => {
      setNow(new Date());
    }, 15_000);

    return () => {
      clearInterval(intervalId);
    };
  }, []);

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <h1 className="flex flex-col font-mono text-3xl font-bold">
        <span>{dateString}</span>

        <span>
          {hours}
          <span className="animate-caret-blink animation-duration-[1s]">:</span>
          {minutes}
        </span>
      </h1>
    </WidgetPositioner>
  );
};

export default Clock;

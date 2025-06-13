import { Bike, Mountain, Turtle, Wind } from "lucide-react";
import React from "react";

const ICON_MAP = {
  Run: <Turtle />,
  Ride: <Bike />,
  Hike: <Mountain />,
  Kite: <Wind />,
};

const TypeIcon: React.FC<{
  type: keyof typeof ICON_MAP;
  className?: string;
}> = ({ type, className }) => {
  return React.cloneElement(ICON_MAP[type], {
    className,
    size: 40,
  });
};

export default TypeIcon;

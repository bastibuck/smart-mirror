import { cn } from "@/lib/utils";

type Postion =
  | "top-left"
  | "top-right"
  | "bottom-left"
  | "bottom-right"
  | "center";
type Size = "large" | "full";

const WidgetPositioner: React.FC<
  React.PropsWithChildren<{ position: Postion; size?: Size }>
> = ({ position, size, children }) => {
  const positionClass = `widget widget--${position}`;
  const sizeClass = size ? `widget--${size}` : "";

  // TODO: try moving this to tailwind
  return <div className={cn(positionClass, sizeClass)}>{children}</div>;
};

export default WidgetPositioner;

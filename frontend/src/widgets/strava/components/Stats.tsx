import { cn } from "@/lib/utils";

const StatCategory: React.FC<
  React.PropsWithChildren<{ name: React.ReactElement }>
> = ({ name, children }) => {
  return (
    <div className="mb-6 grid grid-cols-3 gap-x-2">
      <div className="text-muted-foreground col-span-3 flex justify-end text-3xl">
        {name}
      </div>
      {children}
    </div>
  );
};

const StatValue: React.FC<{
  label: string;
  value: string | number;
  inline?: boolean;
  decimals?: number;
}> = ({ label, value, inline = false, decimals = 0 }) => {
  const formattedValue =
    typeof value === "number"
      ? new Intl.NumberFormat(undefined, {
          maximumFractionDigits: decimals,
          minimumFractionDigits: decimals,
        }).format(value)
      : value;

  return (
    <div
      className={cn({
        "space-y-1": inline === false,
        "flex items-baseline space-x-1": inline === true,
      })}
    >
      <div className="text-3xl font-semibold">{formattedValue}</div>
      <div className="text-muted-foreground text-base leading-2">{label}</div>
    </div>
  );
};

export { StatCategory, StatValue };

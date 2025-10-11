interface CompassProps {
  speed: number;
  gusts: number;
  direction: number;
  size?: number;
}

const cardinalDirections = [
  { label: "N", angle: 0 },
  { label: "NE", angle: 45 },
  { label: "E", angle: 90 },
  { label: "SE", angle: 135 },
  { label: "S", angle: 180 },
  { label: "SW", angle: 225 },
  { label: "W", angle: 270 },
  { label: "NW", angle: 315 },
];

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const getDirectionLabel = (deg: number) => {
  const normalized = ((deg % 360) + 360) % 360;
  const index = Math.round(normalized / 45) % 8;

  return cardinalDirections[index].label;
};

const Compass: React.FC<CompassProps> = ({
  speed,
  gusts,
  direction,
  size = 200,
}) => {
  return (
    <div className="inline-block w-fit">
      <div className="flex flex-col items-center gap-4">
        <div className="relative" style={{ width: size, height: size }}>
          {/* Compass circle */}
          <div className="border-border bg-background absolute inset-0">
            {/* Cardinal direction markers */}
            {cardinalDirections.map(({ label, angle }) => {
              const isMainDirection = angle % 90 === 0;
              const radius = size / 2;
              const markerDistance = radius - 32;

              const x =
                radius + markerDistance * Math.sin((angle * Math.PI) / 180);
              const y =
                radius - markerDistance * Math.cos((angle * Math.PI) / 180);

              return (
                <div
                  key={label}
                  className="absolute -translate-x-1/2 -translate-y-1/2"
                  style={{ left: x, top: y }}
                >
                  <span
                    className={`font-semibold ${isMainDirection ? "text-base" : "text-muted-foreground text-xs"}`}
                  >
                    {label}
                  </span>
                </div>
              );
            })}

            {/* Degree markers */}
            <svg
              className="absolute inset-0 h-full w-full"
              viewBox={`0 0 ${size} ${size}`}
            >
              {Array.from({ length: 36 }).map((_, i) => {
                const angle = i * 10;
                const isMainMarker = angle % 90 === 0;
                const isSecondaryMarker = angle % 45 === 0;
                const radius = size / 2;
                const innerRadius = isMainMarker
                  ? radius - 16
                  : isSecondaryMarker
                    ? radius - 12
                    : radius - 8;
                const outerRadius = radius - 4;

                const x1 =
                  radius + innerRadius * Math.sin((angle * Math.PI) / 180);
                const y1 =
                  radius - innerRadius * Math.cos((angle * Math.PI) / 180);
                const x2 =
                  radius + outerRadius * Math.sin((angle * Math.PI) / 180);
                const y2 =
                  radius - outerRadius * Math.cos((angle * Math.PI) / 180);

                return (
                  <line
                    key={i}
                    x1={x1}
                    y1={y1}
                    x2={x2}
                    y2={y2}
                    stroke="currentColor"
                    strokeWidth={isMainMarker ? 2 : 1}
                    className={
                      isMainMarker
                        ? "text-foreground"
                        : "text-muted-foreground/40"
                    }
                  />
                );
              })}
            </svg>

            {/* Wind direction needle */}
            <div
              className="absolute inset-0 transition-transform duration-500 ease-out"
              style={{ transform: `rotate(${180 + direction}deg)` }}
            >
              <svg className="h-full w-full" viewBox={`0 0 ${size} ${size}`}>
                {/* Needle pointing to wind direction */}
                <path
                  d={`M ${size / 2} ${size / 2 - size / 3} L ${size / 2 + 6} ${size / 2 + size / 4} L ${size / 2} ${size / 2 + size / 5} L ${size / 2 - 6} ${size / 2 + size / 4} Z`}
                  fill="hsl(var(--destructive))"
                  stroke="hsl(var(--destructive-foreground))"
                  strokeWidth="1"
                  className="fill-muted-foreground"
                />
                {/* Tail of needle */}
                <path
                  d={`M ${size / 2} ${size / 2 + size / 5} L ${size / 2 + 4} ${size / 2 + size / 3} L ${size / 2} ${size / 2 + size / 3.5} L ${size / 2 - 4} ${size / 2 + size / 3} Z`}
                  strokeWidth="1"
                  className="fill-muted-foreground"
                />
              </svg>
            </div>

            {/* Center circle */}
            <div className="bg-background border-border absolute top-1/2 left-1/2 h-8 w-8 -translate-x-1/2 -translate-y-1/2 rounded-full border-2 shadow-md" />
          </div>
        </div>

        {/* Wind information */}
        <div className="space-y-0 text-center">
          <div className="text-3xl font-bold">
            {speed}{" "}
            <span className="text-muted-foreground text-lg">{"kn "}</span>
          </div>

          <div className="text-muted-foreground text-sm">
            {gusts}
            {" kn"}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Compass;

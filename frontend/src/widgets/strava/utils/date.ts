const formatDuration = (
  seconds: number,
  options?: {
    showSeconds?: boolean;
    showHours?: boolean;
  },
): string => {
  const showHours = options?.showHours ?? true;
  const showSeconds = options?.showSeconds ?? true;

  const timeString = new Intl.DateTimeFormat("de-DE", {
    timeZone: "UTC",
    hour: showHours ? "2-digit" : undefined,
    minute: "2-digit",
    second: showSeconds ? "2-digit" : undefined,
  }).format(new Date(seconds * 1000));

  return timeString;
};

export { formatDuration };

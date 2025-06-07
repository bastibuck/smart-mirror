const formatDuration = (
  seconds: number,
  options: { mode: "minutes" | "hours" } = { mode: "hours" },
): string => {
  const timeString = new Intl.DateTimeFormat("de-DE", {
    timeZone: "UTC",

    ...(options.mode === "hours" && {
      timeStyle: "short",
    }),

    ...(options.mode === "minutes" && {
      minute: "2-digit",
      second: "2-digit",
    }),
  }).format(new Date(seconds * 1000));

  return timeString;
};

export { formatDuration };

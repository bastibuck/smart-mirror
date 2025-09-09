const formatDuration = (
  seconds: number,
  options?: {
    showSeconds?: boolean;
    showHours?: boolean;
  },
): string => {
  const showHours = options?.showHours ?? true;
  const showSeconds = options?.showSeconds ?? true;

  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const secs = Math.floor(seconds % 60);

  const parts: string[] = [];

  if (showHours) {
    parts.push(hours.toString().padStart(2, "0"));
  }

  parts.push(minutes.toString().padStart(2, "0"));

  if (showSeconds) {
    parts.push(secs.toString().padStart(2, "0"));
  }

  return parts.join(":");
};

export { formatDuration };

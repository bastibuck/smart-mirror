import { useNavigate, useRouter, useRouterState } from "@tanstack/react-router";
import { useCallback, useEffect } from "react";
import { toast } from "sonner";

const useSlidingNavigate = () => {
  const navigate = useNavigate();

  const slidingNavigate = useCallback(
    ({ to, direction }: { to: string; direction: "next" | "previous" }) =>
      navigate({
        to,
        viewTransition: {
          types: [direction === "next" ? "slide-left" : "slide-right"],
        },
      }),
    [navigate],
  );

  return { slidingNavigate };
};

const useAutoRotateRoutes = () => {
  const { slidingNavigate } = useSlidingNavigate();

  const { location } = useRouterState();
  const { routesByPath } = useRouter();

  useEffect(() => {
    const intervalId = setInterval(
      () => {
        const currentPageNumber = parseInt(
          location.pathname.split("/").pop()!,
          10,
        );
        if (typeof currentPageNumber !== "number" || isNaN(currentPageNumber)) {
          return;
        }

        const targetPath = `/${currentPageNumber + 1}`;

        const allRoutes = new Set(Object.keys(routesByPath));
        if (allRoutes.has(targetPath)) {
          slidingNavigate({
            to: targetPath,
            direction: "next",
          });
          return;
        }

        slidingNavigate({
          to: "/1",
          direction: "previous",
        });
      },
      1000 * 30, // 30 seconds,
    );

    return () => {
      clearInterval(intervalId);
    };
  }, [location.pathname, routesByPath, slidingNavigate]);
};

const useDirectionalNavigate = () => {
  const { location } = useRouterState();
  const { routesByPath } = useRouter();

  const { slidingNavigate } = useSlidingNavigate();

  return {
    directionalNavigate: useCallback(
      (direction: "next" | "previous") => {
        const currentPageNumber = parseInt(
          location.pathname.split("/").pop()!,
          10,
        );
        if (typeof currentPageNumber !== "number" || isNaN(currentPageNumber)) {
          return;
        }

        const targetPageNumber =
          direction === "next" ? currentPageNumber + 1 : currentPageNumber - 1;

        const to = `/${targetPageNumber}`;

        const allRoutes = new Set(Object.keys(routesByPath));
        if (!allRoutes.has(to)) {
          toast.error(`No ${direction} page found`, {
            className: "!bg-destructive !text-foreground",
          });
          return;
        }

        slidingNavigate({
          to,
          direction,
        });
      },
      [location.pathname, routesByPath, slidingNavigate],
    ),
  };
};

export { useAutoRotateRoutes, useDirectionalNavigate };

import React from "react";
import WidgetPositioner from "../_layout/WidgetPositioner";
import { useQuery } from "@tanstack/react-query";
import { fetchUtil } from "@/lib/api";
import { z } from "zod/v4";

const DailyRecipesSchema = z
  .object({
    title: z.string(),
  })
  .array();

const DailyRecipes: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const { data, isPending, isError, error } = useQuery({
    queryKey: ["recipes", "daily"],
    queryFn: () => fetchUtil("/recipes/daily", DailyRecipesSchema),
  });

  if (isError) {
    return (
      <WidgetPositioner {...widgetPositionerProps}>
        <p>{error.message}</p>
      </WidgetPositioner>
    );
  }

  if (isPending) {
    return <WidgetPositioner {...widgetPositionerProps} />;
  }

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <div className="space-y-1">
        <div className="text-3xl font-semibold">Kptn Cook</div>

        {data.map((recipe) => (
          <div className="text-lg" key={recipe.title}>
            {recipe.title}
          </div>
        ))}
      </div>
    </WidgetPositioner>
  );
};

export default DailyRecipes;

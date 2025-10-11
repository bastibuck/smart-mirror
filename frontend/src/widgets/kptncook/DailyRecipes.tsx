import React, { useEffect, useState } from "react";
import WidgetPositioner from "../_layout/WidgetPositioner";
import { useQuery } from "@tanstack/react-query";
import { fetchUtil } from "@/lib/api";
import { z } from "zod/v4";
import { HeartIcon } from "lucide-react";

const DailyRecipesSchema = z
  .object({
    title: z.string(),
    favoriteCount: z.number().nonnegative().int(),
    imageUrl: z.url(),
  })
  .array();

const DailyRecipes: React.FC<React.ComponentProps<typeof WidgetPositioner>> = ({
  ...widgetPositionerProps
}) => {
  const [imageIndex, setImageIndex] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setImageIndex((prevIndex) => (prevIndex + 1) % 3); // Assuming there are 3 images
    }, 10_000); // time of screen being displayed divided by number of images

    return () => clearInterval(interval);
  }, []);

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

  const recipe = data[imageIndex];

  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <div className="space-y-2">
        <figure className="relative inline-block">
          <img src={recipe.imageUrl} alt={recipe.title} />
          <div
            className="absolute inset-0 flex h-full w-full items-end bg-gradient-to-b from-transparent via-transparent to-black"
            key={recipe.title}
          >
            <div className="space-y-2">
              <div className="text-2xl leading-tight text-shadow-black text-shadow-xs">
                {recipe.title}
              </div>

              <div className="flex items-center justify-end gap-2 font-mono text-base">
                {Intl.NumberFormat().format(recipe.favoriteCount)}
                <HeartIcon size="1em" />
              </div>
            </div>
          </div>
        </figure>
      </div>
    </WidgetPositioner>
  );
};

export default DailyRecipes;

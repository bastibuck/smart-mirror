"use client";

import * as React from "react";
import * as CheckboxPrimitive from "@radix-ui/react-checkbox";
import { Check } from "lucide-react";

import { cn } from "~/lib/utils";

const sizesDict = {
  base: { text: "text-base", icon: "h-4 w-4" },
  "+1": { text: "text-lg", icon: "h-5 w-5" },
  "+2": { text: "text-xl", icon: "h-6 w-6" },
};

/**
 * Checkbox
 */
const Checkbox = React.forwardRef<
  React.ElementRef<typeof CheckboxPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof CheckboxPrimitive.Root> & {
    size?: keyof typeof sizesDict;
  }
>(({ className, size = "base", ...props }, ref) => (
  <CheckboxPrimitive.Root
    ref={ref}
    className={cn(
      "peer shrink-0 rounded-sm border border-primary ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground",
      sizesDict[size].icon,
      className
    )}
    {...props}
  >
    <CheckboxPrimitive.Indicator
      className={cn("flex items-center justify-center text-current")}
    >
      <Check className={sizesDict[size].icon} />
    </CheckboxPrimitive.Indicator>
  </CheckboxPrimitive.Root>
));

Checkbox.displayName = CheckboxPrimitive.Root.displayName;

/**
 * Checkbox with label
 */
const CheckboxWithLabel: React.FC<
  React.PropsWithChildren<
    React.ComponentProps<typeof Checkbox> & {
      variant?: "default" | "danger";
    }
  >
> = ({ children, variant, size = "base", ...props }) => (
  <div
    className={cn(
      "flex items-center space-x-2",
      variant === "danger" ? "text-red-400" : "text-primary",
    )}
  >
    <Checkbox
      {...props}
      size={size}
      className={cn(
        "flex items-center space-x-2",
        variant === "danger" ? "border-red-400" : "",
        props.className,
      )}
    />
    <label
      className={cn(
        "font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70",
        sizesDict[size].text,
      )}
    >
      {children}
    </label>
  </div>
);

export { Checkbox, CheckboxWithLabel };

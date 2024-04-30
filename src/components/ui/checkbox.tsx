"use client";

import * as React from "react";
import * as CheckboxPrimitive from "@radix-ui/react-checkbox";
import { Check } from "lucide-react";

import { cn } from "~/lib/utils";

const Checkbox = React.forwardRef<
  React.ElementRef<typeof CheckboxPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof CheckboxPrimitive.Root>
>(({ className, ...props }, ref) => (
  <CheckboxPrimitive.Root
    ref={ref}
    className={cn(
      "border-primary ring-offset-background focus-visible:ring-ring data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground peer h-4 w-4 shrink-0 rounded-sm border focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
      className,
    )}
    {...props}
  >
    <CheckboxPrimitive.Indicator
      className={cn("flex items-center justify-center text-current")}
    >
      <Check className="h-4 w-4" />
    </CheckboxPrimitive.Indicator>
  </CheckboxPrimitive.Root>
));
Checkbox.displayName = CheckboxPrimitive.Root.displayName;

const CheckboxWithLabel: React.FC<
  React.PropsWithChildren<
    React.ComponentProps<typeof Checkbox> & { variant?: "default" | "danger" }
  >
> = ({ children, variant, ...props }) => (
  <div
    className={cn(
      "flex items-center space-x-2",
      variant === "danger" ? "text-destructive" : "text-primary",
    )}
  >
    <Checkbox
      {...props}
      className={cn(
        "flex items-center space-x-2",
        variant === "danger" ? "border-destructive" : "",
        props.className,
      )}
    />
    <label className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
      {children}
    </label>
  </div>
);

export { Checkbox, CheckboxWithLabel };

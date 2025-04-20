import WidgetPositioner from "../_layout/WidgetPositioner";

const QuoteOfTheDay: React.FC<
  React.ComponentProps<typeof WidgetPositioner>
> = ({ ...widgetPositionerProps }) => {
  return (
    <WidgetPositioner {...widgetPositionerProps}>
      <h1 className="text-3xl font-bold">Quote of the day</h1>
    </WidgetPositioner>
  );
};

export default QuoteOfTheDay;

import { Card } from "@/components/ui/card";

interface EventStatsProps {
  eventId: string;
}

export function EventStats({ eventId }: EventStatsProps) {
  return (
    <Card className="p-6">
      <h3 className="font-semibold mb-4">Event Stats</h3>
      <div className="space-y-3">
        <div className="flex justify-between">
          <span className="text-muted-foreground">24h Volume</span>
          <span>234 INR</span>
        </div>
        <div className="flex justify-between">
          <span className="text-muted-foreground">Yes Price</span>
          <span>0.65 INR</span>
        </div>
        <div className="flex justify-between">
          <span className="text-muted-foreground">No Price</span>
          <span>0.35 INR</span>
        </div>
      </div>
    </Card>
  );
}

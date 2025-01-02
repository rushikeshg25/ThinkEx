import { Event } from "@/lib/data/event";

interface EventHeaderProps {
  event: Event;
}

export function EventHeader({ event }: EventHeaderProps) {
  return (
    <>
      <h1 className="text-2xl font-bold mb-4">{event.title}</h1>
      <div className="flex items-center gap-4 mb-6">
        <div>
          <p className="text-sm text-muted-foreground">Total Pool</p>
          <p className="font-semibold">{event.totalPool} INR</p>
        </div>
        <div>
          <p className="text-sm text-muted-foreground">Participants</p>
          <p className="font-semibold">{event.participants}</p>
        </div>
        <div>
          <p className="text-sm text-muted-foreground">Ends In</p>
          <p className="font-semibold">{event.endDate}</p>
        </div>
      </div>
    </>
  );
}

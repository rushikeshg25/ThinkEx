"use client";

import { EventCard } from "./event-card";
import { events } from "@/lib/data/event";
import { useSearchParams } from "next/navigation";

export function EventsList() {
  const searchParams = useSearchParams();
  const query = searchParams.get("q")?.toLowerCase() || "";

  const filteredEvents = events.filter((event) =>
    event.title.toLowerCase().includes(query)
  );

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {filteredEvents.map((event) => (
        <EventCard key={event.id} {...event} />
      ))}
    </div>
  );
}

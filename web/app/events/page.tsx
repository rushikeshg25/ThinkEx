import { EventsList } from "@/components/events/events-list";
import { EventSearch } from "@/components/events/event-search";
import { Suspense } from "react";

export default function EventsPage() {
  return (
    <div className="max-w-7xl mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold">Events</h1>
        <Suspense fallback={<div>Loading...</div>}>
          <EventSearch />
        </Suspense>
      </div>
      <Suspense fallback={<div>Loading...</div>}>
        <EventsList />
      </Suspense>
    </div>
  );
}

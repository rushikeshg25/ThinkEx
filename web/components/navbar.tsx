import React from "react";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { ModeToggle } from "./modetoggle";

const Navbar = () => {
  return (
    <nav className="border-b">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16 items-center">
          <div className="flex items-center">
            <Link href="/" className="text-xl font-bold">
              ThinkEx
            </Link>
          </div>

          <div className="flex items-center gap-4">
            <Link href="/events">
              <Button variant="ghost">Events</Button>
            </Link>
            <Link href="/portfolio">
              <Button variant="ghost">Portfolio</Button>
            </Link>

            <ModeToggle />
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;

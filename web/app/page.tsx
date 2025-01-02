"use client";

import { Footer } from "@/components/footer";
import { useToast } from "@/hooks/use-toast";

const page = () => {
  const { toast } = useToast();
  return (
    <div>
      Home Page
      <Footer />
    </div>
  );
};

export default page;

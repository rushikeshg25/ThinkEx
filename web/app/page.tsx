"use client";

import { Button } from "@/components/ui/button";
import { useToast } from "@/hooks/use-toast";

const page = () => {
  const { toast } = useToast();
  return (
    <div>
      page
      <Button onClick={() => toast({ title: "Hello" })}>Hello</Button>
      <a href="http://localhost:8080/auth/google"> Login </a>
    </div>
  );
};

export default page;

import Link from "next/link";

export function Footer() {
  return (
    <footer>
      <div className=" mx-auto px-4 py-12">
        <div className="mt-8 pt-8 border-t text-center text-sm text-muted-foreground">
          <p>
            © {new Date().getFullYear()} ThinkEx. All rights reserved.{" "}
            <Link
              href="https://github.com/rushikeshg25/ThinkEx"
              className="text-muted-foreground hover:text-primary underline"
            >
              Github
            </Link>
          </p>
        </div>
      </div>
    </footer>
  );
}

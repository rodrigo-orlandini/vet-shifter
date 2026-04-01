import { type SVGProps } from "react";

export function TimerIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2} strokeLinecap="round" strokeLinejoin="round" {...props}>
      <line x1="10" x2="14" y1="2" y2="2" />
      <circle cx="12" cy="14" r="8" />
      <path d="M12 10v4l2 2" />
    </svg>
  );
}

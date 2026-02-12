import React from 'react';
import NextLink from 'next/link';

interface LinkProps {
  href: string;
  children: React.ReactNode;
  className?: string;
}

export default function Link({ href, children, className = '' }: LinkProps) {
  return (
    <NextLink
      href={href}
      className={`text-teal-600 hover:text-teal-700 font-medium ${className}`}
    >
      {children}
    </NextLink>
  );
}


import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactCompiler: true,
  output: 'export',
  images: {
    unoptimized: true, // Required for Capacitor
  },
};

export default nextConfig;

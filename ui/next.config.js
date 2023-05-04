/** @type {import('next').NextConfig} */
const nextConfig = {
  headers: () => [
    {
      source: "/dashboard/employees",
      headers: [
        {
          key: "Cache-Control",
          value: "no-store",
        },
      ],
    },
  ],
  reactStrictMode: true,
  generateEtags: false,
};

module.exports = nextConfig;

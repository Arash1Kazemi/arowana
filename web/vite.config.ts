import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      // Go API server runs on :8080. This makes /api/* work in dev
      // with no CORS config and no hardcoded localhost in the frontend.
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
});

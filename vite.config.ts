import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig({
  root: "./frontend",
  plugins: [svelte()],
  server: {
    proxy: {
      "/search": "http://localhost:7500",
    },
  },
});
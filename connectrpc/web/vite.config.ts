import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    port: 3000, // or another port you prefer
  },
  build: {
    target: 'esnext',
  },
});

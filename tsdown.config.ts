import { defineConfig } from "tsdown";

export default defineConfig({
  entry: "src/index.ts",
  platform: "node",
  env: {
    NODE_ENV: "production",
  },
  shims: {
    true: false,
  },
});

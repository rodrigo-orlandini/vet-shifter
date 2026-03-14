import { defineConfig } from "orval";

export default defineConfig({
  vetShifterApi: {
    input: "../backend/cmd/api/docs/swagger.json",
    output: {
      target: "./src/api/generated/api.ts",
      client: "axios",
      mode: "single",
      override: {
        mutator: {
          path: "./src/api/axios-instance.ts",
          name: "customInstance",
        },
      },
    },
  },
});

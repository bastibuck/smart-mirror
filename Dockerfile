##### DEPENDENCIES

FROM  --platform=linux/arm64 arm64v8/node:20-alpine AS deps

WORKDIR /app

# Install Prisma Client - remove if not using Prisma

COPY prisma ./

# Install dependencies based on the preferred package manager

COPY package.json package-lock.json* ./

RUN npm ci

##### BUILDER

FROM  --platform=linux/arm64 arm64v8/node:20-alpine AS builder

# List of environment variables to be passed in from the .end.docker file
ARG DATABASE_URL
ARG TODOIST_API_TOKEN

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

ENV NEXT_TELEMETRY_DISABLED 1

RUN SKIP_ENV_VALIDATION=1 npm run build

##### RUNNER

FROM  --platform=linux/arm64 arm64v8/node:20-alpine AS runner
WORKDIR /app

ENV NODE_ENV production

ENV NEXT_TELEMETRY_DISABLED 1

COPY --from=builder /app/prisma ./
COPY --from=builder /app/next.config.js ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json

COPY --from=builder /app/entrypoint.sh ./
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static


EXPOSE 80
ENV PORT 80

RUN chmod +x /app/entrypoint.sh
ENTRYPOINT [ "/app/entrypoint.sh"]

CMD ["node", "server.js"]

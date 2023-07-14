##### DEPENDENCIES

FROM node:18-alpine AS deps
RUN apk add --no-cache libc6-compat openssl1.1-compat
WORKDIR /app

# Install Prisma Client - remove if not using Prisma

COPY prisma ./

# Install dependencies based on the preferred package manager

COPY package.json package-lock.json* ./

RUN npm ci

##### BUILDER

FROM node:18-alpine AS builder
ARG DATABASE_URL
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

ENV NEXT_TELEMETRY_DISABLED 1

RUN SKIP_ENV_VALIDATION=1 npm run build

##### RUNNER

FROM node:18-alpine AS runner
WORKDIR /app

ENV NODE_ENV production

ENV NEXT_TELEMETRY_DISABLED 1

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/prisma ./
COPY --from=builder /app/next.config.mjs ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json

COPY --from=builder --chown=nextjs:nodejs /app/entrypoint.sh ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static


USER nextjs
EXPOSE 3000
ENV PORT 3000

RUN chmod +x /app/entrypoint.sh
ENTRYPOINT [ "/app/entrypoint.sh"]

CMD ["node", "server.js"]

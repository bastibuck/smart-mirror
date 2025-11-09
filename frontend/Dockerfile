# Base stage
FROM node:24-slim AS base
RUN apt-get update && apt-get install -y ca-certificates # install SSL certificates

ARG VERSION_HASH
ENV VITE_VERSION_HASH=$VERSION_HASH

WORKDIR /app

COPY package*.json ./

RUN npm ci

# Development stage (target: dev)
FROM base AS dev

COPY . .

CMD ["npm", "run", "dev", "--", "--host"]

# Production stage (target: prod)
FROM base AS build

COPY . .

RUN npm run build

FROM nginx:alpine AS prod

COPY --from=build /app/dist /usr/share/nginx/html

# Add custom Nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

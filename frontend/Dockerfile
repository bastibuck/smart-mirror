# Base stage
FROM node:23 AS base

WORKDIR /app

COPY package*.json ./

RUN npm install

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

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

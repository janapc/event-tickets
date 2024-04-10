FROM node:18-alpine as base
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .

FROM base as production
RUN npm run build
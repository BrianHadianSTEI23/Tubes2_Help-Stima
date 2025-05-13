# Install dependencies and build the app
FROM node:18-alpine AS builder
WORKDIR /app

# Install dependencies
COPY client/package*.json ./
RUN npm install

# Copy app files and build
COPY client/. .
RUN npm run build

# Production image
FROM node:18-alpine AS runner
WORKDIR /app

# Copy only necessary files
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/.next .next
COPY --from=builder /app/public ./public
COPY --from=builder /app/node_modules ./node_modules

ENV PORT=3000
EXPOSE 3000

CMD ["npm", "start"]

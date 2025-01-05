FROM node:20.14.0-alpine AS builder

WORKDIR /app
COPY . ./
RUN npm install
RUN npm install vue-router

FROM node:20.14.0-alpine

WORKDIR /app
COPY --from=builder /app /app
EXPOSE 5173
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
# syntax=docker/dockerfile:1
# frontend
FROM node:20.17-alpine AS frontend

# Set destination for COPY
WORKDIR /web

# Download npm modules
COPY web/package.json .
COPY web/yarn.lock .
RUN yarn --frozen-lockfile

COPY web/ .

# Preload data for homepage
RUN curl https://pit-796768497423.europe-west3.run.app/api/v1/seasons/5/submissions\?page\=1\&size\=50\&class\=\&mode\=softcore\&sort\=-tier,duration -o ./src/lib/assets/preloaded.json

# Build
RUN yarn build

# backend
FROM golang:1.23.1-alpine AS api

# Set destination for COPY
WORKDIR /api

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -v -o=/bin/api ./cmd/api

FROM nginx:1.26.2-alpine-slim
WORKDIR /app
COPY --from=frontend /web/build /var/www/
COPY --from=api /bin/api /bin/api
COPY ./scripts/run.sh /run.sh
COPY ./nginx.conf ./nginx.conf

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

CMD ["/run.sh"]

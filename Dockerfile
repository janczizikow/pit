# syntax=docker/dockerfile:1
# frontend
FROM node:20.17 AS frontend

# Set destination for COPY
WORKDIR /web

# Download npm modules
COPY web/package.json .
COPY web/yarn.lock .
RUN yarn --frozen-lockfile

COPY web/ .

# Build
RUN yarn build

# backend
FROM golang:1.23.1

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# intstall go-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate $GOPATH/bin/migrate

COPY . .
COPY ./scripts/run.sh /bin/run.sh
COPY --from=frontend /web /bin/web/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -v -o=/bin/api ./cmd/api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

CMD ["sh", "/bin/run.sh"]

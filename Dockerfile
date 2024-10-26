# syntax=docker/dockerfile:1
FROM golang:1.23.1-alpine AS api

ARG TARGETOS
ARG TARGETARCH

# Set destination for COPY
WORKDIR /api

# Copy local code to the container image.
COPY . .
RUN go mod download && go mod verify

# Build
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -mod=readonly -v -o=/bin/api ./cmd/api

FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the api stage.
COPY --from=api /bin/api /bin/api
# Copy the migrations to the production image from the api stage.
COPY --from=api api/migrations/ /migrations

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run the web service on container startup.
CMD ["/bin/api"]

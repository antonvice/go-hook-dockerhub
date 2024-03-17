# Use the official Golang image to create a build artifact.
FROM golang:1.21.8 as builder

WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.mod

# Install dependencies
# This step can cache the modules as a separate layer, speeding up subsequent builds
RUN go mod download

# Copy the source code
COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# Use an Alpine image for the runtime container.
FROM alpine:latest  

WORKDIR /
# Copy the compiled binary.
COPY --from=builder /app/server /server
# Copy the update script.
COPY --from=builder /app/update_container.sh /update_container.sh
RUN chmod +x ./update_container.sh

# Install Docker CLI to interact with the host's Docker daemon.
RUN apk add --no-cache docker-cli

CMD ["/server"]

# syntax=docker/dockerfile:1

FROM golang:1.21.4

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

ENV WRAPTIME="30"

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-amigowraptime

# Run
CMD ["/docker-amigowraptime"]
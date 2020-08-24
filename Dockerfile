############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /var/www
COPY /src .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN GOOS=linux go build -ldflags="-s -w" -o ./app_build ./server.go

############################
# STEP 2 build a small image
############################
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin

# Copy our static executable.
COPY --from=builder /var/www /go/bin

# Run the binary.
ENTRYPOINT ["/go/bin/app_build"]

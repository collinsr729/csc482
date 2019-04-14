# Start the Go app build
FROM golang:latest AS build

# Copy source
WORKDIR /go/src/go-docker
COPY . .

# Get required modules (assumes packages have been added to ./vendor)
RUN go get -d -v ./...

# Build a statically-linked Go binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

# New build phase -- create binary-only image
FROM alpine:latest

# Add support for HTTPS and time zones
RUN apk update && \
    apk upgrade

WORKDIR /root/

# Copy files from previous build container
#COPY --from=build /go/src/my-golang-source-code/main ./

# Add environment variables
# ENV ...

# Check results
RUN env && pwd && find .

# Start the application
CMD ["./server"]

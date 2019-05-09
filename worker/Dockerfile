# Start the Go app build
FROM golang:latest AS build

# Copy source
WORKDIR /go/src/go-docker/http
COPY . .

# Get required modules (assumes packages have been added to ./vendor)
RUN go get -d -v ./...

# Build a statically-linked Go binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

# New build phase -- create binary-only image
FROM alpine:latest

# Add support for HTTPS and time zones
RUN apk update && \
    apk upgrade && \
    apk add ca-certificates

WORKDIR /root/

# Copy files from previous build container
COPY --from=build /go/src/go-docker/http/main ./

# Add environment variables
 ENV AWS_ACCESS_KEY_ID=AKIA34XNLPJYGHRQ6YUF
 ENV AWS_SECRET_ACCESS_KEY=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
 ENV LOGGLY_TOKEN=7653fab5-7716-4ad3-8d33-b655e4ef3ced
 ENV TOKEN=7653fab5-7716-4ad3-8d33-b655e4ef3ced

EXPOSE 8080

# Check results
RUN env && pwd && find .

# Start the application
CMD ["./main"]

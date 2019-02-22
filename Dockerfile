
FROM golang:latest


WORKDIR src/go-docker

COPY . .

RUN go get -d -v ./...

# Install the package
#RUN go install -v ./...

CMD ["go","run","http.go"]
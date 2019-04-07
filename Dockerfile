# Build Stage
FROM golang:1.12.2 as builder
LABEL maintainer="Jonathan Braley<braley.jo@gmail.com>"

WORKDIR /go/src/github.com/callicoder/go-docker

COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/go-docker .

# Image creation time.
FROM alpine:3.9.2
LABEL maintainer="Jonathan Braley<braley.jo@gmail.com>"

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/go-docker .

EXPOSE 8080
CMD ["./go-docker"]
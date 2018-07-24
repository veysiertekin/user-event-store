# Build app
FROM golang:1.10.3 as builder

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/user-event-store
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

# Create a clean image with compiled app
FROM alpine:3.8
RUN apk add --no-cache curl
COPY --from=builder /app ./
COPY docker/app-config-docker.json /app-config.json

HEALTHCHECK CMD curl --fail http://localhost:8080/health || exit 1

CMD ["./app"]

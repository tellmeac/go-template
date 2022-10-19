FROM golang:1.19.2-alpine as base

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go generate ./...
RUN go test ./...
RUN go build -o app .

FROM alpine:3

WORKDIR /app
COPY --from=base /src/app /usr/bin/app
CMD ["/usr/bin/app"]
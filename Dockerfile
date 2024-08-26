FROM golang:1.22 AS build

WORKDIR /go/src/fiber

COPY ./backend/go.mod ./backend/go.sum ./
RUN --mount=type=cache,target=/go/.cache GOCACHE=/go/.cache go mod download

COPY ./backend .

WORKDIR /go/src/fiber/cmd/server

RUN --mount=type=cache,target=/go/.cache GOCACHE=/go/.cache CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/fiber/cmd/server/app .

EXPOSE 4000

CMD ["./app"]

FROM golang:1.22 AS build

WORKDIR /go/src/fiber

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

WORKDIR /go/src/fiber/cmd/server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/fiber/cmd/server/app .

EXPOSE 4000

CMD ["./app"]

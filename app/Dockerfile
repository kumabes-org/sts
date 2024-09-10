FROM golang:1.20.2-alpine3.17 AS build

WORKDIR /app

COPY . /app/

RUN go build main.go

FROM alpine:3.17

COPY --from=build /app/main /app/main

EXPOSE 7788

WORKDIR /app

CMD [ "./main" ]
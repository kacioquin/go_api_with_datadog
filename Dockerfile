FROM golang:alpine as builder

LABEL maintainer="Cassio de Freitas <cassio.quintino@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

COPY /api .

RUN go build -o main

## Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8081

CMD ["./main"]
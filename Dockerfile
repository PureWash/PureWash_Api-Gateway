FROM golang:1.22.6 AS builder

WORKDIR /api-app

COPY . ./
RUN go mod download

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp

FROM alpine:latest

WORKDIR /api-app

COPY --from=builder /api-app/myapp .
COPY --from=builder /api-app/.env .

EXPOSE 0420

CMD [ "./myapp" ]
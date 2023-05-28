# implementing multistage to reduce image size

# Buid Stage
FROM golang:1.20.4-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run Stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .


EXPOSE 8080
CMD ["/app/main"] 
FROM golang:1.27.3-alpine AS builder

WORKDIR /app
COPY src/ .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8000
CMD [ "/app/main" ]

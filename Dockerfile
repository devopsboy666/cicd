# Build stage
FROM golang:1.20-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .
# Deploy stage
FROM debian:bullseye-slim
USER root
RUN adduser --disabled-password --gecos '' --uid 10000 --gid 10000 godev
COPY --from=build /app/app /app/app
USER godev
EXPOSE 3000
ENTRYPOINT ["/app/app"]
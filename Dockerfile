# Build stage
FROM golang:1.20-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .
# Deploy stage
FROM gcr.io/distroless/base-debian11
COPY --from=build /app/app /app/app
EXPOSE 3000
ENTRYPOINT ["/app/app"]
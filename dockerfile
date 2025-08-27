#build
FROM golang:1.22.4-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /api main.go

#prod
FROM alpine:latest
WORKDIR /
COPY --from=build /api /api
COPY ./.env .
COPY --from=build /app/database /app/database
COPY --from=build /app/models /app/models

EXPOSE 8080
CMD ["/api"]
FROM golang:1.25-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o statsync .

FROM alpine:3.21
COPY --from=build /app/statsync /usr/local/bin/statsync
EXPOSE 8080
CMD ["statsync"]

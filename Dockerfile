FROM golang:1.22.3 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-start ./cmd/server/main.go

# FROM build-stage AS run-test-stage
# RUN go test -v ./...

FROM scratch
WORKDIR /
COPY --from=build-stage /go-start /go-start
COPY --from=build-stage /app/.env .env
EXPOSE 8080
ENTRYPOINT ["/go-start"]

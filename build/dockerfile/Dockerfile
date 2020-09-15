FROM golang:1.15-alpine as build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM build_base AS server_builder
COPY . .
RUN go build -o /app ./cmd/main.go

FROM plugins/base:linux-amd64
COPY --from=server_builder /app /app
CMD /app
FROM golang:1.22.3-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/api
EXPOSE 8080
CMD ["./main"]
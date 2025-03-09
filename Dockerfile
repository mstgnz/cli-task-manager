FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o issue-tracker ./cmd

FROM alpine:latest
RUN apk --no-cache add bash
WORKDIR /root/
COPY --from=builder /app/issue-tracker /usr/local/bin/
RUN mkdir -p /root/.cli-task-manager
ENTRYPOINT ["issue-tracker"]
CMD ["help"] 
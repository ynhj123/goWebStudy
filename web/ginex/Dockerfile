FROM golang AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOPROXY=https://goproxy.cn CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/ginexample .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN addgroup -S ginexample && adduser -S ginexample -G ginexample
USER ginexample
WORKDIR /home/ginexample
COPY --from=builder /bin/ginexample ./
EXPOSE 8080
ENTRYPOINT ["./ginexample"]
FROM golang:1.14-alpine
WORKDIR /app
COPY . .
RUN apk add git --no-cache
RUN go mod download
RUN go build -o bottomsUp
CMD ["./bottomsUp"]

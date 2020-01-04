FROM golang:1.13-alpine

WORKDIR /app
COPY . .
RUN apk add git --no-cache
RUN go get github.com/gin-gonic/gin

RUN go build -o bottoms-up

EXPOSE 9000

CMD ["./bottoms-up"]

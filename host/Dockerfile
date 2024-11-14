FROM golang:alpine

WORKDIR /app
COPY . .

RUN go build -o host .

EXPOSE 8080

CMD ["./host"]

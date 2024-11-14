FROM golang:alpine

WORKDIR /app
COPY . .

RUN go build -o content .

EXPOSE 8081

CMD ["./content"]

FROM golang:latest

COPY . /app

WORKDIR /app

RUN go mod tidy

RUN go build -o app .

CMD ["/app/app"]
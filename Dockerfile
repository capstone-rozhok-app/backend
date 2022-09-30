FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o rozhok

CMD ["./rozhok"]

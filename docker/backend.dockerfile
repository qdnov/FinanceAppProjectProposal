FROM golang AS base

COPY go.mod /usr/app/go.mod

WORKDIR /usr/app

RUN go mod download

COPY . /usr/app

RUN go build -o build/main cmd/main.go

CMD ["./build/main"]
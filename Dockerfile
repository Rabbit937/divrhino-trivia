FROM golang:latest

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT ["air"]

COPY . .
RUN go mod tidy

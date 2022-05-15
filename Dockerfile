FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN go mod download


EXPOSE 8080

CMD [ "go", "run", "." ]

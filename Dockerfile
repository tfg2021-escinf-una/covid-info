FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
ADD docs ./docs

RUN go mod download


EXPOSE 8080

CMD [ "go", "run", "." ]

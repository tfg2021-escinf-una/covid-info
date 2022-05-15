FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
ADD docs ./docs

RUN go mod download

RUN go build -o /covid-info

EXPOSE 8080

CMD [ "/covid-info" ]

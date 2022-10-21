FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go build -o todo-api

EXPOSE 8080

CMD ./todo-api
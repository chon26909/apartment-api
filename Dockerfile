FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download 

RUN go build main.go

ENV ENVIRONMENT=dev

EXPOSE $PORT

ENTRYPOINT ["./main"]
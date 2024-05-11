FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download 

RUN go build main.go

EXPOSE 4000

ENTRYPOINT ["./main"]
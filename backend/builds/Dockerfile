FROM golang:1.17-alpine

RUN mkdir /hello
WORKDIR /hello

COPY go.mod .
COPY go.sum .


RUN go mod download 

RUN apk update && apk add ca-certificates

COPY . .

RUN go build -o ./main

EXPOSE 5005

CMD ["./main"]

#  docker build -t backend_dev3 -f builds/Dockerfile .
#  docker run -p 5005:5005 backend


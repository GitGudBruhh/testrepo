FROM golang:1.23rc1-alpine3.20
WORKDIR /home/workshop/server
COPY . /home/workshop/server
RUN go mod download
RUN go build -o ./build/server
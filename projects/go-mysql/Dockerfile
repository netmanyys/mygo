FROM golang:latest

WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=off
RUN go get github.com/go-sql-driver/mysql
RUN go build -o myapp

CMD ["./myapp"]

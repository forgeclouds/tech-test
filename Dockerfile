FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod init simple-api && go mod tidy

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

# docker build -t simple-go-api .
# docker run -p 8080:8080 simple-go-api
# curl localhost/health > health-check path
# curl localhost/api > endpoint path

# set the base image
FROM golang:1.23

# set working directory inside the app
WORKDIR /app


COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

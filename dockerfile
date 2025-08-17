# # set the base image
# FROM golang:1.23

# # set working directory inside the app
# WORKDIR /app


# COPY . .

# RUN go mod download

# RUN go build -o main .

# EXPOSE 8080

# CMD ["./main"]

FROM golang:1.23

WORKDIR /app

# Copy dependency files first
COPY go.mod ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the app
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

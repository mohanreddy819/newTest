# # base image
# FROM golang:1.24

# WORKDIR /app

# # Copy dependency files first
# COPY go.mod ./
# RUN go mod download

# # Copy the rest of the code
# COPY . .

# # Build the app
# RUN go build -o main .

# EXPOSE 8080

# CMD ["./main"]


# more optimized multi stage job

FROM golang:1.24-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go application into a static binary.
# CGO_ENABLED=0 is crucial for building a static binary for a minimal image like alpine.
RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Final Image 

FROM alpine:latest

# Set the working directory.
WORKDIR /app

# Copy only the compiled binary from the builder stage.
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

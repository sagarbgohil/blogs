FROM golang:1.24.4-alpine

# Set work dir
WORKDIR /app

# Install tools
# RUN apk update && apk add --no-cache git

# Copy go mod & vendor
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN go build -o blog-api ./cmd

# Expose port
EXPOSE 4000

# Run app
CMD ["./blog-api"]

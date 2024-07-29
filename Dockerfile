# Use the official Go image as the base image
FROM golang:1.22.5-alpine3.19

# Install necessary build tools and libraries
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Enable CGO and set GOOS and GOARCH
ENV CGO_ENABLED=1
ENV GOOS=linux
# ENV GOARCH=amd64

# removing apk cache
RUN rm -rf /var/cache/apk/*

# Copy the Go module files
COPY . .

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the application listens on
EXPOSE 8080

# Update the alpine
# RUN apk --update upgrade

# Add sqlite
# RUN apk add sqlite sqlite-dev

# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# Set the entry point for the container
CMD ["./main"]
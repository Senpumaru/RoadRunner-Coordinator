# Use Ubuntu 22.04 as base image
FROM golang:1.23.0-bullseye

# Set up a non-root user with sudo privileges
RUN useradd -ms /bin/bash user && \
    echo "user:user" | chpasswd && \
    usermod -aG sudo user

# Set the working directory
WORKDIR /app

# Install any necessary packages
# Install any necessary packages
RUN apt-get update && apt-get install -y ca-certificates
RUN apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN apt-get install -y git make curl

ENV GOMODCACHE=/go/pkg/mod


# Define the command to run when the container starts
CMD ["bash"]
# start from the latest golang base image
FROM golang:latest

# add Maintainer Info
LABEL maintainer="Bassem <ibassemtarek@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mods dependency requirements file
COPY go.mod .

# copy go mods expected hashes file
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy all app sources from current directory to working directory inside the container
COPY . .

# set http port
ENV PORT=8082

# Build the Go app
RUN go build -o server ./cmd/server

# remove go source files
RUN find . -name "*.go" -type f -delete

# make port 5000 available to the world outside this container
EXPOSE $PORT

# run the app
CMD ["./server"]



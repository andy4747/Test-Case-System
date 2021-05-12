FROM golang:1.16

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/angeldhakal/testcase-ms

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Build the Go app
RUN go build -o ./out/testcase-ms .

# This container exposes port 8080 to the host os
EXPOSE 8080

# Run the executable
CMD ["testcase-ms"]

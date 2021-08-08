# Use an official Go runtime as a parent image
FROM  golang:1.16
# Set the Current Working Directory inside the container
WORKDIR /app/create-token
# Copy all files from the current directory to the PWD (Present Working Directory) inside the container
COPY . .
# Download all the dependencies from go.mod file
RUN go mod download
# tidy up dependencies in go.sum file and project files
RUN go mod tidy
# Build the Go app
RUN go build -o ./app/out/create-token .
# This container exposes port 8083 to the outside world
EXPOSE 8083
# Run the binary program produced by `go build` from out folder
CMD ["./app/out/create-token"]
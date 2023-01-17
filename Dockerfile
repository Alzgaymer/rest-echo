# Start with the official MongoDB image
FROM mongo:latest

# Expose the MongoDB port
EXPOSE 27017

# Start MongoDB as the entrypoint
ENTRYPOINT ["mongod"]

# Start the Go app in a new container
FROM golang:latest

# Set the working directory
WORKDIR /go/src/app

# Copy the source files
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the app port
EXPOSE 8080

# Connect to the MongoDB container
ENV MONGO_URL=mongodb://mongo:27017

# Run the binary
CMD ["./main"]

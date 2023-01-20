# Start the Go app in a new container
FROM golang:latest

# Set the working directory
WORKDIR /go/src/app

# Copy the source files
COPY . .

#install mod dependensies
RUN go install

# Build the Go app
RUN go build -o main .

# Expose the app port
EXPOSE 8080

# Connect to the MongoDB container
ENV MONGO_URI=mongodb://mongo:27017

# Run the binary
CMD ["./main"]

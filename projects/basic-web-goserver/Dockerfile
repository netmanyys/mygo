# Start with a base Golang image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Golang code to the working directory
COPY . /app

# Build the Golang application
RUN go get -u github.com/gin-gonic/gin
RUN go build -o myapp

# Expose port 9999 for the Golang server
EXPOSE 9999


# Set the command to run the executable
CMD ["./myapp"]%
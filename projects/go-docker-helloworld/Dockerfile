# Start with a base Golang image
FROM golang:1.19-alpine

# Set the working directory inside the container 
WORKDIR /app 

# Copy the Golang code to the working directory 
COPY . /app

# Build the Golang application 
RUN go env -w GO111MODULE=off
RUN go build -o myapp

# Set the command to run the executable 
CMD ["./myapp"]
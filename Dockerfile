# Use a minimal base image for the container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

COPY ./script .

RUN apk add go
RUN go version

# To rebuild go.mod and go.sum files
RUN go clean -modcache
#RUN go mod init student-api
RUN go mod download
RUN go mod tidy

RUN go build -o student-api ./cmd

# Add execute permission to application
RUN chmod +x ./student-api

# Expose the port
EXPOSE 8080

# Command to run the application
CMD [ "./student-api" ]
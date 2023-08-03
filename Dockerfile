# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the application files to the container's working directory
COPY . .

# Build the Golang application
RUN go build -o myapp

# Set environment variables for PostgreSQL and Server configurations
ENV POSTGRES_USERNAME=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5454
ENV POSTGRES_DBNAME=postgres
ENV POSTGRES_SCHEMA=public
ENV POSTGRES_SSLMODE=disable
ENV POSTGRES_DIALECT=postgres
ENV SERVER_PORT=8085

# Expose the Server port (8085)
EXPOSE ${SERVER_PORT}

# Set the entry point for the container to run your Golang application
CMD ["./myapp"]

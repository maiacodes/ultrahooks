# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM alpine:latest

# Add Maintainer Info
LABEL maintainer="UltraHooks"
LABEL com.centurylinklabs.watchtower.enable=true

# Set the Current Working Directory inside the container
WORKDIR /app

EXPOSE 1323

# Copy binary
COPY UltraHooks.linux ./UltraHooks

# Copy public data
COPY /public ./public

# Command to run the executable
CMD ["./UltraHooks"]
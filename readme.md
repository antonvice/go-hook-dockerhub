# Docker Webhook Receiver and Updater

This Go application serves as a webhook receiver for Docker Hub notifications. Upon receiving a webhook indicating an update to a Docker repository, it automatically runs a script to update a Docker container based on the latest image.

## Overview

The application listens for POST requests on port 8080. When a webhook from Docker Hub is received, it parses the JSON payload to check the repository status. If the status is "Active", indicating a final update, it executes a predefined script (`update_container.sh`) that can pull the latest Docker image and restart the container.

## Prerequisites

- Go 1.15 or newer
- Docker installed on the host system
- The Docker CLI available in the environment where the application is running

## Setup

1. Clone the repository to your local machine.

2. Ensure `update_container.sh` script is present in the root directory of the application and is executable. This script should contain Docker commands to pull the latest image and restart the container.

    Example `update_container.sh`:

    ```bash
    #!/bin/sh
    docker pull your-image:latest
    docker stop your-container
    docker rm your-container
    docker run -d -p 8090:8080 -v /var/run/docker.sock:/var/run/docker.sock -e DOCKER_IMAGE='yourdockerhubuser/yourimage:tag' --name your-container-name your-image

    ```

3. Build the Go application:

    ```bash
    go build -o docker-webhook-receiver .
    ```

4. Run the application:

    ```bash
    ./docker-webhook-receiver
    ```

    The application will start and immediately execute `update_container.sh` to ensure the container is using the latest image. It then listens for incoming webhooks on port 8080.

## Usage

Configure your Docker Hub repository to send webhook notifications to the URL where this application is accessible, appending `/webhook` to the end of the URL.

Example: `http://your-server-address:8080/webhook`

Upon receiving a webhook with the repository status "Active", the application executes the `update_container.sh` script, pulling the latest image and updating the container accordingly.

## License

This project is open-source and available under the MIT License.

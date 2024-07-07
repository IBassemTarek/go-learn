#!/bin/bash

case $1 in
  "start")
    go run cmd/server/main.go
    ;;
  "swagger")
    swag init -g cmd/server/main.go
    ;;
  "docker-build")
    if [ "$#" -ne 3 ]; then
      echo "Usage: ./run.sh docker-build <image_name> <tag>"
      exit 1
    fi
    image_name=$2
    tag=$3
    container_name="${image_name}_container"
    
    echo "Building Docker image: ${image_name}:${tag}"
    docker build -t ${image_name}:${tag} .
    
    if [ $? -eq 0 ]; then
      echo "Docker image built successfully"
      
      # Check if the container already exists
      if [ "$(docker ps -aq -f name=^/${container_name}$)" ]; then
        echo "Container already exists. Stopping and removing..."
        docker stop ${container_name}
        docker rm ${container_name}
      fi
      
      echo "Running new container..."
      docker run -d -p 8082:8082 --name ${container_name} ${image_name}:${tag}
      
      if [ $? -eq 0 ]; then
        echo "Container is now running"
        echo "You can access it at http://localhost:8082"
        echo "To stop the container, run: docker stop ${container_name}"
      else
        echo "Failed to start the container"
      fi
    else
      echo "Docker build failed"
    fi
    ;;
  *)
    echo "Unknown command. Available commands: start, docker-build"
    exit 1
    ;;
esac
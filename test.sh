#! /bin/bash

echo "Starting local Rabbitmq container..."
if ! docker start benchmark-rabbit 2>&1  /dev/null; then
    docker run -d -p 5672:5672 --hostname my-rabbit --name benchmark-rabbit rabbitmq:3
fi

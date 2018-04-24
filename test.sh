#! /bin/bash

# this can be used for testing locally

# set up the queue
echo "Starting local Rabbitmq container..."
if ! docker start benchmark-rabbit 2>&1  /dev/null; then
    docker run -d -p 5672:5672 --hostname my-rabbit --name benchmark-rabbit rabbitmq:3
fi

# set up Vault
echo "Starting local Vault server..."
vault server -dev -dev-root

# Run the vault server
echo "Starting local Vault server..."
vault server -dev -dev-root-token-id=root &

# Set the vault address environment variable
export VAULT_ADDR='http://127.0.0.1:8200'

# enable the transit secret engine
echo "Enabling the transit engine"
vault secrets enable transit

# create an encryption key to use for transit
echo "Creating a key named \"benchmark\""
vault write -f transit/keys/benchmark


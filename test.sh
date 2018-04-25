#! /bin/bash

# this can be used for testing locally

# set up the queue
echo "Starting local Rabbitmq container..."
docker start benchmark-rabbit || docker run -d -p 5672:5672 \
    --hostname my-rabbit --name benchmark-rabbit rabbitmq:3

# set up database for results
echo "Creating local DB for to store results"
mkdir -p ~/benchmark-data
docker start mysql-benchmark || docker run --name mysql-benchmark \
    -p 3306:3306 \
    -v ~/benchmark-data:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=root \
    -e MYSQL_ROOT_HOST=% \
    -e MYSQL_DATABASE=benchmark \
    -e MYSQL_USER=vault \
    -e MYSQL_PASSWORD=vaultpw \
    -d mysql/mysql-server:5.7


# Run the vault server
echo "Starting local Vault server..."
vault server -dev -dev-root-token-id=root &

# Set the vault address environment variable
export VAULT_ADDR=http://127.0.0.1:8200

# auth
#echo "Logging in with root token"
vault login root

# enable the transit secret engine
echo "Enabling the transit engine"
vault secrets enable transit

# create an encryption key to use for transit
echo "Creating a key named \"benchmark\""
vault write -f transit/keys/benchmark
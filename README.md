# INCOMPLETE!!  

# Vault Transit Benchmarking Tool

This tool aims to demonstrate requests per second performance using the transit secret engine in Hashicorp's Vault project.

## Requirements

This test utilizes a number of components.  You need a Vault server with the transit secret engine enabled along with a key created for the test.  

```
vault server -dev -dev-root-token-id=root &
export VAULT_ADDR='http://127.0.0.1:8200'
vault secrets enable transit
vault write -f transit/keys/benchmark_key
```

A database in which to store the test results is also required.  We will use mysqldb.

```
docker pull mysql/mysql-server:5.7
mkdir ~/transit-benchmark
docker run --name mysql-transit-benchmark \
  -p 3306:3306 \
  -v ~/transit-benchmark:/var/lib/mysql \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_ROOT_HOST=% \
  -e MYSQL_DATABASE=benchmark \
  -e MYSQL_USER=vault \
  -e MYSQL_PASSWORD=vaultpw \
  -d mysql/mysql-server:5.7
```

We'll need a queue from which to pull test data:
```
docker run -d -p 5672:5672 --hostname my-rabbit --name some-rabbit rabbitmq:3
```

And a worker process to perform the writes.


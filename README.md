# T-Meter

The temperature logging service.

Demonstration project, REST API in Go with zero dependencies.

## Build project

### Local build

```bash
# Change dir to sources folder
cd ./tmeter
go build -o ../build/tmeter_server .
```

### Build Docker container

```bash
# On repository root
docker build -t tmeter -f docker/tmeter/Dockerfile .
```

or with docker-compose:

```bash
# On repository root
docker-compose build
```

## Start server

Run server on specified port:

```bash
LISTEN_HOST=0.0.0.0 LISTEN_PORT=8080 build/tmeter_server
```

### Using Docker

Run container and bind to host on port 8080:

```bash
docker run --rm -p 8080:8080 tmeter
```

or with docker-compose:

```bash
# On repository root
cp docker-compose.override.yml.dist docker-compose.override.yml 
docker-compose up
```

## Connect to service

Execute example script with specified API address:

```bash
# On repository root
cli/example_case_1.sh --server localhost:8580 --debug
```

or inside Docker-Compose project network:

```bash
docker-compose run cli example_case_1.sh --debug
```

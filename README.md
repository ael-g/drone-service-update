# Drone service update plugin
Drone plugin triggering image updates on swarm services.

## Build
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o service-update
docker build -t plugins/service-update .
```

## Run
```
docker run -it -e PLUGIN_SERVICE_NAME=your_service -v/var/run/docker.sock:/var/run/docker.sock plugins/service-update
```

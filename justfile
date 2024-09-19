POD_NAME := "unotone"
DEBUG := "1"

dev:
    air

tidy:
    go mod tidy

podman-build:
    podman build \\
        --file dockerfile \\
        --tag unotone-server:latest

run-pod: podman-build
    podman pod create \\
        --publish 8080:8080 \\
        --name {{ POD_NAME }}

    podman run \
        --pod {{ POD_NAME }} \
        --restart unless-stopped \
        --name unotone-server \
        --detach \
        -e UNOTONE_ADDR=0.0.0.0:8080 \
        -e UNOTONE_STATIC_DIR=./controllers/web/static \
        -e UNOTONE_DEBUG={{ DEBUG }} \
        unotone-server:latest

rm-pod:
    podman pod stop {{ POD_NAME }}
    podman pod rm {{ POD_NAME }}

rerun-pod: rm-pod run-pod

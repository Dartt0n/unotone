POD_NAME := "unotone"

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
        unotone-server:latest

rm-pod:
    podman pod stop {{ POD_NAME }}
    podman pod rm {{ POD_NAME }}

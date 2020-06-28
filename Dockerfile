FROM  docker.io/ubuntu:16.04
WORKDIR /hook
COPY bin/hook /bin/hook
COPY config.yaml /hook/config.yaml
ENTRYPOINT hook run
EXPOSE 9999

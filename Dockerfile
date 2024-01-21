FROM ubuntu:23.04
WORKDIR /app
COPY . .
CMD ["/bin/sh", "-c", "while true; do sleep 1000; done"]

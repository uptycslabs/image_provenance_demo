FROM ubuntu:latest
WORKDIR /app
COPY . .
CMD ["/bin/sh", "-c", "while true; do sleep 1000; done"]

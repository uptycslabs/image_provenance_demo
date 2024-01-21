FROM uptqa007/vulsecmal:v1
WORKDIR /app
COPY . .
CMD ["/bin/sh", "-c", "while true; do sleep 1000; done"]

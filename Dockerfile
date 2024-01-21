FROM uptqa007/vulsecmal:v1
RUN mkdir /tmp/app
WORKDIR /tmp/app
COPY . .
CMD ["/bin/sh", "-c", "while true; do sleep 1000; done"]

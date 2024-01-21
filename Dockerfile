FROM quay.io/petr_ruzicka/malware-cryptominer-container:1.4.0
WORKDIR /app
COPY . .
CMD ["/bin/sh", "-c", "while true; do sleep 1000; done"]

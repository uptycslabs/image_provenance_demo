FROM ubuntu:latest
WORKDIR /app
COPY . .
CMD ["sh -c", "sleep 10000"]

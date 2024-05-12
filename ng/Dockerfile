FROM ubuntu:20.04

RUN apt update && apt-get install -y wget

COPY .secret ngserver.sh /

CMD ["bash" , "/ngserver.sh"]

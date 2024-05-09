FROM balenalib/raspberrypi3-debian

RUN apt-get update && apt-get install -y golang-go gcc libc6-dev
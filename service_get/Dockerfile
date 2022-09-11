FROM golang:latest
MAINTAINER Stepan K. <xamust@gmail.com>
WORKDIR /service_get/
VOLUME ["/opt/service_get"]
COPY . ./
RUN make build
CMD [ "build/service_get" ]
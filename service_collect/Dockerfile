FROM golang:latest
MAINTAINER Stepan K. <xamust@gmail.com>
WORKDIR /service_collect/
VOLUME ["/opt/service_collect"]
COPY . ./
RUN make build
CMD [ "build/service_collect" ]
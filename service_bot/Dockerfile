FROM xamust/myalpine:golang
MAINTAINER Stepan K. <xamust@gmail.com>
WORKDIR /service_bot/
VOLUME ["/opt/service_bot"]
COPY . ./
RUN make build
CMD [ "build/service_bot" ]
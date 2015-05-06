FROM ubuntu:latest

MAINTAINER maaaato

ENV PATH $PATH:/usr/local/bin

ADD app_linux_amd64 /usr/local/bin/server

ENTRYPOINT ["/usr/local/bin/server"]

EXPOSE 8080
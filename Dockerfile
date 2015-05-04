FROM ubuntu:latest

MAINTAINER maaaato

RUN \
  apt-get update && \
  apt-get install -y nginx && \
  echo "example docker contena nginx server " > /usr/share/nginx/html/index.html


ENTRYPOINT /usr/sbin/nginx -g 'daemon off;' -c /etc/nginx/nginx.conf

EXPOSE 80
  
  
  

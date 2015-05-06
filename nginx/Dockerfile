FROM ubuntu:latest

MAINTAINER maaaato

RUN \
  apt-get update && \
  apt-get install -y nginx && \
  echo "example docker contena nginx server " > /usr/share/nginx/html/index.html

ADD ./sites-available/proxy /etc/nginx/sites-available/proxy

# delete symlink
RUN \
  rm /etc/nginx/sites-enabled/*
  
# create symlink
RUN \
  ln -s /etc/nginx/sites-available/proxy /etc/nginx/sites-enabled/proxy
  
ENTRYPOINT /usr/sbin/nginx -g 'daemon off;' -c /etc/nginx/nginx.conf

EXPOSE 80
  
  
  

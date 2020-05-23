FROM debian:stretch-slim

LABEL maintainer="Sanjeev Rohila <justsanjeev@gmail.com>"

# Install NGINX
RUN set -x \
 && apt-get update && apt-get install nginx -y

EXPOSE 80

STOPSIGNAL SIGTERM

CMD ["nginx", "-g", "daemon off;"]

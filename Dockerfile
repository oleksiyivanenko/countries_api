FROM golang:1.6-alpine

RUN apk add --update \
    nginx \
    supervisor \
  && rm -rf /var/cache/apk/*

ADD . /go/src
WORKDIR /go/src

RUN go build -o run

RUN rm /etc/nginx/nginx.conf
RUN ln -s /go/src/conf/nginx/nginx.conf /etc/nginx/
RUN ln -s /go/src/conf/nginx/prod.conf /etc/nginx/conf.d/

RUN rm /etc/supervisord.conf
RUN ln -s /go/src/conf/supervisord/supervisord.conf /etc/

CMD ["supervisord", "-n"]

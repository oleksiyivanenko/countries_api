FROM golang:1.6-alpine

RUN apk add --update \
    nginx \
    supervisor \
    py-sphinx \
    py-pip \
    py-docutils \
    git \
  && rm -rf /var/cache/apk/*

RUN pip install imagesize recommonmark
RUN go get github.com/gorilla/mux

ADD . /go/src

WORKDIR /go/src/docs
RUN make html

WORKDIR /go/src
RUN go build -o run

RUN rm /etc/nginx/nginx.conf
RUN ln -s /go/src/conf/nginx/nginx.conf /etc/nginx/
RUN ln -s /go/src/conf/nginx/prod.conf /etc/nginx/conf.d/

RUN rm /etc/supervisord.conf
RUN ln -s /go/src/conf/supervisord/supervisord.conf /etc/

CMD ["supervisord", "-n"]

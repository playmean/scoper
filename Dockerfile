FROM golang:1.14.4

ADD . /src

WORKDIR /src

RUN [ "/src/docker/build.sh" ]

ADD ./docker/config.json /data

EXPOSE 3000

CMD [ "/app/scoper", "-config", "/data/config.json" ]

FROM golang:1.14.4

ADD . /src

WORKDIR /src

RUN [ "/src/docker/build.sh" ]

ADD ./docker/config.json /app

EXPOSE 3000

CMD [ "/app/scoper", "-config", "/app/config.json" ]

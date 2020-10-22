# Build golang executable from sources

FROM golang:1.14.4 AS builder

ADD . /src

WORKDIR /src

RUN [ "/src/docker/build.sh" ]

# Build application container

FROM scratch

COPY --from=builder /app/scoper /app/scoper

ADD ./docker/config.json /data/

EXPOSE 8080

CMD [ "/app/scoper", "-config", "/data/config.json" ]

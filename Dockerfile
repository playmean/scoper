FROM scratch

WORKDIR /app

ADD ./scoper /app

EXPOSE 3000

CMD [ "/app/scoper", "-config", "/data/config.json" ]

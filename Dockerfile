FROM scratch

WORKDIR /app

ADD ./error-tracking /app

EXPOSE 3000

CMD [ "/app/error-tracking", "-config", "/data/config.json" ]

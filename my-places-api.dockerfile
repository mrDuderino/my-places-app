FROM alpine:latest

RUN mkdir /app

COPY myPlacesApp /app
COPY configs /app/configs

CMD [ "/app/myPlacesApp" ]
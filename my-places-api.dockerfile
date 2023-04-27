FROM alpine:latest

RUN mkdir /app

COPY build/myPlacesApp /app

CMD [ "/app/myPlacesApp" ]
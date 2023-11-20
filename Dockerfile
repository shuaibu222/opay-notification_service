FROM alpine:latest

RUN mkdir /app

COPY notificationService /app

CMD [ "/app/notificationService"]
FROM alpine:3.5

COPY ./microservice /app/microservice
RUN chmod +x /app/microservice

ENV PORT 9988
EXPOSE 9988

ENTRYPOINT /app/microservice

FROM alpine:3.5

COPY ./microservice-with-go /app/microservice-with-go
RUN chmod +x /app/microservice-with-go

ENV PORT 9988
EXPOSE 9988

ENTRYPOINT /app/microservice-with-go
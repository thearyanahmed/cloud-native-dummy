FROM alpine:3.5

COPY ./service /app/service
RUN chmod +x /app/service

ENV PORT 8877
EXPOSE 8877

ENTRYPOINT /app/service

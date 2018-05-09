FROM alpine:3.5

RUN apk add --no-cache ca-certificates

WORKDIR /app

EXPOSE 8888

COPY swaggerui /app/swaggerui
COPY frontend/dist /app/dist
COPY doc-server /app/doc-server

CMD ["/app/doc-server"]

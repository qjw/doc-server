FROM alpine:3.5

EXPOSE 8888

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY swaggerui /app/swaggerui
COPY frontend/dist /app/dist
COPY doc-server /app/doc-server

CMD ["/app/doc-server"]

#FROM docker.io/alpine
FROM ubuntu:16.04

WORKDIR /app

EXPOSE 8888

RUN \
	apt-get update && \
	apt-get install -y ca-certificates && \
	apt-get -y autoremove && \
	apt-get clean

ADD ./swaggerui /app/swaggerui
ADD ./frontend/dist /app/dist
ADD ./doc-server /app/doc-server

ENV LOCAL_DIR /tmp
ENV SWAGGER_UI /app/swaggerui
ENV FRONTEND /app/dist

#ARG PARAM
CMD ["/app/doc-server"]



FROM alpine:latest

ADD cmd/check/check /opt/resource/check
ADD cmd/in/in /opt/resource/in
ADD cmd/out/out /opt/resource/out

RUN chmod +x /opt/resource/*

RUN apk add bash ca-certificates --no-cache

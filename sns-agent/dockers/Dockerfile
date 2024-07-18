FROM ubuntu:18.04
LABEL maintainer="iron.zhang@foxmail.com"

WORKDIR /sns
COPY ./conf ./conf
COPY ./entrypoint.sh ./
COPY ./sns-agent ./

EXPOSE 1789

ENTRYPOINT ["./entrypoint.sh"]

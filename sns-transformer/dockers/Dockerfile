FROM ubuntu:18.04
LABEL maintainer="iron.zhang@foxmail.com"

WORKDIR /sns
COPY ./conf ./conf
COPY ./entrypoint.sh ./
COPY ./sns-transformer ./

ENTRYPOINT ["./entrypoint.sh"]

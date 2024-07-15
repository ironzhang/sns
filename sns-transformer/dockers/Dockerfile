FROM ubuntu:18.04
LABEL maintainer="iron.zhang@foxmail.com"

WORKDIR /sns
COPY ./conf ./conf
COPY ./sns-transformer ./
COPY ./entrypoint.sh ./

ENTRYPOINT ["./entrypoint.sh"]

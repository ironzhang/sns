#!/bin/bash

mkdir -p /root/.supername/bin
cp ./sns-lookup /root/.supername/bin/
./sns-agent -config ./conf/cfg.docker.toml >run.log 2>&1

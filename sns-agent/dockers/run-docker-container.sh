#!/bin/bash

docker run -d -p 1789:1789 -v $HOME/.supername/resource:/root/.supername/resource --rm --network minikube --link minikube:minikubeCA --name sns-agent sns/agent:v0.0.1

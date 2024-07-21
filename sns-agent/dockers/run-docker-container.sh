#!/bin/bash

docker run -d -p 1789:1789 -v $HOME/.supername:/root/.supername --rm --network minikube --link minikube:minikubeCA --name sns-agent sns/agent:v0.0.1

#!/bin/bash

RESOURCE=$HOME/.supername/resource

docker run -d -p 1789:1789 --name sns-agent --network minikube --link minikube:minikubeCA -v $RESOURCE:/sns/resource sns/agent:v0.0.1

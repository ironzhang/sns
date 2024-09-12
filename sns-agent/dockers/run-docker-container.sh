#!/bin/bash

docker run -d -p 1789:1789 -v $HOME/.supername/resource:/var/supername/resource --rm --network minikube --link minikube:minikubeCA --name sns-agent sns/agent

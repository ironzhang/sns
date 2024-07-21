#!/bin/bash

docker run -d -p 8000:8000 -v $HOME/.supername:/root/.supername --rm --network minikube --name sns-eg-caller sns/examples/caller

#!/bin/bash

docker run -d -p 8001:8001 --rm --network minikube --name sns-eg-callee sns/examples/callee

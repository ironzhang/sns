#!/bin/bash

docker run -d --network minikube --link minikube:minikubeCA --rm --name sns-transformer sns/transformer

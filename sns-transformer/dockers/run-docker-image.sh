#!/bin/bash

docker run -d --name sns-transformer --network minikube --link minikube:minikubeCA sns/transformer:v0.0.1

#!/bin/bash

MINIKUBE=$HOME/.minikube

cp $MINIKUBE/ca.crt ./conf/kube/
cp $MINIKUBE/profiles/minikube/client.crt ./conf/kube/
cp $MINIKUBE/profiles/minikube/client.key ./conf/kube/

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ../
docker build -t sns/transformer:v0.0.1 .

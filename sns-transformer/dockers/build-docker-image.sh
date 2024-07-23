#!/bin/bash

MINIKUBE=$HOME/.minikube

cp $MINIKUBE/ca.crt ./conf/kube/
cp $MINIKUBE/profiles/minikube/client.crt ./conf/kube/
cp $MINIKUBE/profiles/minikube/client.key ./conf/kube/

(cd .. && make linux && mv ./sns-transformer ./dockers/)
docker build -t sns/transformer .

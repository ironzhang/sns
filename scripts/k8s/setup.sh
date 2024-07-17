#!/bin/bash

function StartMinikube() {
	echo "minikube start --registry-mirror=https://docker.m.daocloud.io --image-mirror-country=cn --insecure-registry=docker-registry:5000"
	minikube start --registry-mirror=https://docker.m.daocloud.io --image-mirror-country=cn --insecure-registry=docker-registry:5000
}

function RunDockerRegistry() {
	echo "docker run -d --name docker-registry --network minikube -p 5000:5000 --restart always registry:2"
	docker run -d --name docker-registry --network minikube -p 5000:5000 --restart always registry:2
}

function BuildTransformerImage() {
	echo "(cd ../../sns-transformer/dockers && ./build-docker-image.sh)"
	(cd ../../sns-transformer/dockers && ./build-docker-image.sh)
}

function PushTransformerImage() {
	echo "docker tag sns/transformer:v0.0.1 127.0.0.1:5000/sns/transformer:v0.0.1"
	docker tag sns/transformer:v0.0.1 127.0.0.1:5000/sns/transformer:v0.0.1

	echo "docker push 127.0.0.1:5000/sns/transformer:v0.0.1"
	docker push 127.0.0.1:5000/sns/transformer:v0.0.1
}

function CreateSNSCRDs() {
	echo "(cd ../../kernel/artifacts && kubectl apply -f snsclusters.core.sns.io.yaml)"
	(cd ../../kernel/artifacts && kubectl apply -f snsclusters.core.sns.io.yaml)

	echo "(cd ../../kernel/artifacts && kubectl apply -f sns.namespace.yaml)"
	(cd ../../kernel/artifacts && kubectl apply -f sns.namespace.yaml)

	echo "(cd ../../kernel/artifacts && kubectl apply -f sns.roles.yaml)"
	(cd ../../kernel/artifacts && kubectl apply -f sns.roles.yaml)
}

function RunTransformerPod() {
	echo "kubectl apply -f transformer-deployment.yaml"
	kubectl apply -f transformer-deployment.yaml
}

function Setup() {
	StartMinikube
	CreateSNSCRDs
	RunDockerRegistry

	BuildTransformerImage
	PushTransformerImage
	RunTransformerPod
}

function Clean() {
	echo "minikube delete"
	minikube delete

	echo "docker stop docker-registry"
	docker stop docker-registry

	echo "docker rm docker-registry"
	docker rm docker-registry
}

function main() {
	local action=$1
	case X$action in
		"XStartMinikube" )
			StartMinikube
			;;
		"XCreateSNSCRDs" )
			CreateSNSCRDs
			;;
		"XRunDockerRegistry" )
			RunDockerRegistry
			;;
		"XBuildTransformerImage" )
			BuildTransformerImage
			;;
		"XPushTransformerImage" )
			PushTransformerImage
			;;
		"XRunTransformerPod" )
			RunTransformerPod
			;;
		"XClean" )
			Clean
			;;
		* )
			Setup
			;;
	esac
}

main $@

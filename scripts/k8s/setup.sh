#!/bin/bash

function StartMinikube() {
	echo "minikube start --registry-mirror=https://docker.m.daocloud.io --image-mirror-country=cn --insecure-registry=docker-registry:5000"
	minikube start --registry-mirror=https://docker.m.daocloud.io --image-mirror-country=cn --insecure-registry=docker-registry:5000
}

function CreateSNSCRDs() {
	echo "(cd ../../kernel/artifacts && kubectl apply -f snsclusters.core.sns.io.yaml)"
	(cd ../../kernel/artifacts && kubectl apply -f snsclusters.core.sns.io.yaml)

	echo "(cd ../../kernel/artifacts && kubectl apply -f sns.namespace.yaml)"
	(cd ../../kernel/artifacts && kubectl apply -f sns.namespace.yaml)

	echo "(cd ../../kernel/artifacts && kubectl apply -f sns.roles.yaml)"
	(cd ../../kernel/artifacts && kubectl apply -f sns.roles.yaml)
}

function RunDockerRegistry() {
	echo "docker run -d --name docker-registry --network minikube -p 5000:5000 --restart always registry:2"
	docker run -d --name docker-registry --network minikube -p 5000:5000 --restart always registry:2
}

function BuildTransformerImage() {
	echo "(cd ../../sns-transformer/dockers && ./build-docker-image.sh)"
	(cd ../../sns-transformer/dockers && ./build-docker-image.sh)

	echo "docker tag sns/transformer:v0.0.1 127.0.0.1:5000/sns/transformer:v0.0.1"
	docker tag sns/transformer:v0.0.1 127.0.0.1:5000/sns/transformer:v0.0.1

	echo "docker push 127.0.0.1:5000/sns/transformer:v0.0.1"
	docker push 127.0.0.1:5000/sns/transformer:v0.0.1
}

function BuildAgentImage() {
	echo "(cd ../../sns-agent/dockers && ./build-docker-image.sh)"
	(cd ../../sns-agent/dockers && ./build-docker-image.sh)

	echo "docker tag sns/agent:v0.0.1 127.0.0.1:5000/sns/agent:v0.0.1"
	docker tag sns/agent:v0.0.1 127.0.0.1:5000/sns/agent:v0.0.1

	echo "docker push 127.0.0.1:5000/sns/agent:v0.0.1"
	docker push 127.0.0.1:5000/sns/agent:v0.0.1
}

function RunTransformerPod() {
	echo "kubectl apply -f transformer-deployment.yaml"
	kubectl apply -f transformer-deployment.yaml
}

function BuildExampleImages() {
	echo "(cd ../../examples/callee/dockers && ./build-docker-image.sh)"
	(cd ../../examples/callee/dockers && ./build-docker-image.sh)

	echo "docker tag sns/examples/callee:latest 127.0.0.1:5000/sns/examples/callee:latest"
	docker tag sns/examples/callee:latest 127.0.0.1:5000/sns/examples/callee:latest

	echo "docker push 127.0.0.1:5000/sns/examples/callee:latest"
	docker push 127.0.0.1:5000/sns/examples/callee:latest

	echo "(cd ../../examples/caller/dockers && ./build-docker-image.sh)"
	(cd ../../examples/caller/dockers && ./build-docker-image.sh)

	echo "docker tag sns/examples/caller:latest 127.0.0.1:5000/sns/examples/caller:latest"
	docker tag sns/examples/caller:latest 127.0.0.1:5000/sns/examples/caller:latest

	echo "docker push 127.0.0.1:5000/sns/examples/caller:latest"
	docker push 127.0.0.1:5000/sns/examples/caller:latest
}

function RunExamplePods() {
	echo "kubectl apply -f callee-deployment.yaml"
	kubectl apply -f callee-deployment.yaml

	echo "kubectl apply -f caller-deployment.yaml"
	kubectl apply -f caller-deployment.yaml
}

function DebugUsage() {
	echo "setup.sh debug examples:"
	echo "	./setup.sh debug StartMinikube"
	echo "	./setup.sh debug CreateSNSCRDs"
	echo "	./setup.sh debug RunDockerRegistry"
	echo "	./setup.sh debug BuildTransformerImage"
	echo "	./setup.sh debug BuildAgentImage"
	echo "	./setup.sh debug RunTransformerPod"
	echo "	./setup.sh debug BuildExampleImages"
	echo "	./setup.sh debug RunExamplePods"
	echo "	./setup.sh debug Clean"
}

function Debug() {
	local funcName=$1
	case X$funcName in
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
		"XBuildAgentImage" )
			BuildAgentImage
			;;
		"XRunTransformerPod" )
			RunTransformerPod
			;;
		"XBuildExampleImages" )
			BuildExampleImages
			;;
		"XRunExamplePods" )
			RunExamplePods
			;;
		"XClean" )
			Clean
			;;
		* )
			echo "Unknown function: $funcName"
			DebugUsage
			;;
	esac
}

function Clean() {
	echo "minikube delete"
	minikube delete

	echo "docker stop docker-registry"
	docker stop docker-registry

	echo "docker rm docker-registry"
	docker rm docker-registry
}

function Setup() {
	StartMinikube
	CreateSNSCRDs
	RunDockerRegistry

	BuildTransformerImage
	BuildAgentImage

	RunTransformerPod
}

function SetupExamples() {
	BuildExampleImages
	RunExamplePods
}

function Help() {
	echo "Usage of ./setup.sh"
	echo "	./setup.sh"
	echo "	./setup.sh clean"
	echo "	./setup.sh examples"
}

function main() {
	local action=$1
	case X$action in
		"Xdebug" )
			Debug $2
			;;
		"Xclean" )
			Clean
			;;
		"Xhelp" )
			Help
			;;
		"Xexamples" )
			SetupExamples
			;;
		* )
			Setup
			;;
	esac
}

main $@

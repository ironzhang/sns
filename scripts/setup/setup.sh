#!/bin/bash

function kubectl() {
	minikube kubectl -- $@
}

function StartMinikube() {
	echo "minikube start --registry-mirror=https://docker.m.daocloud.io --image-mirror-country=cn --insecure-registry=docker-registry:5000"
	minikube start --registry-mirror=https://docker.m.daocloud.io --image-mirror-country=cn --insecure-registry=docker-registry:5000
}

function CreateSNSCRDs() {
	echo "(cd ../../kernel/crds && kubectl apply -f core.sns.io_snsclusters.yaml)"
	(cd ../../kernel/crds && kubectl apply -f core.sns.io_snsclusters.yaml)

	echo "(cd ../../kernel/crds && kubectl apply -f core.sns.io_snsroutepolicies.yaml)"
	(cd ../../kernel/crds && kubectl apply -f core.sns.io_snsroutepolicies.yaml)

	echo "(cd ../../kernel/artifacts && kubectl apply -f sns.namespace.yaml)"
	(cd ../../kernel/artifacts && kubectl apply -f sns.namespace.yaml)

	echo "(cd ../../kernel/artifacts && kubectl apply -f sns.roles.yaml)"
	(cd ../../kernel/artifacts && kubectl apply -f sns.roles.yaml)
}

function RunDockerRegistry() {
	echo "docker run -d --name docker-registry --network minikube -p 5000:5000 --restart always registry:2"
	docker run -d --name docker-registry --network minikube -p 5000:5000 --restart always registry:2
}

function InstallTools() {
	echo "go install ../../tools/sns-lookup"
	go install ../../tools/sns-lookup
}

function BuildTransformerImage() {
	echo "(cd ../../sns-transformer/dockers && ./build-docker-image.sh)"
	(cd ../../sns-transformer/dockers && ./build-docker-image.sh)

	echo "docker tag sns/transformer:latest 127.0.0.1:5000/sns/transformer:latest"
	docker tag sns/transformer:latest 127.0.0.1:5000/sns/transformer:latest

	echo "docker push 127.0.0.1:5000/sns/transformer:latest"
	docker push 127.0.0.1:5000/sns/transformer:latest
}

function BuildAgentImage() {
	echo "(cd ../../sns-agent/dockers && ./build-docker-image.sh)"
	(cd ../../sns-agent/dockers && ./build-docker-image.sh)

	echo "docker tag sns/agent:latest 127.0.0.1:5000/sns/agent:latest"
	docker tag sns/agent:latest 127.0.0.1:5000/sns/agent:latest

	echo "docker push 127.0.0.1:5000/sns/agent:latest"
	docker push 127.0.0.1:5000/sns/agent:latest
}

function RunTransformerPod() {
	echo "kubectl apply -f transformer-deployment.yaml"
	kubectl apply -f transformer-deployment.yaml
}

function RunAgentDaemonSet() {
	echo "kubectl apply -f agent-daemonset.yaml"
	kubectl apply -f agent-daemonset.yaml
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
	echo "Usage of setup.sh debug:"
	echo "	./setup.sh debug StartMinikube"
	echo "	./setup.sh debug CreateSNSCRDs"
	echo "	./setup.sh debug RunDockerRegistry"
	echo "	./setup.sh debug InstallTools"
	echo "	./setup.sh debug BuildTransformerImage"
	echo "	./setup.sh debug BuildAgentImage"
	echo "	./setup.sh debug RunTransformerPod"
	echo "	./setup.sh debug RunAgentDaemonSet"
	echo "	./setup.sh debug BuildExampleImages"
	echo "	./setup.sh debug RunExamplePods"
}

function DebugFunc() {
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
		"XInstallTools" )
			InstallTools
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
		"XRunAgentDaemonSet" )
			RunAgentDaemonSet
			;;
		"XBuildExampleImages" )
			BuildExampleImages
			;;
		"XRunExamplePods" )
			RunExamplePods
			;;
		* )
			echo "Unknown function: $funcName"
			;;
	esac
}

function Debug() {
	if [ $# -eq 0 ]; then
		DebugUsage
		exit 0
	fi

	for arg in $@; do
		echo "debug $arg"
		DebugFunc $arg
	done
}

function Init() {
	StartMinikube
	CreateSNSCRDs
	RunDockerRegistry
	InstallTools

	BuildTransformerImage
	BuildAgentImage

	RunTransformerPod
	RunAgentDaemonSet
}

function Clean() {
	echo "minikube delete"
	minikube delete

	echo "docker stop docker-registry"
	docker stop docker-registry

	echo "docker rm docker-registry"
	docker rm docker-registry
}

function SetupExamples() {
	BuildExampleImages
	RunExamplePods
}

function Help() {
	echo "Usage of ./setup.sh"
	echo "	./setup.sh init"
	echo "	./setup.sh clean"
	echo "	./setup.sh examples"
}

function main() {
	local action=$1
	case X$action in
		"Xdebug" )
			shift
			Debug $@
			;;
		"Xinit" )
			Init
			;;
		"Xclean" )
			Clean
			;;
		"Xexamples" )
			SetupExamples
			;;
		* )
			Help
			;;
	esac
}

main $@

#!/bin/bash

function GetTransformerImagePath() {
	if [ -z "$1" ]; then
		echo "registry.cn-hangzhou.aliyuncs.com/ironzhang/sns-transformer:latest"
	else
		echo $1
	fi
}

function GenerateSNSCRDFiles() {
	cp ../../kernel/crds/core.sns.io_snsclusters.yaml ./prod/
	cp ../../kernel/crds/core.sns.io_snsroutepolicies.yaml ./prod/
	cp ../../kernel/artifacts/sns.namespace.yaml ./prod/
	cp ../../kernel/artifacts/sns.roles.yaml ./prod/
}

function GenerateTransformerDeploymentFile() {
	local transformerImage=$1
	cat <<'EOF' | sed "s#TRANSFORMER_IMAGE#$transformerImage#g" >./prod/transformer-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sns-transformer-deployment
  labels:
    app: sns-transformer-deployment-v1
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sns-transformer
  template:
    metadata:
      labels:
        app: sns-transformer
    spec:
      containers:
      - name: transformer
        image: TRANSFORMER_IMAGE
EOF
}

function GenerateAgentDaemonSetFile() {
	cp ./agent-daemonset.yaml ./prod/
}

function GenerateSetupFile() {
	cat <<'EOF' > ./prod/setup.sh
kubectl apply -f core.sns.io_snsclusters.yaml
kubectl apply -f core.sns.io_snsroutepolicies.yaml
kubectl apply -f sns.namespace.yaml
kubectl apply -f sns.roles.yaml
kubectl apply -f transformer-deployment.yaml
kubectl apply -f agent-daemonset.yaml
EOF
	chmod u+x ./prod/setup.sh
}

function main() {
	local transformerImage=$(GetTransformerImagePath $1)

	mkdir -p prod
	GenerateSNSCRDFiles
	GenerateTransformerDeploymentFile $transformerImage
	GenerateAgentDaemonSetFile
	GenerateSetupFile
}

main $@

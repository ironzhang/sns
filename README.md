# SNS

## Overview

SNS(super-name-system) is a service discovery system similar to DNS. Users can use the SDKs to resolve domain names into IP:Port. The difference from DNS in that:

* SNS supports resolve Port.
* SNS provides flexible traffic scheduling capabilities.
* SNS supports users to define their own load balancing algorithms.

> the domain name in SNS is a specific concept, similar to the domain name in DNS, but not exactly the same.

## Quick Start

### Requirements

* go version >= 1.22.3
* A working docker environment

### Installation

step 1: use minikube to start k8s server
```
minikube start
```

step 2: create sns CRD
```
curl https://raw.githubusercontent.com/ironzhang/sns/master/kernel/artifacts/snsclusters.core.sns.io.yaml >snsclusters.core.sns.io.yaml
kubectl apply -f snsclusters.core.sns.io.yaml
```

step 3: install sns services
```
go install github.com/ironzhang/sns/sns-agent@latest
go install github.com/ironzhang/sns/sns-transformer@latest
go install github.com/ironzhang/supernamego/examples/sns-lookup@latest
```

step 4: start sns services
```
mkdir -p sns-transformer; cd sns-transformer; nohup sns-transformer >run.log &; cd ..;
mkdir -p sns-agent; cd sns-agent; nohup sns-agent >run.log &; cd ..;
```

### Usage

first, we create some k8s pods
```
kubectl apply -f deployment.yaml
```

The deployment.yaml is as follows:
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: zone00-myapp-deployment
  labels:
    app: zone00-myapp-deployment-v1
spec:
  replicas: 1
  selector:
    matchLabels:
      app: zone00.myapp
  template:
    metadata:
      labels:
        app: zone00.myapp
    spec:
      containers:
      - name: nginx
        image: nginx:1.22
        ports:
        - name: http
          containerPort: 80
        - name: https
          containerPort: 443
```

then, we can use sns-lookup to resolve the domains
```
sns-lookup sns/http.myapp
sns-lookup sns/https.myapp
```

we can use the SDK to resolve the domains too, see [supernamego](https://github.com/ironzhang/supernamego?tab=readme-ov-file#supernamego)


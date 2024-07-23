English | [中文](./README_CN.md)

# SNS

## Overview

SNS(super-name-system) is a DNS-like product developed for intranet service discovery. Users can use the SDKs to resolve domain names into IP:Port. The difference from DNS is SNS provides flexible traffic scheduling capabilities. Base on this, users can quickly build functions such as Sentinel-Stress-Testing, Blue-Green-Deployment. 

> the domain name in SNS is a specific concept, similar to the domain name in DNS, but not exactly the same.

## Quick Start

### Requirements

* go version >= 1.22.3
* A working docker environment

### Installation

step 1: setup
```
git clone git@github.com:ironzhang/sns.git
(cd sns/scripts/k8s && ./setup.sh init)
```

step 2: run sns-agent
```
(cd sns/sns-agent && make && ./sns-agent)
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
  name: k8s-myapp-deployment
  labels:
    app: k8s-myapp-deployment-v1
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp
      cluster: k8s
  template:
    metadata:
      labels:
        app: myapp
        cluster: k8s
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

we can use the SDK to resolve the domains too, see [supernamego](https://github.com/ironzhang/supernamego?tab=readme-ov-file#supernamego).


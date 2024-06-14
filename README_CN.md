[English](./README.md) | 中文

# SNS

## 概述

SNS(super-name-system) 是一个类似 DNS 的服务发现系统。用户可以使用 sns SDK 将域名解析为 IP:Port。不同于 DNS 之处在于：

* SNS 支持解析端口。
* SNS 提供了灵活的流量调度能力。
* SNS 支持用户自定义负载均衡算法。

> SNS 中的域名是一个特定概念，类似 DNS 的域名，但并不完全相同。

## 快速开始

### 要求

* go 版本 >= 1.22.3
* 一个可用的 docker 环境

### 安装

步骤1：使用 minikube 启动 k8s 服务
```
minikube start --registry-mirror=https://registry.docker-cn.com --image-mirror-country=cn
```

步骤2：创建 sns CRD
```
curl https://raw.githubusercontent.com/ironzhang/sns/master/kernel/artifacts/snsclusters.core.sns.io.yaml >snsclusters.core.sns.io.yaml
kubectl apply -f snsclusters.core.sns.io.yaml
```

步骤3：安装 sns 服务
```
go install github.com/ironzhang/sns/sns-agent@latest
go install github.com/ironzhang/sns/sns-transformer@latest
go install github.com/ironzhang/supernamego/examples/sns-lookup@latest
```

步骤4：启动 sns 服务
```
mkdir -p sns-transformer; cd sns-transformer; nohup sns-transformer >run.log &; cd ..;
mkdir -p sns-agent; cd sns-agent; nohup sns-agent >run.log &; cd ..;
```

### 用法

首先，我们需要创建一些 k8s pod
```
kubectl apply -f deployment.yaml
```

deployment.yaml 文件内容如下：
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

然后，我们可以使用 sns-lookup 来解析域名
```
sns-lookup sns/http.myapp
sns-lookup sns/https.myapp
```

也可以使用 SDK 来解析域名，参见 [supernamego](https://github.com/ironzhang/supernamego?tab=readme-ov-file#supernamego)。


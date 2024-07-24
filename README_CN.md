[English](./README.md) | 中文

# SNS

## 概述

SNS（super-name-system）是一款专为内网服务发现而研发的类 DNS 产品。用户可以使用 SNS SDK 将域名解析为 IP:Port。相比 DNS，它提供了非常灵活的流量调度能力，基于此，用户可以快速构建出哨兵压测、蓝绿发布等功能。

> SNS 中的域名是一个特定概念，类似 DNS 的域名，但并不完全相同。

## 快速开始

### 要求

* go 版本 >= 1.22.3
* 一个可用的 docker 环境

### 安装

步骤1：初始化安装
```
git clone git@github.com:ironzhang/sns.git
(cd sns/scripts/setup && ./setup.sh init)
```

步骤2：运行 sns-agent
```
(cd sns/sns-agent && make && ./sns-agent)
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

然后，我们可以使用 sns-lookup 来解析域名
```
sns-lookup sns/http.myapp
sns-lookup sns/https.myapp
```

也可以使用 SDK 来解析域名，参见 [supernamego](https://github.com/ironzhang/supernamego?tab=readme-ov-file#supernamego)。


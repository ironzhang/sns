# SNS 总体设计

[toc]

## 1. 概述

### 1.1. 什么是 SNS？

SNS 是 super name system 的简称，是一个类似 DNS（domain-name-system）的服务发现系统。它提供了若干语言的 SDK：

* [supernamego](https://github.com/ironzhang/supernamego)
* supernamec(TODO)
* supernamejava(TODO)
* supernamepython(TODO)
* more

用户可以通过 SDK 将一个 SNS 域名解析为一个 IP:Port，以此达到服务发现的目的。

### 1.2. 为什么要做 SNS？

在互联网环境下，DNS 是一个非常成功的服务发现系统，但因其配置生效慢，不支持端口等特点，DNS 不是内网服务发现的最佳选择。SNS 的目标是构建一个专用于内网的、类似 DNS 的、具备灵活的流量调度能力的服务发现系统，帮助开发者更快速、更容易地构建和交付微服务。

### 1.3. SNS 特性概览

SNS 支持特性如下图所示：

![](./diagram/sns-features.svg)

图1 SNS 特性概览

## 2. 总体设计

### 2.1. 概念和术语

|中文|英文|释义|
|----|----|----|
|应用|application|应用代表一个可被部署的程序|
|集群|cluster|一个应用可按需求部署多个集群，以实现小流量、多活等能力|
|地址节点|endpoint|一个地址节点代表一个可以被访问的服务实例|
|SNS 域名|SNS domain name|可以被 SNS SDK 解析为 IP:Port 的一串字符串，如 sns/http.myapp|

![](./diagram/svc-model.png)

图2 SNS 服务模型

如上图所示，一个应用可通过多个不同的域名来对外提供不同的服务，而一个域名下又可分为多个集群，如小流量集群、线上生产环境集群等，而一个集群下则包含多个地址节点。

### 2.2. 逻辑架构

SNS 系统的逻辑架构如下图所示：

![](./diagram/architecture.png)

图3 SNS 系统逻辑架构

各模块的职责如下表所示：

|模块|职责|
|----|----|
|k8s-api-server|作为数据存储中心使用，所有服务发现数据都存储在 k8s-api-server。|
|sns-transformer|转换程序，负责将 pod 等 k8s 原生资源转成供 sns-agent 订阅的自定义资源。|
|sns-agent|代理程序，负责订阅域名并更新到本机文件，以供 SDK 读取。|
|supernamec/supernamego/supernamejava|各语言 SDK，为用户提供将域名解析为 IP:Port 的功能。|

## 3. 开发路线图

![](./diagram/roadmap.svg)

图4 SNS 开发路线路


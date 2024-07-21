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
go install github.com/ironzhang/supernamego/examples/sns-lookup@latest
git clone git@github.com:ironzhang/sns.git
(cd sns/scripts/k8s && ./setup.sh && ./setup.sh examples)
```

step 2: run sns-agent
```
(cd sns/sns-agent && go build && ./sns-agent)
```

### Usage

now we can use sns-lookup to resolve the domains
```
sns-lookup sns/http.callee
sns-lookup sns/http.caller
```

we can use the SDK to resolve the domains too, see [supernamego](https://github.com/ironzhang/supernamego?tab=readme-ov-file#supernamego).


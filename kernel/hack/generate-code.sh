#!/bin/bash

CONTROLLER_GEN="controller-gen"
GENERATE_GROUPS="/Users/iron/workspace/src/golang/github.com/code-generator/generate-groups.sh"

$CONTROLLER_GEN crd:crdVersions=v1 paths=../apis/core.sns.io/v1 output:dir=../crds
$GENERATE_GROUPS all github.com/ironzhang/sns/kernel/clients/coresnsclients github.com/ironzhang/sns/kernel/apis core.sns.io:v1 --go-header-file ./boilerplate.go.txt


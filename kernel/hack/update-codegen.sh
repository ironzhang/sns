#!/bin/bash

GENERATE_GROUPS="/Users/iron/workspace/src/golang/github.com/code-generator/generate-groups.sh"

$GENERATE_GROUPS all github.com/ironzhang/sns/kernel/clients/coresnsclients github.com/ironzhang/sns/kernel/apis core.sns.io:v1 --go-header-file ./boilerplate.go.txt


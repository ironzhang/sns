#!/bin/bash

function GenerateSupernameConf() {
	local hostIP=$(hostname -i)
	cat <<'EOF' | sed "s#HOST_IP#$hostIP#g" > /var/supername/supername.conf
[Agent]
	Server = "HOST_IP:1789"
	SkipError = true
	Timeout = 2
	KeepAliveTTL = 600
	KeepAliveInterval = 10

[Watch]
	ResourcePath = "/var/supername/resource"
	WatchInterval = 1
EOF
}

function Setup() {
	mkdir -p /var/supername/bin
	cp ./sns-lookup /var/supername/bin/
	GenerateSupernameConf
}

Setup

./sns-agent -config ./conf/cfg.docker.toml >run.log 2>&1

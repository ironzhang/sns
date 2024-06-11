#!/bin/bash

Addr="127.0.0.1:1789"

curl -X GET "http://$Addr/sns/agent/api/v1/list/watch/domains"

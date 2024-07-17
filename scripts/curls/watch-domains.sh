#!/bin/bash

Addr="127.0.0.1:1789"

ContentTypeHeader='Content-Type: application/json'

Data='{
	"Domains": ["sns.https.nginx"],
	"TTL": "0s"
}'

curl -X POST "http://$Addr/sns/agent/api/v1/watch/domains" --header "${ContentTypeHeader}" -d "$Data"

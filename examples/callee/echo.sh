#!/bin/bash

curl -X POST http://127.0.0.1:8001/echo --header 'Content-Type: application/json' -d '"hello, world"'

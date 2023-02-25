#!/bin/bash

mkdir -p output
go build && mv gpt3-http-server ./output/gpt3-http-server && cp config.json ./output/config.json && mv run.sh ./output/run.sh
echo "./gpt3-http-server > out.log 2>&1" > run.sh
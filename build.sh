# !/bin/bash
if [ !-d "$output"];
then
  mkdir output
fi

go build && mv gpt3-http-server ./output/gpt3-http-server
#!/bin/bash

gox -output "dist/{{.OS}}_{{.Arch}}/kaifa"

cd dist
for file in *; do
  if [[ -d "${file}" ]]; then
    tar zcvf "${file}".tar.gz "${file}"
    rm -rf "${file}"
  fi
done

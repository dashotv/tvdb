#!/usr/bin/env bash

VERSION=$(git describe --tags --always --dirty --match=v* 2>/dev/null || cat "${PWD}/.version" 2>/dev/null || echo v0)

mkdir -p openapi
printf "go:\n  package: github.com/dashotv/tvdb/openapi\n  version: %s" "$VERSION" >openapi/gen.yml
speakeasy generate sdk -l go -o openapi -s ./openapi.yml
rm -rf openapi/go.*
mv openapi/pkg/* openapi/
rm -rf openapi/pkg

find ./openapi -type f -print0 |
  while IFS= read -r -d '' line; do
    echo "$line"
  done

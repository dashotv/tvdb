#!/usr/bin/env bash

VERSION=$(git describe --tags --always --dirty --match=v* 2>/dev/null || cat "${PWD}/.version" 2>/dev/null || echo v0)

mkdir -p openapi
printf "go:\n  packageName: github.com/dashotv/tvdb/openapi\n  version: %s" "$VERSION" >openapi/gen.yaml
speakeasy generate sdk -l go -o openapi -s ./openapi.yml
rm -rf openapi/go.*
mv openapi/pkg/* openapi/
rm -rf openapi/pkg

find ./openapi -type f -exec sed -i '.backup' "s/pkg\///g" {} \;
find ./openapi -type f -name '*.backup' -delete

echo "package tvdb" >operations.go
echo "" >>operations.go
grep -E '^func \(\w+ \*\w+\) ' openapi/*.go | awk -F':' '{print $2}' >>operations.go

git add .

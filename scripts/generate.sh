#!/usr/bin/env bash
if [[ "$#" -ne 2 ]]; then
  echo "Usage: $0 <name> <spec>"
  exit 1
fi

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
VERSION=$(cat "${PWD}/.version" 2>/dev/null || echo v0)
NAME="$1"
SPEC="$2"

rm -rf openapi
mkdir -p openapi
# set golang defaults for generator
printf "go:\n  packageName: github.com/dashotv/$NAME/openapi\n  version: %s" "$VERSION" >openapi/gen.yaml
# generate go sdk
speakeasy generate sdk -l go -o openapi -s "./$SPEC"
# cleanup generated mod files
rm -rf openapi/go.*
# remove pkg folder
mv openapi/pkg/* openapi/
rm -rf openapi/pkg
# remove pkg from imports
find ./openapi -type f -exec sed -i '.backup' "s/pkg\///g" {} \;
find ./openapi -type f -name '*.backup' -delete

# copy types to root
{
  echo "package $NAME"
  echo
  echo 'import "github.com/dashotv/'"$NAME"'/openapi/models/operations"'
  echo
  echo "// pointers"
  grep '^func ' openapi/types/pointers.go
  echo
  echo "// request, response"
  grep '^type ' openapi/models/operations/*.go | awk -F':' '{print $2}' | awk '{print $2}' | while read -r l; do
    if [[ $l == *200ApplicationJSON* ]]; then
      echo "type ${l/200ApplicationJSON/Response} = operations.${l}"
    elif [[ $l == *Response* ]]; then
      echo "type ${l/Response/FullResponse} = operations.${l}"
    else
      echo "type $l = operations.${l}"
    fi
  done
} >types.go

# copy functions to root
{
  echo "package $NAME"
  echo
  echo 'import ('
  echo '	"github.com/pkg/errors"'
  echo
  echo '	"github.com/dashotv/'"$NAME"'/openapi/models/operations"'
  echo ')'
  echo
  grep -E '^func \(' openapi/*.go | awk -F':' '{print $2}' | "$SCRIPT_DIR/rewrite_funcs.rb"
} >functions.go

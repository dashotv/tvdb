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
cat <<HERE >openapi/gen.yaml
configVersion: 1.0.0
generation:
  comments: {}
  sdkClassName: SDK
  usageSnippets:
    optionalPropertyRendering: withExample
features:
  go:
    constsAndDefaults: 0.1.1
    core: 3.1.5
    flattening: 2.81.1
    globalSecurity: 2.82.2
    globalServerURLs: 2.82.0
go:
  version: $VERSION
  clientServerStatusCodesAsErrors: true
  flattenGlobalSecurity: false
  imports:
    option: openapi
    paths:
      callbacks: models/callbacks
      errors: models/sdkerrors
      operations: models/operations
      shared: models/shared
      webhooks: models/webhooks
  inputModelSuffix: input
  maxMethodParams: 4
  outputModelSuffix: output
  packageName: github.com/dashotv/$NAME/openapi
HERE
# generate go sdk
speakeasy generate sdk -l go -o openapi -s "./$SPEC"
# cleanup generated mod files
rm -rf openapi/go.*

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
    if [[ $l == *ResponseBody* ]]; then
      echo "type ${l/ResponseBody/Response} = operations.${l}"
    elif [[ $l == *Response* ]]; then
      echo "type ${l/Response/Raw} = operations.${l}"
    else
      echo "type $l = operations.${l}"
    fi
  done
} >types.go

# copy functions to root
cat <<HERE >functions.go
package ${NAME}

import (
	"github.com/pkg/errors"

	"github.com/dashotv/${NAME}/openapi/models/operations"
)

HERE
cat <<HERE >functions_test.go
package ${NAME}

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dashotv/${NAME}/openapi/models/operations"
)

HERE

grep --with-filename -E '^func \(' openapi/*.go | awk -F':' '{print $2}' | "$SCRIPT_DIR/rewrite_funcs.rb"

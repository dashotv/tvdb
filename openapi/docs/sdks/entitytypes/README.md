# EntityTypes
(*EntityTypes*)

### Available Operations

* [GetEntityTypes](#getentitytypes) - returns the active entity types

## GetEntityTypes

returns the active entity types

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/dashotv/tvdb/openapi"
	"github.com/dashotv/tvdb/openapi/models/shared"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.EntityTypes.GetEntityTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEntityTypes200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetEntityTypesResponse](../../models/operations/getentitytypesresponse.md), error**


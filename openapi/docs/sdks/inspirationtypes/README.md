# InspirationTypes
(*InspirationTypes*)

### Available Operations

* [GetAllInspirationTypes](#getallinspirationtypes) - returns list of inspiration types records

## GetAllInspirationTypes

returns list of inspiration types records

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
    res, err := s.InspirationTypes.GetAllInspirationTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllInspirationTypes200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllInspirationTypesResponse](../../models/operations/getallinspirationtypesresponse.md), error**


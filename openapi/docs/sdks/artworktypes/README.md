# ArtworkTypes
(*ArtworkTypes*)

### Available Operations

* [GetAllArtworkTypes](#getallartworktypes) - Returns a list of artworkType records

## GetAllArtworkTypes

Returns a list of artworkType records

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
    res, err := s.ArtworkTypes.GetAllArtworkTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllArtworkTypes200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllArtworkTypesResponse](../../models/operations/getallartworktypesresponse.md), error**


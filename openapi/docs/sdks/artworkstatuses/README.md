# ArtworkStatuses
(*ArtworkStatuses*)

### Available Operations

* [GetAllArtworkStatuses](#getallartworkstatuses) - Returns list of artwork status records.

## GetAllArtworkStatuses

Returns list of artwork status records.

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
    res, err := s.ArtworkStatuses.GetAllArtworkStatuses(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllArtworkStatuses200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllArtworkStatusesResponse](../../models/operations/getallartworkstatusesresponse.md), error**


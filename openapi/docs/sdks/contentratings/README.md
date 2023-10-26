# ContentRatings
(*ContentRatings*)

### Available Operations

* [GetAllContentRatings](#getallcontentratings) - returns list content rating records

## GetAllContentRatings

returns list content rating records

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
    res, err := s.ContentRatings.GetAllContentRatings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllContentRatings200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllContentRatingsResponse](../../models/operations/getallcontentratingsresponse.md), error**


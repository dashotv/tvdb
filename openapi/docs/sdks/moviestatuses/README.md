# MovieStatuses
(*MovieStatuses*)

### Available Operations

* [GetAllMovieStatuses](#getallmoviestatuses) - returns list of status records

## GetAllMovieStatuses

returns list of status records

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
    res, err := s.MovieStatuses.GetAllMovieStatuses(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllMovieStatuses200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllMovieStatusesResponse](../../models/operations/getallmoviestatusesresponse.md), error**


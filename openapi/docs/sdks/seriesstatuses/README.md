# SeriesStatuses
(*SeriesStatuses*)

### Available Operations

* [GetAllSeriesStatuses](#getallseriesstatuses) - returns list of status records

## GetAllSeriesStatuses

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
    res, err := s.SeriesStatuses.GetAllSeriesStatuses(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllSeriesStatuses200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllSeriesStatusesResponse](../../models/operations/getallseriesstatusesresponse.md), error**

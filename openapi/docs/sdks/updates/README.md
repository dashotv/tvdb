# Updates
(*Updates*)

### Available Operations

* [Updates](#updates) - Returns updated entities.  methodInt indicates a created record (1), an updated record (2), or a deleted record (3).  If a record is deleted because it was a duplicate of another record, the target record's information is provided in mergeToType and mergeToId.

## Updates

Returns updated entities.  methodInt indicates a created record (1), an updated record (2), or a deleted record (3).  If a record is deleted because it was a duplicate of another record, the target record's information is provided in mergeToType and mergeToId.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/dashotv/tvdb/openapi"
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi/models/operations"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(""),
    )


    var since int64 = 75660

    var action *operations.UpdatesAction = operations.UpdatesActionDelete

    var page *int64 = 745624

    var type_ *operations.UpdatesType = operations.UpdatesTypeMovies

    ctx := context.Background()
    res, err := s.Updates.Updates(ctx, since, action, page, type_)
    if err != nil {
        log.Fatal(err)
    }

    if res.Updates200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |                                                                       |
| `since`                                                               | *int64*                                                               | :heavy_check_mark:                                                    | N/A                                                                   |                                                                       |
| `action`                                                              | [*operations.UpdatesAction](../../models/operations/updatesaction.md) | :heavy_minus_sign:                                                    | N/A                                                                   | movies                                                                |
| `page`                                                                | **int64*                                                              | :heavy_minus_sign:                                                    | name                                                                  |                                                                       |
| `type_`                                                               | [*operations.UpdatesType](../../models/operations/updatestype.md)     | :heavy_minus_sign:                                                    | N/A                                                                   | movies                                                                |


### Response

**[*operations.UpdatesResponse](../../models/operations/updatesresponse.md), error**


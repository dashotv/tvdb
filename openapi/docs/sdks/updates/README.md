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
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"github.com/dashotv/tvdb/openapi/models/operations"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
        }),
    )
    var since int64 = 75660

    var action *operations.Action = operations.ActionDelete.ToPointer()

    var page *float64 = openapi.Float64(7456.24)

    var type_ *operations.Type = operations.TypeMovies.ToPointer()
    ctx := context.Background()
    res, err := s.Updates.Updates(ctx, since, action, page, type_)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                               | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ctx`                                                   | [context.Context](https://pkg.go.dev/context#Context)   | :heavy_check_mark:                                      | The context to use for the request.                     |                                                         |
| `since`                                                 | *int64*                                                 | :heavy_check_mark:                                      | N/A                                                     |                                                         |
| `action`                                                | [*operations.Action](../../models/operations/action.md) | :heavy_minus_sign:                                      | N/A                                                     | movies                                                  |
| `page`                                                  | **float64*                                              | :heavy_minus_sign:                                      | name                                                    |                                                         |
| `type_`                                                 | [*operations.Type](../../models/operations/type.md)     | :heavy_minus_sign:                                      | N/A                                                     | movies                                                  |


### Response

**[*operations.UpdatesResponse](../../models/operations/updatesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

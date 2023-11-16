# SourceTypes
(*SourceTypes*)

### Available Operations

* [GetAllSourceTypes](#getallsourcetypes) - returns list of sourceType records

## GetAllSourceTypes

returns list of sourceType records

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SourceTypes.GetAllSourceTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllSourceTypesResponse](../../models/operations/getallsourcetypesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

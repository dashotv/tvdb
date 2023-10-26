# Genders
(*Genders*)

### Available Operations

* [GetAllGenders](#getallgenders) - returns list of gender records

## GetAllGenders

returns list of gender records

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
    res, err := s.Genders.GetAllGenders(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllGenders200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllGendersResponse](../../models/operations/getallgendersresponse.md), error**


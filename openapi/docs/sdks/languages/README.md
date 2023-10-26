# Languages
(*Languages*)

### Available Operations

* [GetAllLanguages](#getalllanguages) - returns list of language records

## GetAllLanguages

returns list of language records

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
    res, err := s.Languages.GetAllLanguages(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllLanguages200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllLanguagesResponse](../../models/operations/getalllanguagesresponse.md), error**


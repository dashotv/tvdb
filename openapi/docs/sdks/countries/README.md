# Countries
(*Countries*)

### Available Operations

* [GetAllCountries](#getallcountries) - returns list of country records

## GetAllCountries

returns list of country records

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
    res, err := s.Countries.GetAllCountries(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllCountries200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllCountriesResponse](../../models/operations/getallcountriesresponse.md), error**


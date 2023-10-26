# PeopleTypes
(*PeopleTypes*)

### Available Operations

* [GetAllPeopleTypes](#getallpeopletypes) - returns list of peopleType records

## GetAllPeopleTypes

returns list of peopleType records

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
    res, err := s.PeopleTypes.GetAllPeopleTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllPeopleTypes200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllPeopleTypesResponse](../../models/operations/getallpeopletypesresponse.md), error**


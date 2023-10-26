# AwardCategories
(*AwardCategories*)

### Available Operations

* [GetAwardCategory](#getawardcategory) - Returns a single award category base record
* [GetAwardCategoryExtended](#getawardcategoryextended) - Returns a single award category extended record

## GetAwardCategory

Returns a single award category base record

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


    var id float64 = 6788.97

    ctx := context.Background()
    res, err := s.AwardCategories.GetAwardCategory(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAwardCategory200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *float64*                                             | :heavy_check_mark:                                    | id                                                    |


### Response

**[*operations.GetAwardCategoryResponse](../../models/operations/getawardcategoryresponse.md), error**


## GetAwardCategoryExtended

Returns a single award category extended record

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


    var id float64 = 4271.53

    ctx := context.Background()
    res, err := s.AwardCategories.GetAwardCategoryExtended(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAwardCategoryExtended200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *float64*                                             | :heavy_check_mark:                                    | id                                                    |


### Response

**[*operations.GetAwardCategoryExtendedResponse](../../models/operations/getawardcategoryextendedresponse.md), error**


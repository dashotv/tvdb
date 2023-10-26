# Lists
(*Lists*)

### Available Operations

* [GetAllLists](#getalllists) - returns list of list base records
* [GetList](#getlist) - returns an list base record
* [GetListBySlug](#getlistbyslug) - returns an list base record search by slug
* [GetListExtended](#getlistextended) - returns a list extended record
* [GetListTranslation](#getlisttranslation) - Returns list translation record

## GetAllLists

returns list of list base records

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


    var page *float64 = 6231.5

    ctx := context.Background()
    res, err := s.Lists.GetAllLists(ctx, page)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllLists200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `page`                                                | **float64*                                            | :heavy_minus_sign:                                    | page number                                           |


### Response

**[*operations.GetAllListsResponse](../../models/operations/getalllistsresponse.md), error**


## GetList

returns an list base record

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


    var id float64 = 4515.12

    ctx := context.Background()
    res, err := s.Lists.GetList(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetList200ApplicationJSONObject != nil {
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

**[*operations.GetListResponse](../../models/operations/getlistresponse.md), error**


## GetListBySlug

returns an list base record search by slug

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


    var slug string = "string"

    ctx := context.Background()
    res, err := s.Lists.GetListBySlug(ctx, slug)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetListBySlug200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `slug`                                                | *string*                                              | :heavy_check_mark:                                    | slug                                                  |


### Response

**[*operations.GetListBySlugResponse](../../models/operations/getlistbyslugresponse.md), error**


## GetListExtended

returns a list extended record

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


    var id float64 = 1520.84

    ctx := context.Background()
    res, err := s.Lists.GetListExtended(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetListExtended200ApplicationJSONObject != nil {
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

**[*operations.GetListExtendedResponse](../../models/operations/getlistextendedresponse.md), error**


## GetListTranslation

Returns list translation record

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


    var id float64 = 3693.37

    var language string = "string"

    ctx := context.Background()
    res, err := s.Lists.GetListTranslation(ctx, id, language)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetListTranslation200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *float64*                                             | :heavy_check_mark:                                    | id                                                    |
| `language`                                            | *string*                                              | :heavy_check_mark:                                    | language                                              |


### Response

**[*operations.GetListTranslationResponse](../../models/operations/getlisttranslationresponse.md), error**


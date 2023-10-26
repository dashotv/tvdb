# People
(*People*)

### Available Operations

* [GetAllPeople](#getallpeople) - Returns a list of people base records with the basic attributes.
* [GetPeopleBase](#getpeoplebase) - Returns people base record
* [GetPeopleExtended](#getpeopleextended) - Returns people extended record
* [GetPeopleTranslation](#getpeopletranslation) - Returns people translation record

## GetAllPeople

Returns a list of people base records with the basic attributes.

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


    var page *float64 = 8592.53

    ctx := context.Background()
    res, err := s.People.GetAllPeople(ctx, page)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllPeople200ApplicationJSONObject != nil {
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

**[*operations.GetAllPeopleResponse](../../models/operations/getallpeopleresponse.md), error**


## GetPeopleBase

Returns people base record

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


    var id float64 = 7419.73

    ctx := context.Background()
    res, err := s.People.GetPeopleBase(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPeopleBase200ApplicationJSONObject != nil {
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

**[*operations.GetPeopleBaseResponse](../../models/operations/getpeoplebaseresponse.md), error**


## GetPeopleExtended

Returns people extended record

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


    var id float64 = 9525.83

    var meta *operations.GetPeopleExtendedMeta = operations.GetPeopleExtendedMetaTranslations

    ctx := context.Background()
    res, err := s.People.GetPeopleExtended(ctx, id, meta)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPeopleExtended200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |                                                                                       |
| `id`                                                                                  | *float64*                                                                             | :heavy_check_mark:                                                                    | id                                                                                    |                                                                                       |
| `meta`                                                                                | [*operations.GetPeopleExtendedMeta](../../models/operations/getpeopleextendedmeta.md) | :heavy_minus_sign:                                                                    | meta                                                                                  | translations                                                                          |


### Response

**[*operations.GetPeopleExtendedResponse](../../models/operations/getpeopleextendedresponse.md), error**


## GetPeopleTranslation

Returns people translation record

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


    var id float64 = 150.51

    var language string = "string"

    ctx := context.Background()
    res, err := s.People.GetPeopleTranslation(ctx, id, language)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPeopleTranslation200ApplicationJSONObject != nil {
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

**[*operations.GetPeopleTranslationResponse](../../models/operations/getpeopletranslationresponse.md), error**


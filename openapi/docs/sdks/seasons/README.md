# Seasons
(*Seasons*)

### Available Operations

* [GetAllSeasons](#getallseasons) - returns list of seasons base records
* [GetSeasonBase](#getseasonbase) - Returns season base record
* [GetSeasonExtended](#getseasonextended) - Returns season extended record
* [GetSeasonTranslation](#getseasontranslation) - Returns season translation record
* [GetSeasonTypes](#getseasontypes) - Returns season type records

## GetAllSeasons

returns list of seasons base records

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


    var page *float64 = 8175.94

    ctx := context.Background()
    res, err := s.Seasons.GetAllSeasons(ctx, page)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllSeasons200ApplicationJSONObject != nil {
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

**[*operations.GetAllSeasonsResponse](../../models/operations/getallseasonsresponse.md), error**


## GetSeasonBase

Returns season base record

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


    var id float64 = 9987.23

    ctx := context.Background()
    res, err := s.Seasons.GetSeasonBase(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeasonBase200ApplicationJSONObject != nil {
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

**[*operations.GetSeasonBaseResponse](../../models/operations/getseasonbaseresponse.md), error**


## GetSeasonExtended

Returns season extended record

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


    var id float64 = 8663.11

    ctx := context.Background()
    res, err := s.Seasons.GetSeasonExtended(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeasonExtended200ApplicationJSONObject != nil {
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

**[*operations.GetSeasonExtendedResponse](../../models/operations/getseasonextendedresponse.md), error**


## GetSeasonTranslation

Returns season translation record

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


    var id float64 = 5660.49

    var language string = "string"

    ctx := context.Background()
    res, err := s.Seasons.GetSeasonTranslation(ctx, id, language)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeasonTranslation200ApplicationJSONObject != nil {
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

**[*operations.GetSeasonTranslationResponse](../../models/operations/getseasontranslationresponse.md), error**


## GetSeasonTypes

Returns season type records

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
    res, err := s.Seasons.GetSeasonTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeasonTypes200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetSeasonTypesResponse](../../models/operations/getseasontypesresponse.md), error**


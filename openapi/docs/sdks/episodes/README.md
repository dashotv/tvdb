# Episodes
(*Episodes*)

### Available Operations

* [GetAllEpisodes](#getallepisodes) - Returns a list of episodes base records with the basic attributes.<br> Note that all episodes are returned, even those that may not be included in a series' default season order.
* [GetEpisodeBase](#getepisodebase) - Returns episode base record
* [GetEpisodeExtended](#getepisodeextended) - Returns episode extended record
* [GetEpisodeTranslation](#getepisodetranslation) - Returns episode translation record

## GetAllEpisodes

Returns a list of episodes base records with the basic attributes.<br> Note that all episodes are returned, even those that may not be included in a series' default season order.

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


    var page *int64 = 672768

    ctx := context.Background()
    res, err := s.Episodes.GetAllEpisodes(ctx, page)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllEpisodes200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `page`                                                | **int64*                                              | :heavy_minus_sign:                                    | page number                                           |


### Response

**[*operations.GetAllEpisodesResponse](../../models/operations/getallepisodesresponse.md), error**


## GetEpisodeBase

Returns episode base record

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


    var id int64 = 711597

    ctx := context.Background()
    res, err := s.Episodes.GetEpisodeBase(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEpisodeBase200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *int64*                                               | :heavy_check_mark:                                    | id                                                    |


### Response

**[*operations.GetEpisodeBaseResponse](../../models/operations/getepisodebaseresponse.md), error**


## GetEpisodeExtended

Returns episode extended record

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


    var id int64 = 20932

    var meta *operations.GetEpisodeExtendedMeta = operations.GetEpisodeExtendedMetaTranslations

    ctx := context.Background()
    res, err := s.Episodes.GetEpisodeExtended(ctx, id, meta)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEpisodeExtended200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `id`                                                                                    | *int64*                                                                                 | :heavy_check_mark:                                                                      | id                                                                                      |                                                                                         |
| `meta`                                                                                  | [*operations.GetEpisodeExtendedMeta](../../models/operations/getepisodeextendedmeta.md) | :heavy_minus_sign:                                                                      | meta                                                                                    | translations                                                                            |


### Response

**[*operations.GetEpisodeExtendedResponse](../../models/operations/getepisodeextendedresponse.md), error**


## GetEpisodeTranslation

Returns episode translation record

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


    var id int64 = 388621

    var language string = "string"

    ctx := context.Background()
    res, err := s.Episodes.GetEpisodeTranslation(ctx, id, language)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetEpisodeTranslation200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *int64*                                               | :heavy_check_mark:                                    | id                                                    |
| `language`                                            | *string*                                              | :heavy_check_mark:                                    | language                                              |


### Response

**[*operations.GetEpisodeTranslationResponse](../../models/operations/getepisodetranslationresponse.md), error**


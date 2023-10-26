# Movies
(*Movies*)

### Available Operations

* [GetAllMovie](#getallmovie) - returns list of movie base records
* [GetMovieBase](#getmoviebase) - Returns movie base record
* [GetMovieBaseBySlug](#getmoviebasebyslug) - Returns movie base record search by slug
* [GetMovieExtended](#getmovieextended) - Returns movie extended record
* [GetMovieTranslation](#getmovietranslation) - Returns movie translation record
* [GetMoviesFilter](#getmoviesfilter) - Search movies based on filter parameters

## GetAllMovie

returns list of movie base records

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


    var page *float64 = 9632.57

    ctx := context.Background()
    res, err := s.Movies.GetAllMovie(ctx, page)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllMovie200ApplicationJSONObject != nil {
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

**[*operations.GetAllMovieResponse](../../models/operations/getallmovieresponse.md), error**


## GetMovieBase

Returns movie base record

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


    var id float64 = 3592.32

    ctx := context.Background()
    res, err := s.Movies.GetMovieBase(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMovieBase200ApplicationJSONObject != nil {
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

**[*operations.GetMovieBaseResponse](../../models/operations/getmoviebaseresponse.md), error**


## GetMovieBaseBySlug

Returns movie base record search by slug

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
    res, err := s.Movies.GetMovieBaseBySlug(ctx, slug)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMovieBaseBySlug200ApplicationJSONObject != nil {
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

**[*operations.GetMovieBaseBySlugResponse](../../models/operations/getmoviebasebyslugresponse.md), error**


## GetMovieExtended

Returns movie extended record

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


    var id float64 = 4597.44

    var meta *operations.GetMovieExtendedMeta = operations.GetMovieExtendedMetaTranslations

    var short *bool = false

    ctx := context.Background()
    res, err := s.Movies.GetMovieExtended(ctx, id, meta, short)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMovieExtended200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `id`                                                                                                       | *float64*                                                                                                  | :heavy_check_mark:                                                                                         | id                                                                                                         |                                                                                                            |
| `meta`                                                                                                     | [*operations.GetMovieExtendedMeta](../../models/operations/getmovieextendedmeta.md)                        | :heavy_minus_sign:                                                                                         | meta                                                                                                       | translations                                                                                               |
| `short`                                                                                                    | **bool*                                                                                                    | :heavy_minus_sign:                                                                                         | reduce the payload and returns the short version of this record without characters, artworks and trailers. |                                                                                                            |


### Response

**[*operations.GetMovieExtendedResponse](../../models/operations/getmovieextendedresponse.md), error**


## GetMovieTranslation

Returns movie translation record

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


    var id float64 = 3544.29

    var language string = "string"

    ctx := context.Background()
    res, err := s.Movies.GetMovieTranslation(ctx, id, language)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMovieTranslation200ApplicationJSONObject != nil {
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

**[*operations.GetMovieTranslationResponse](../../models/operations/getmovietranslationresponse.md), error**


## GetMoviesFilter

Search movies based on filter parameters

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

    ctx := context.Background()
    res, err := s.Movies.GetMoviesFilter(ctx, operations.GetMoviesFilterRequest{
        Company: openapi.Float64(1),
        ContentRating: openapi.Float64(245),
        Country: "usa",
        Genre: openapi.Float64(3),
        Lang: "eng",
        Year: openapi.Float64(2020),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetMoviesFilter200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetMoviesFilterRequest](../../models/operations/getmoviesfilterrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetMoviesFilterResponse](../../models/operations/getmoviesfilterresponse.md), error**


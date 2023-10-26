# Series
(*Series*)

### Available Operations

* [GetAllSeries](#getallseries) - returns list of series base records
* [GetSeriesArtworks](#getseriesartworks) - Returns series artworks base on language and type. <br> Note&#58; Artwork type is an id that can be found using **/artwork/types** endpoint.
* [GetSeriesBase](#getseriesbase) - Returns series base record
* [GetSeriesBaseBySlug](#getseriesbasebyslug) - Returns series base record searched by slug
* [GetSeriesEpisodes](#getseriesepisodes) - Returns series episodes from the specified season type, default returns the episodes in the series default season type
* [GetSeriesExtended](#getseriesextended) - Returns series extended record
* [GetSeriesFilter](#getseriesfilter) - Search series based on filter parameters
* [GetSeriesNextAired](#getseriesnextaired) - Returns series base record including the nextAired field. <br> Note&#58; nextAired was included in the base record endpoint but that field will deprecated in the future so developers should use the nextAired endpoint.
* [GetSeriesSeasonEpisodesTranslated](#getseriesseasonepisodestranslated) - Returns series base record with episodes from the specified season type and language. Default returns the episodes in the series default season type.
* [GetSeriesTranslation](#getseriestranslation) - Returns series translation record

## GetAllSeries

returns list of series base records

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


    var page *float64 = 2844.5

    ctx := context.Background()
    res, err := s.Series.GetAllSeries(ctx, page)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllSeries200ApplicationJSONObject != nil {
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

**[*operations.GetAllSeriesResponse](../../models/operations/getallseriesresponse.md), error**


## GetSeriesArtworks

Returns series artworks base on language and type. <br> Note&#58; Artwork type is an id that can be found using **/artwork/types** endpoint.

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


    var id float64 = 681.34

    var lang *string = "eng, spa"

    var type_ *int64 = 1,2,3

    ctx := context.Background()
    res, err := s.Series.GetSeriesArtworks(ctx, id, lang, type_)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeriesArtworks200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `id`                                                  | *float64*                                             | :heavy_check_mark:                                    | id                                                    |                                                       |
| `lang`                                                | **string*                                             | :heavy_minus_sign:                                    | lang                                                  | eng, spa                                              |
| `type_`                                               | **int64*                                              | :heavy_minus_sign:                                    | type                                                  | 1,2,3                                                 |


### Response

**[*operations.GetSeriesArtworksResponse](../../models/operations/getseriesartworksresponse.md), error**


## GetSeriesBase

Returns series base record

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


    var id float64 = 7004.04

    ctx := context.Background()
    res, err := s.Series.GetSeriesBase(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeriesBase200ApplicationJSONObject != nil {
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

**[*operations.GetSeriesBaseResponse](../../models/operations/getseriesbaseresponse.md), error**


## GetSeriesBaseBySlug

Returns series base record searched by slug

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
    res, err := s.Series.GetSeriesBaseBySlug(ctx, slug)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeriesBaseBySlug200ApplicationJSONObject != nil {
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

**[*operations.GetSeriesBaseBySlugResponse](../../models/operations/getseriesbasebyslugresponse.md), error**


## GetSeriesEpisodes

Returns series episodes from the specified season type, default returns the episodes in the series default season type

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
    res, err := s.Series.GetSeriesEpisodes(ctx, operations.GetSeriesEpisodesRequest{
        ID: 3113.63,
        Page: 737433,
        SeasonType: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeriesEpisodes200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetSeriesEpisodesRequest](../../models/operations/getseriesepisodesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetSeriesEpisodesResponse](../../models/operations/getseriesepisodesresponse.md), error**


## GetSeriesExtended

Returns series extended record

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


    var id float64 = 4780.9

    var meta *operations.GetSeriesExtendedMeta = operations.GetSeriesExtendedMetaTranslations

    var short *bool = false

    ctx := context.Background()
    res, err := s.Series.GetSeriesExtended(ctx, id, meta, short)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeriesExtended200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     | Example                                                                                         |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |                                                                                                 |
| `id`                                                                                            | *float64*                                                                                       | :heavy_check_mark:                                                                              | id                                                                                              |                                                                                                 |
| `meta`                                                                                          | [*operations.GetSeriesExtendedMeta](../../models/operations/getseriesextendedmeta.md)           | :heavy_minus_sign:                                                                              | meta                                                                                            | translations                                                                                    |
| `short`                                                                                         | **bool*                                                                                         | :heavy_minus_sign:                                                                              | reduce the payload and returns the short version of this record without characters and artworks |                                                                                                 |


### Response

**[*operations.GetSeriesExtendedResponse](../../models/operations/getseriesextendedresponse.md), error**


## GetSeriesFilter

Search series based on filter parameters

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
    res, err := s.Series.GetSeriesFilter(ctx, operations.GetSeriesFilterRequest{
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

    if res.GetSeriesFilter200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetSeriesFilterRequest](../../models/operations/getseriesfilterrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetSeriesFilterResponse](../../models/operations/getseriesfilterresponse.md), error**


## GetSeriesNextAired

Returns series base record including the nextAired field. <br> Note&#58; nextAired was included in the base record endpoint but that field will deprecated in the future so developers should use the nextAired endpoint.

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


    var id float64 = 7399.64

    ctx := context.Background()
    res, err := s.Series.GetSeriesNextAired(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeriesNextAired200ApplicationJSONObject != nil {
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

**[*operations.GetSeriesNextAiredResponse](../../models/operations/getseriesnextairedresponse.md), error**


## GetSeriesSeasonEpisodesTranslated

Returns series base record with episodes from the specified season type and language. Default returns the episodes in the series default season type.

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


    var id float64 = 4718.81

    var lang string = "string"

    var page int64 = 498852

    var seasonType string = "string"

    ctx := context.Background()
    res, err := s.Series.GetSeriesSeasonEpisodesTranslated(ctx, id, lang, page, seasonType)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeriesSeasonEpisodesTranslated200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *float64*                                             | :heavy_check_mark:                                    | id                                                    |
| `lang`                                                | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `page`                                                | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `seasonType`                                          | *string*                                              | :heavy_check_mark:                                    | season-type                                           |


### Response

**[*operations.GetSeriesSeasonEpisodesTranslatedResponse](../../models/operations/getseriesseasonepisodestranslatedresponse.md), error**


## GetSeriesTranslation

Returns series translation record

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


    var id float64 = 8327.38

    var language string = "string"

    ctx := context.Background()
    res, err := s.Series.GetSeriesTranslation(ctx, id, language)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSeriesTranslation200ApplicationJSONObject != nil {
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

**[*operations.GetSeriesTranslationResponse](../../models/operations/getseriestranslationresponse.md), error**


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
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )


    var page *int64 = 284450

    ctx := context.Background()
    res, err := s.Series.GetAllSeries(ctx, page)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
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

**[*operations.GetAllSeriesResponse](../../models/operations/getallseriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSeriesArtworks

Returns series artworks base on language and type. <br> Note&#58; Artwork type is an id that can be found using **/artwork/types** endpoint.

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )


    var id int64 = 68134

    var lang *string = "eng, spa"

    var type_ *int64 = 1,2,3

    ctx := context.Background()
    res, err := s.Series.GetSeriesArtworks(ctx, id, lang, type_)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `id`                                                  | *int64*                                               | :heavy_check_mark:                                    | id                                                    |                                                       |
| `lang`                                                | **string*                                             | :heavy_minus_sign:                                    | lang                                                  | eng, spa                                              |
| `type_`                                               | **int64*                                              | :heavy_minus_sign:                                    | type                                                  | 1,2,3                                                 |


### Response

**[*operations.GetSeriesArtworksResponse](../../models/operations/getseriesartworksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSeriesBase

Returns series base record

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )


    var id int64 = 700404

    ctx := context.Background()
    res, err := s.Series.GetSeriesBase(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
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

**[*operations.GetSeriesBaseResponse](../../models/operations/getseriesbaseresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSeriesBaseBySlug

Returns series base record searched by slug

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )


    var slug string = "string"

    ctx := context.Background()
    res, err := s.Series.GetSeriesBaseBySlug(ctx, slug)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSeriesEpisodes

Returns series episodes from the specified season type, default returns the episodes in the series default season type

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"github.com/dashotv/tvdb/openapi/models/operations"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Series.GetSeriesEpisodes(ctx, operations.GetSeriesEpisodesRequest{
        ID: 311363,
        Page: 737433,
        SeasonType: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSeriesExtended

Returns series extended record

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"github.com/dashotv/tvdb/openapi/models/operations"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )


    var id int64 = 478090

    var meta *operations.GetSeriesExtendedQueryParamMeta = operations.GetSeriesExtendedQueryParamMetaTranslations

    var short *bool = false

    ctx := context.Background()
    res, err := s.Series.GetSeriesExtended(ctx, id, meta, short)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               | Example                                                                                                   |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |                                                                                                           |
| `id`                                                                                                      | *int64*                                                                                                   | :heavy_check_mark:                                                                                        | id                                                                                                        |                                                                                                           |
| `meta`                                                                                                    | [*operations.GetSeriesExtendedQueryParamMeta](../../models/operations/getseriesextendedqueryparammeta.md) | :heavy_minus_sign:                                                                                        | meta                                                                                                      | translations                                                                                              |
| `short`                                                                                                   | **bool*                                                                                                   | :heavy_minus_sign:                                                                                        | reduce the payload and returns the short version of this record without characters and artworks           |                                                                                                           |


### Response

**[*operations.GetSeriesExtendedResponse](../../models/operations/getseriesextendedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSeriesFilter

Search series based on filter parameters

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"github.com/dashotv/tvdb/openapi/models/operations"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Series.GetSeriesFilter(ctx, operations.GetSeriesFilterRequest{
        Company: openapi.Int64(1),
        ContentRating: openapi.Int64(245),
        Country: "usa",
        Genre: operations.QueryParamGenreThree.ToPointer(),
        Lang: "eng",
        Year: openapi.Int64(2020),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSeriesNextAired

Returns series base record including the nextAired field. <br> Note&#58; nextAired was included in the base record endpoint but that field will deprecated in the future so developers should use the nextAired endpoint.

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )


    var id int64 = 739964

    ctx := context.Background()
    res, err := s.Series.GetSeriesNextAired(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
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

**[*operations.GetSeriesNextAiredResponse](../../models/operations/getseriesnextairedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSeriesSeasonEpisodesTranslated

Returns series base record with episodes from the specified season type and language. Default returns the episodes in the series default season type.

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )


    var id int64 = 471881

    var lang string = "string"

    var page int64 = 498852

    var seasonType string = "string"

    ctx := context.Background()
    res, err := s.Series.GetSeriesSeasonEpisodesTranslated(ctx, id, lang, page, seasonType)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *int64*                                               | :heavy_check_mark:                                    | id                                                    |
| `lang`                                                | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `page`                                                | *int64*                                               | :heavy_check_mark:                                    | N/A                                                   |
| `seasonType`                                          | *string*                                              | :heavy_check_mark:                                    | season-type                                           |


### Response

**[*operations.GetSeriesSeasonEpisodesTranslatedResponse](../../models/operations/getseriesseasonepisodestranslatedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetSeriesTranslation

Returns series translation record

### Example Usage

```go
package main

import(
	"github.com/dashotv/tvdb/openapi/models/shared"
	"github.com/dashotv/tvdb/openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )


    var id int64 = 832738

    var language string = "string"

    ctx := context.Background()
    res, err := s.Series.GetSeriesTranslation(ctx, id, language)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
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

**[*operations.GetSeriesTranslationResponse](../../models/operations/getseriestranslationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

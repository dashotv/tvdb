# Genres
(*Genres*)

### Available Operations

* [GetAllGenres](#getallgenres) - returns list of genre records
* [GetGenreBase](#getgenrebase) - Returns genre record

## GetAllGenres

returns list of genre records

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
    res, err := s.Genres.GetAllGenres(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllGenres200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetAllGenresResponse](../../models/operations/getallgenresresponse.md), error**


## GetGenreBase

Returns genre record

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


    var id int64 = 723237

    ctx := context.Background()
    res, err := s.Genres.GetGenreBase(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetGenreBase200ApplicationJSONObject != nil {
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

**[*operations.GetGenreBaseResponse](../../models/operations/getgenrebaseresponse.md), error**


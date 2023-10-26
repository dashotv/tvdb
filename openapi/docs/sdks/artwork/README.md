# Artwork
(*Artwork*)

### Available Operations

* [GetArtworkBase](#getartworkbase) - Returns a single artwork base record.
* [GetArtworkExtended](#getartworkextended) - Returns a single artwork extended record.

## GetArtworkBase

Returns a single artwork base record.

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


    var id float64 = 6050.48

    ctx := context.Background()
    res, err := s.Artwork.GetArtworkBase(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetArtworkBase200ApplicationJSONObject != nil {
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

**[*operations.GetArtworkBaseResponse](../../models/operations/getartworkbaseresponse.md), error**


## GetArtworkExtended

Returns a single artwork extended record.

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


    var id float64 = 111.13

    ctx := context.Background()
    res, err := s.Artwork.GetArtworkExtended(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetArtworkExtended200ApplicationJSONObject != nil {
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

**[*operations.GetArtworkExtendedResponse](../../models/operations/getartworkextendedresponse.md), error**


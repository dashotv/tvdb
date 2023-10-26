# Favorites
(*Favorites*)

### Available Operations

* [CreateUserFavorites](#createuserfavorites) - creates a new user favorite
* [GetUserFavorites](#getuserfavorites) - returns user favorites

## CreateUserFavorites

creates a new user favorite

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
    res, err := s.Favorites.CreateUserFavorites(ctx, &shared.FavoriteRecord{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.FavoriteRecord](../../models/shared/favoriterecord.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.CreateUserFavoritesResponse](../../models/operations/createuserfavoritesresponse.md), error**


## GetUserFavorites

returns user favorites

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
    res, err := s.Favorites.GetUserFavorites(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUserFavorites200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetUserFavoritesResponse](../../models/operations/getuserfavoritesresponse.md), error**


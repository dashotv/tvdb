# Search
(*Search*)

### Available Operations

* [GetSearchResults](#getsearchresults) - Our search index includes series, movies, people, and companies. Search is limited to 5k results max.
* [GetSearchResultsByRemoteID](#getsearchresultsbyremoteid) - Search a series, movie, people, episode, company or season by specific remote id and returns a base record for that entity.

## GetSearchResults

Our search index includes series, movies, people, and companies. Search is limited to 5k results max.

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
    res, err := s.Search.GetSearchResults(ctx, operations.GetSearchResultsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSearchResults200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetSearchResultsRequest](../../models/operations/getsearchresultsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetSearchResultsResponse](../../models/operations/getsearchresultsresponse.md), error**


## GetSearchResultsByRemoteID

Search a series, movie, people, episode, company or season by specific remote id and returns a base record for that entity.

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


    var remoteID string = "string"

    ctx := context.Background()
    res, err := s.Search.GetSearchResultsByRemoteID(ctx, remoteID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSearchResultsByRemoteID200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `remoteID`                                                                              | *string*                                                                                | :heavy_check_mark:                                                                      | Search for a specific remote id.  Allows searching for an IMDB or EIDR id, for example. |


### Response

**[*operations.GetSearchResultsByRemoteIDResponse](../../models/operations/getsearchresultsbyremoteidresponse.md), error**

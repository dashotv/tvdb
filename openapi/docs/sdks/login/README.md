# Login
(*Login*)

### Available Operations

* [PostLogin](#postlogin) - create an auth token. The token has one month validation length.

## PostLogin

create an auth token. The token has one month validation length.

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
    res, err := s.Login.PostLogin(ctx, operations.PostLoginRequestBody{
        Apikey: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostLogin200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PostLoginRequestBody](../../models/operations/postloginrequestbody.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.PostLoginResponse](../../models/operations/postloginresponse.md), error**


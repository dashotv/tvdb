# UserInfo
(*UserInfo*)

### Available Operations

* [GetUserInfo](#getuserinfo) - returns user info
* [GetUserInfoByID](#getuserinfobyid) - returns user info by user id

## GetUserInfo

returns user info

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
    res, err := s.UserInfo.GetUserInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUserInfo200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetUserInfoResponse](../../models/operations/getuserinforesponse.md), error**


## GetUserInfoByID

returns user info by user id

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


    var id int64 = 339298

    ctx := context.Background()
    res, err := s.UserInfo.GetUserInfoByID(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUserInfoByID200ApplicationJSONObject != nil {
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

**[*operations.GetUserInfoByIDResponse](../../models/operations/getuserinfobyidresponse.md), error**


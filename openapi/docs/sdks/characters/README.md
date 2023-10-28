# Characters
(*Characters*)

### Available Operations

* [GetCharacterBase](#getcharacterbase) - Returns character base record

## GetCharacterBase

Returns character base record

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


    var id int64 = 195650

    ctx := context.Background()
    res, err := s.Characters.GetCharacterBase(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCharacterBase200ApplicationJSONObject != nil {
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

**[*operations.GetCharacterBaseResponse](../../models/operations/getcharacterbaseresponse.md), error**


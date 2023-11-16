# People
(*People*)

### Available Operations

* [GetAllPeople](#getallpeople) - Returns a list of people base records with the basic attributes.
* [GetPeopleBase](#getpeoplebase) - Returns people base record
* [GetPeopleExtended](#getpeopleextended) - Returns people extended record
* [GetPeopleTranslation](#getpeopletranslation) - Returns people translation record

## GetAllPeople

Returns a list of people base records with the basic attributes.

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


    var page *int64 = 859253

    ctx := context.Background()
    res, err := s.People.GetAllPeople(ctx, page)
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

**[*operations.GetAllPeopleResponse](../../models/operations/getallpeopleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetPeopleBase

Returns people base record

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


    var id int64 = 741973

    ctx := context.Background()
    res, err := s.People.GetPeopleBase(ctx, id)
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

**[*operations.GetPeopleBaseResponse](../../models/operations/getpeoplebaseresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetPeopleExtended

Returns people extended record

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


    var id int64 = 952583

    var meta *operations.GetPeopleExtendedQueryParamMeta = operations.GetPeopleExtendedQueryParamMetaTranslations

    ctx := context.Background()
    res, err := s.People.GetPeopleExtended(ctx, id, meta)
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
| `meta`                                                                                                    | [*operations.GetPeopleExtendedQueryParamMeta](../../models/operations/getpeopleextendedqueryparammeta.md) | :heavy_minus_sign:                                                                                        | meta                                                                                                      | translations                                                                                              |


### Response

**[*operations.GetPeopleExtendedResponse](../../models/operations/getpeopleextendedresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetPeopleTranslation

Returns people translation record

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


    var id int64 = 15051

    var language string = "string"

    ctx := context.Background()
    res, err := s.People.GetPeopleTranslation(ctx, id, language)
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

**[*operations.GetPeopleTranslationResponse](../../models/operations/getpeopletranslationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

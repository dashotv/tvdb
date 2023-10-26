# Companies
(*Companies*)

### Available Operations

* [GetAllCompanies](#getallcompanies) - returns a paginated list of company records
* [GetCompany](#getcompany) - returns a company record
* [GetCompanyTypes](#getcompanytypes) - returns all company type records

## GetAllCompanies

returns a paginated list of company records

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


    var page *float64 = 2609.1

    ctx := context.Background()
    res, err := s.Companies.GetAllCompanies(ctx, page)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAllCompanies200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `page`                                                | **float64*                                            | :heavy_minus_sign:                                    | name                                                  |


### Response

**[*operations.GetAllCompaniesResponse](../../models/operations/getallcompaniesresponse.md), error**


## GetCompany

returns a company record

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


    var id float64 = 310.35

    ctx := context.Background()
    res, err := s.Companies.GetCompany(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCompany200ApplicationJSONObject != nil {
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

**[*operations.GetCompanyResponse](../../models/operations/getcompanyresponse.md), error**


## GetCompanyTypes

returns all company type records

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
    res, err := s.Companies.GetCompanyTypes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCompanyTypes200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetCompanyTypesResponse](../../models/operations/getcompanytypesresponse.md), error**


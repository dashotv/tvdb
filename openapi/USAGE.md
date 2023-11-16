<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/dashotv/tvdb/openapi"
	"github.com/dashotv/tvdb/openapi/models/shared"
	"log"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity(shared.Security{
			BearerAuth: "",
		}),
	)

	ctx := context.Background()
	res, err := s.ArtworkStatuses.GetAllArtworkStatuses(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->
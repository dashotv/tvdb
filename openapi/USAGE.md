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
		openapi.WithSecurity(""),
	)

	var id int64 = 605048

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
<!-- End SDK Example Usage -->
# GetMovieExtendedRequest


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                       | *float64*                                                                                                  | :heavy_check_mark:                                                                                         | id                                                                                                         |                                                                                                            |
| `Meta`                                                                                                     | [*GetMovieExtendedMeta](../../models/operations/getmovieextendedmeta.md)                                   | :heavy_minus_sign:                                                                                         | meta                                                                                                       | translations                                                                                               |
| `Short`                                                                                                    | **bool*                                                                                                    | :heavy_minus_sign:                                                                                         | reduce the payload and returns the short version of this record without characters, artworks and trailers. |                                                                                                            |
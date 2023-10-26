# GetSeriesSeasonEpisodesTranslated200ApplicationJSONData


## Fields

| Field                                                                                                                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                                                                                                                           | Required                                                                                                                                                                                                                                                                                                                       | Description                                                                                                                                                                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Series`                                                                                                                                                                                                                                                                                                                       | [*shared.SeriesBaseRecord](../../models/shared/seriesbaserecord.md)                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                             | The base record for a series. All series airs time like firstAired, lastAired, nextAired, etc. are in US EST for US series, and for all non-US series, the time of the showâ€™s country capital or most populous city. For streaming services, is the official release time. See https://support.thetvdb.com/kb/faq.php?id=29. |
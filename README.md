# cloudflare-cache-warmup
This is a simple golang-app, that takes a csv-file (in a format described below) as stdin and warming-up 
[the CloudFlare CDN-caches](https://www.cloudflare.com/learning/cdn/what-is-caching/). 
This is achieved by sending http-get-requests for each url until 
`CF-Cache-Status = HIT` http-response-header is appeared.

CSV-file format:
```text
id,Title,Permalink,Status
9,"Some title",https://example.com/example-resource/,publish
```


## Running last release
```shell
docker pull iakunin/cloudflare-cache-warmup:latest \
&& cat ~/Downloads/Zapisi-Export-2020-December-14-1031.csv \
| docker run -i iakunin/cloudflare-cache-warmup:latest
```

## Building

Following command builds this project into executable binary file:
```shell
go build .
```

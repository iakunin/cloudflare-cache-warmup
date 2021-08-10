# wp-cloudflare-cache-warmup
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
LAST_RELEASE=$(curl -s https://api.github.com/repos/iakunin/wp-cloudflare-cache-warmup/releases/latest | grep -E 'browser_download_url' | cut -d '"' -f 4); \
cd $(mktemp -d) \
&& curl -sL $LAST_RELEASE --output application \
&& chmod u+x application \
&& cat ~/Downloads/Zapisi-Export-2020-December-14-1031.csv | ./application
```

## Building

Following command builds this project into executable binary file:
```shell
go build .
```

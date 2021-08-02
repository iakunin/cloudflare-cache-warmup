# wp-cloudflare-cache-warmup

## Running last release
```shell
LAST_RELEASE=$(curl -s https://api.github.com/repos/iakunin/wp-cloudflare-cache-warmup/releases/latest | grep -E 'browser_download_url' | cut -d '"' -f 4); \
cd $(mktemp -d) \
&& curl -sL $LAST_RELEASE --output wp-cloudflare-cache-warmup \
&& chmod u+x wp-cloudflare-cache-warmup \
&& ./wp-cloudflare-cache-warmup '/home/yakunin/Downloads/Zapisi-Export-2020-December-14-1031.csv'
```

## Building from source files
```shell
go build .
```

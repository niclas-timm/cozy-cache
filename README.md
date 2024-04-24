# Cozy Cache

A Go program to keep your cache warm and cozy by pinging urls from your sitemap.

## Getting started

Create a `cozy-cache.json` file with the following content:

```json
{
    "url": "https://your-domain.com",
    "priorities": [
        "/paths-that-should",
        "/be-cached",
        "/with-priority"
    ],
    "exclude": [
        "/paths-you-",
        "/do-not-want",
        "/to-cache",
    ],
    "limit" 1000
}
```

`url` is the domain you want to run the cache warmer on. Make sure it has a sitemap.xml located at `your-domain.com/sitemap.cml`, as cozy-cache will grab all urls from that file.

`priorities` is an array of paths (or any substring of a url) that you want to cache with priority. Paths that match the patters in `priorities` will be cached before any other paths.

`exclude` is the opposite of priorities. Any path matching this pattern will not be cached (pinged).

`limit` determines after how many pings cozy-cache will stop. Set to 0 or omit if you don't need a limit.


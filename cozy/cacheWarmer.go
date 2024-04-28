package cozy

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"niclas-timm/cozy-cache/config"
	"path/filepath"
)

func RunCacheWarmer(config config.Config) {
	sitemap := parseSitemap(config)
	walkSitemapUrls(config, sitemap)
}

func walkSitemapUrls(config config.Config, sitemap Sitemap) {
	var skipped []SitemapUrl
	pingedPriorityUrls := 0
	pingedNonPriorityUrls := 0
	totalPinged := 0

	// Ping priority urls and collect non-priorities in `skipped`.
	fmt.Println("🚀 Start pinging urls...")
	for _, url := range sitemap.Urls {
		if config.Limit != 0 && totalPinged >= config.Limit {
			fmt.Printf("Reached ping limit of %d. Aborting.\n", config.Limit)
			break
		}
		if !checkForSubstringInSlice(config.Priorities, url.Url) {
			skipped = append(skipped, url)
			continue
		}
		ping(url.Url)
		pingedPriorityUrls++
		totalPinged++
	}
	fmt.Printf("Pinged %d priority urls 👍\n", pingedPriorityUrls)

	if config.Limit != 0 && totalPinged >= config.Limit {
		return
	}

	// Ping non-priority urls.
	fmt.Printf("Continuing with non-priority urls...\n")
	for _, url := range skipped {
		if config.Limit != 0 && totalPinged >= config.Limit {
			fmt.Printf("Reached ping limit of %d. Aborting.\n", config.Limit)
			break
		}
		if checkForSubstringInSlice(config.Exclude, url.Url) {
			continue
		}
		ping(url.Url)
		pingedNonPriorityUrls++
		totalPinged++
	}

	fmt.Printf("Pinged %d non-priority urls 👍\n", pingedNonPriorityUrls)
}

func checkForSubstringInSlice(find []string, search string) bool {
	url, err := url.Parse(search)
	if err != nil {
		panic(err)
	}
	for _, s := range find {
		match, err := filepath.Match(s, url.Path)
		if err != nil {
			panic(err)
		}
		if match {
			return true
		}
	}
	return false
}

func ping(url string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Cozy Cache 1.0")
	client.Do(req)
}

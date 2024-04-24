package cozy

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"niclas-timm/cozy-cache/config"
	"os"
)

type Sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	Urls    []SitemapUrl `xml:"url"`
}

type SitemapUrl struct {
	XMLName xml.Name `xml:"url"`
	Url     string   `xml:"loc"`
}

func parseSitemap(config config.Config) Sitemap {
	res, err := http.Get(config.Url + "/sitemap.xml?page=2")
	if err != nil {
		fmt.Printf("Error while fetching sitemap.xml: %s", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Unable to parse sitemap response body")
	}

	var sitemap Sitemap

	xml.Unmarshal(body, &sitemap)
	return sitemap
}

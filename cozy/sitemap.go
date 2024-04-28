package cozy

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"niclas-timm/cozy-cache/config"
	"os"
	"strings"
)

type Sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	Urls    []SitemapUrl `xml:"url"`
}

type SitemapUrl struct {
	XMLName xml.Name `xml:"url"`
	Url     string   `xml:"loc"`
}

type SitemapIndex struct {
	XMLName xml.Name          `xml:"sitemapindex"`
	Urls    []SitemapIndexUrl `xml:"sitemap"`
}

type SitemapIndexUrl struct {
	XMLName xml.Name `xml:"sitemap"`
	Url     string   `xml:"loc"`
}

func parseSitemap(config config.Config) Sitemap {
	res, err := http.Get(config.Url + "/sitemap.xml")
	if err != nil {
		fmt.Printf("Error while fetching sitemap.xml: %s", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Unable to parse sitemap response body")
	}

	sitemaps := getSitemaps(body)
	return combineSitemaps(sitemaps)
}

func getSitemaps(rootSitemap []byte) []Sitemap {
	var sitemaps []Sitemap

	if hasMultipleSitemaps(rootSitemap) {
		fmt.Println("Detected multiple sitemaps 🕵️")
		var sitemapIndex SitemapIndex
		xml.Unmarshal(rootSitemap, &sitemapIndex)
		for _, s := range sitemapIndex.Urls {
			sitemaps = append(sitemaps, loadSitemap(s.Url))
		}
		return sitemaps
	}

	var sitemap Sitemap
	xml.Unmarshal(rootSitemap, &sitemap)
	sitemaps = append(sitemaps, sitemap)
	return sitemaps
}

func hasMultipleSitemaps(sitemap []byte) bool {
	return strings.Contains(string(sitemap), "sitemapindex")
}

func loadSitemap(url string) Sitemap {
	var sitemap Sitemap
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error while fetching sitemap.xml: %s", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("No body")
	}

	xml.Unmarshal(body, &sitemap)
	return sitemap
}

func combineSitemaps(sitemaps []Sitemap) Sitemap {
	var combined Sitemap
	for _, sitemap := range sitemaps {
		combined.Urls = append(combined.Urls, sitemap.Urls...)
	}
	return combined
}

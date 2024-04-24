package main

import (
	"niclas-timm/cozy-cache/config"
	"niclas-timm/cozy-cache/cozy"
)

func main() {
	config := config.ReadConfig()
	cozy.RunCacheWarmer(config)
}

package cozy

import "niclas-timm/cozy-cache/config"

func RunCacheWarmer(config config.Config) {
	parseSitemap(config)
}

package main

import (
	"fmt"
	"niclas-timm/cozy-cache/config"
)

func main() {
	config := config.ReadConfig()
	fmt.Println(config.Url)
}

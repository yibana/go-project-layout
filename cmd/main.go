package main

import "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/configs"

func main() {
	println(configs.PROXY)
	configs.LOGGER.Info("Hello")
	configs.LOGGER.Error("Hello")
}

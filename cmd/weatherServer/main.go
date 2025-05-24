package main

import "github.com/riccruzdev/zip-weather-ws-cloudrun/config"

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
}

package main

import "prj/internal/app"

const configPath = "configs/config.yml"

func main() {
	app.Run(configPath)
}

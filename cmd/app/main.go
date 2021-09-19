package main

import "prj/internal/app"

const configPath = "configs/config"

func main() {
	app.Run(configPath)
}

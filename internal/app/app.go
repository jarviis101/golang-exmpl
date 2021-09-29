package app

import (
	"fmt"
	"net/http"
	delivery "prj/internal/delivery/http"
)

func Run(configPath string) {
	cfg, err := NewConfig(configPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	r := delivery.New()
	http.Handle("/", r)
	fmt.Println("Server is listen...")
	if err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), nil); err != nil {
		return
	}
}

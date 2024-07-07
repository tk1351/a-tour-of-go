package controllers

import (
	"a-tour-of-go/config"
	"fmt"
	"net/http"
)

func StartMainServer() error {
	fmt.Println("Start Server 🚀")
	http.HandleFunc("/", index)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

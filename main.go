package main

import (
	"flag"
	"fmt"
	"net/http"
	"node-x/utils"
	"time"
)

var explorer *utils.Console
var err error
var RUNNING_AT = time.Now().Format("2006-01-02 15:04:05")

func main() {
	updater := ""
	flag.StringVar(&updater, "update_me_by_updater_this_is_a_tag_has_no_meaning", "", "updater path")
	flag.Parse()
	explorer, err = utils.NewBrowserContext()
	if err != nil {
		fmt.Println("Failed to create browser context:", err)
		return
	}
	LoadConfig()
	RUNNING_AT = time.Now().Format("2006-01-02 15:04:05")

	http.HandleFunc("/v1/web", webHandler)
	http.HandleFunc("/v1/raw", rawHandler)
	http.HandleFunc("/v1/info", infoHandler)
	http.HandleFunc("/v1/install", installHandler)
	http.HandleFunc("/v1/update", updateInstaller)
	http.HandleFunc("/v1/config", configupdateInstaller)
	http.HandleFunc("/v1/png", downloadImgHandler)
	time.Sleep(3 * time.Second)
	fmt.Println("Server started on :31111")
	http.ListenAndServe(":31111", nil)
}

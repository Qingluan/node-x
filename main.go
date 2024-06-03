package main

import (
	"flag"
	"fmt"
	"net/http"
	"node-x/utils"
	"time"

	"gitee.com/dark.H/gs"
	"gopkg.in/ini.v1"
)

var explorer *utils.Console
var err error
var configPath = gs.HOME.PathJoin(".config", "node-x.ini").ExpandUser().Str()
var ROLE = "master"
var MASTER = ""
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

	// parse config ini file, get default's role and master key
	cfg, err := ini.Load(configPath)
	if err == nil {
		// return
		ROLE = cfg.Section("").Key("role").String()
		MASTER = cfg.Section("").Key("master").String()

	}
	RUNNING_AT = time.Now().Format("2006-01-02 15:04:05")

	http.HandleFunc("/v1/web", webHandler)
	http.HandleFunc("/v1/raw", rawHandler)
	http.HandleFunc("/v1/info", infoHandler)
	http.HandleFunc("/v1/install", installHandler)
	http.HandleFunc("/v1/update", updateInstaller)
	http.HandleFunc("/v1/png", downloadImgHandler)
	time.Sleep(3 * time.Second)
	fmt.Println("Server started on :31111")
	http.ListenAndServe(":31111", nil)
}

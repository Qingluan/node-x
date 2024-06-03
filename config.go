package main

import (
	"gitee.com/dark.H/gs"
	"gopkg.in/ini.v1"
)

var configPath = gs.HOME.PathJoin(".config", "node-x.ini").ExpandUser().Str()
var ROLE = "master"
var TEXT_JS = ""
var MASTER = ""
var NODES = gs.Dict[string]{}

func LoadConfig() {
	// parse config ini file, get default's role and master key
	cfg, err := ini.Load(configPath)
	if err == nil {
		// return
		section := cfg.Section("")
		if section.Key("role") != nil {
			ROLE = section.Key("role").String()
		}

		if section.Key("master") != nil {

			MASTER = section.Key("master").String()
		}

		if section.Key("text_js") != nil {

			TEXT_JS = section.Key("text_js").String()
		}

		children := cfg.Section("nodes")
		if children != nil {

			for _, key := range children.Keys() {
				NODES[key.Name()] = gs.Str(key.String()).Trim().Str()
			}
		}

	}
}

package main

import (
	"node-x/asset"

	"gitee.com/dark.H/gs"
	"gopkg.in/ini.v1"
)

var CONFIG_PATH = gs.HOME.PathJoin(".config", "node-x.ini").ExpandUser().Str()
var ROLE = "master"
var TEXT_JS = ""
var MASTER = ""
var NODES = gs.Dict[string]{}

func LoadALlConfig() {
	// parse config ini file, get default's role and master key
	cfg, err := ini.Load(CONFIG_PATH)
	if err == nil {
		// return
		section := cfg.Section("default")
		if section.Key("role") != nil {
			ROLE = section.Key("role").String()
		}

		if section.Key("master") != nil {
			MASTER = section.Key("master").String()
		}

		if section.Key("text_js") != nil {
			TEXT_JS = gs.HOME.PathJoin(".config", "node-x-js", gs.Str(section.Key("text_js").String()).Basename().Str()).MustAsFile().Str()
		}

		children := cfg.Section("roles")
		if children != nil {
			for _, key := range children.Keys() {
				NODES[key.Name()] = gs.Str(key.String()).Trim().Str()
			}
		}
	}
}

func LoadSearchEngine(name string) (searchUrl string, javascript string) {
	if name == "google" {
		return "https://www.google.com/search?q=${KEY}&num=100", ""
	}
	cfg, err := ini.Load(CONFIG_PATH)
	if err == nil {
		// return
		section := cfg.Section("search_engine")
		if section != nil {
			if key := section.Key(name); key != nil {
				javascript = gs.HOME.PathJoin(".config", "node-x-js", name+".js").MustAsFile().String()
				return gs.Str(key.String()).Trim().Str(), javascript
			}
		}

	}
	return "", ""
}

func Release() {
	for _, name := range asset.AssetNames() {
		if gs.Str(name).EndsWith(".js") {
			f := gs.HOME.PathJoin(".config", "node-x-js", gs.Str(name).Basename().String())
			if !f.IsExists() {
				dir := gs.HOME.PathJoin(".config", "node-x-js").String()
				asset.RestoreAsset(dir, name)
			}
		} else if gs.Str(name).EndsWith(".ini") {
			f := gs.HOME.PathJoin(".config", "node-x.ini")
			if !f.IsExists() {
				asset.RestoreAsset(gs.HOME.PathJoin(".config").String(), name)
			}
		}

	}

}

func SetConfig(section string, key string, val string) (err error) {
	cfg, err := ini.Load(CONFIG_PATH)
	if err == nil {
		return err
	}
	if section == "" {
		section = "default"
	}
	if key == "" {
		return
	}

	if sec := cfg.Section(section); sec == nil {
		sec, err = cfg.NewSection(section)
		if err != nil {
			return err
		}
		if _key := sec.Key(key); _key == nil {
			sec.NewKey(key, val)
		} else {
			_key.SetValue(val)
		}

	} else {
		if _key := sec.Key(key); _key == nil {
			sec.NewKey(key, val)
		} else {
			_key.SetValue(val)
		}
	}
	LoadALlConfig()
	return nil
}

func LoadConfig(section string, key string) (val string) {
	cfg, err := ini.Load(CONFIG_PATH)
	if err == nil {
		return ""
	}
	if section == "" {
		section = "default"
	}
	if key == "" {
		return
	}

	if sec := cfg.Section(section); sec == nil {
		return ""
	} else {
		if _key := sec.Key(key); _key == nil {
			return ""
		} else {
			return _key.String()
		}

	}
}

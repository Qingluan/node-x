package utils

import (
	"fmt"

	"gitee.com/dark.H/gs"
)

type AccountModel struct {
	Name       string `json:"name" ini:"name"`
	Url        string `json:"url" ini:"url"`
	AccountCss string `json:"account" ini:"account"`
	Valid      string `json:"valid" ini:"valid"`
	Before     string `json:"before" ini:"before"`
	PhoneCode  string `json:"phone_code" ini:"phone_code"`
}

func (model *AccountModel) ActionChains() (as gs.List[gs.Str]) {
	// 登录前
	if model.Before != "" {
		for _, v := range gs.Str(model.Before).Split("|") {
			as = append(as, v.Trim())
		}
	}
	return
}

func ParseIni(content gs.Str) (d gs.List[*AccountModel], err error) {
	name := ""
	url := ""
	input := ""
	before := ""
	valid := ""
	PhoneCode := ""
	content.EveryLine(func(lineno int, line gs.Str) {
		if line.StartsWith("[") && line.Trim().EndsWith("]") {
			if name != "" && url != "" && input != "" && valid != "" {
				d = append(d, &AccountModel{
					Name:       name,
					Url:        url,
					AccountCss: input,
					Valid:      valid,
					Before:     before,
					PhoneCode:  PhoneCode,
				})
			}
			name = ""
			url = ""
			input = ""
			before = ""
			valid = ""
			PhoneCode = ""
			// fmt.Println(line)
			name = string(line[1:].Trim().Replace("]", ""))
			return
		}
		if line.In("=") {
			fs := line.Split("=", 2)
			key := fs[0].Trim()
			value := fs[1].Trim()
			switch key.String() {
			case "url":
				url = value.Trim().String()
			case "input":
				input = value.Trim().String()
			case "before":
				before = value.Trim().String()
			case "valid":
				valid = value.Trim().String()
			case "phone_code":
				PhoneCode = value.Trim().String()
			}
		}
		fmt.Println(name, url, input, before, valid)
	})
	if name != "" && url != "" {
		d = append(d, &AccountModel{
			Name:       name,
			Url:        url,
			AccountCss: input,
			Valid:      valid,
			Before:     before,
			PhoneCode:  PhoneCode,
		})
	}

	return d, nil

}

func ReadAccountModels(path string) (d gs.List[*AccountModel], err error) {
	return ParseIni(gs.Str(path).MustAsFile())
}

func SaveConfig(path string, ds gs.List[*AccountModel]) {
	lines := ""
	ds.Every(func(no int, i *AccountModel) {
		lines += fmt.Sprintf("[%s]\n", i.Name)
		lines += fmt.Sprintf("url = %s\n", i.Url)
		lines += fmt.Sprintf("input = %s\n", i.AccountCss)
		lines += fmt.Sprintf("before = %s\n", i.Before)
		lines += fmt.Sprintf("valid = %s\n", i.Valid)
		lines += fmt.Sprintf("phone_code = %s\n", i.PhoneCode)
		lines += "\n"
	})
	gs.Str(lines).ToFile(path, gs.O_NEW_WRITE)
}

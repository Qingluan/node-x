package utils

import (
	"encoding/json"
	"io"
)

type R struct {
	URL        string            `json:"url"`
	URLS       []string          `json:"urls"`
	Headers    map[string]string `json:"headers"`
	Proxy      string            `json:"proxy"`
	Script     string            `json:"script"`
	Screenshot bool              `json:"screenshot"`
}

func RFromJson(data []byte) (r *R, err error) {
	r = new(R)
	r.Headers = make(map[string]string)
	// 解析JSON数据
	err = json.Unmarshal(data, r)
	return
}

func RFromJsonReader(reader io.Reader) (r *R, err error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	r = new(R)
	r.Headers = make(map[string]string)
	// 解析JSON数据
	err = json.Unmarshal(data, r)
	return
}

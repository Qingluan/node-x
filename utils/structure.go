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
	LoadImage  bool              `json:"load_image"`
}

type Searcher struct {
	Query      string            `json:"script"`
	Name       string            `json:"name"`
	URL        string            `json:"url"`
	Headers    map[string]string `json:"headers"`
	Proxy      string            `json:"proxy"`
	Screenshot bool              `json:"screenshot"`
	LoadImage  bool              `json:"load_image"`
}

type ConnectRequest struct {
	URL string `json:"url"`
}

type ConnectResponse struct {
	ID   string `json:"ID"`
	Port int    `json:"port"`
}

type SearchItem struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Desc  string `json:"desc"`
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

func SearcherFromReader(reader io.Reader) (s *Searcher, err error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	s = new(Searcher)
	s.Headers = make(map[string]string)
	// 解析JSON数据
	err = json.Unmarshal(data, s)
	if s.Name == "" {
		s.Name = "google"
	}
	return
}

func SearcherFromJson(data []byte) (s *Searcher, err error) {
	s = new(Searcher)
	s.Headers = make(map[string]string)
	// 解析JSON数据
	err = json.Unmarshal(data, s)
	if s.Name == "" {
		s.Name = "google"
	}
	return
}

func SearchItemsFromJson(data []byte) (s []*SearchItem, err error) {
	err = json.Unmarshal(data, &s)
	return
}

func ConnectRequestFromReader(reader io.Reader) (r *ConnectRequest, err error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	r = new(ConnectRequest)
	// 解析JSON数据
	err = json.Unmarshal(data, r)
	return
}

func ConnectResponseFromJson(data []byte) (r *ConnectResponse, err error) {
	r = new(ConnectResponse)
	// 解析JSON数据
	err = json.Unmarshal(data, r)
	return
}

func ConnectResponseFromReader(reader io.Reader) (s *ConnectResponse, err error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	s = new(ConnectResponse)
	// 解析JSON数据
	err = json.Unmarshal(data, s)
	return
}

package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"gitee.com/dark.H/gs"
)

var (
	MY_ID  = gs.HOME.RandStr(32)
	UP_KEY = ""
)

func Check(url string) bool {
	// post req to url's /info . if "status","start" in json response
	// then it is a valid url
	// else it is not a valid url
	// return true or false

	// ignore ssl errors
	http.DefaultClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	resp, err := http.Post(url+"/v1/join", "application/json", strings.NewReader(gs.Dict[string]{
		"ID":   MY_ID.String(),
		"Port": fmt.Sprint(PORT),
	}.Json().String()))

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		data := gs.Dict[string]{}
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			return false
		}
		json.Unmarshal(buf, &data)

		if status, ok := data["status"]; ok && status == "alive" {
			if key, ok := data["key"]; ok {
				UP_KEY = key
			}
			return true
		}
	}
	return false
}

func RedirectToChildren(w http.ResponseWriter, requestData []byte, url string) {
	// reqbuf, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(gs.Dict[any]{
	// 		"err": err.Error(),
	// 	})
	// 	return
	// }
	http.DefaultClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	resp, err := http.Post(url, "application/json", strings.NewReader(string(requestData)))
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gs.Dict[any]{
			"err": err.Error(),
		})
		return
	}

	defer resp.Body.Close()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	reader := bufio.NewReader(resp.Body)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if bytes.HasSuffix(line, []byte("\n")) {
			w.Write(line)
		} else {
			io.WriteString(w, string(line)+"\n")
		}
		flusher.Flush()
	}
}

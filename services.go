package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"

	"node-x/utils"

	"gitee.com/dark.H/gs"
	"github.com/playwright-community/playwright-go"
	"golang.org/x/net/html"
)

var (
	NoJavascriptRegex = regexp.MustCompile(`<script[\w\W]+?</script>`)
	NocssRegex        = regexp.MustCompile(`<style[\w\W]+?</style>`)
	NoSVG             = regexp.MustCompile(`<svg[\w\W]+?</svg>`)
	NoIframe          = regexp.MustCompile(`<iframe[\w\W]+?</iframe>`)
	LinkRegex         = regexp.MustCompile(`<a[\w\W]+?</a>`)
	HrefRegex         = regexp.MustCompile(`href="([\w\W]+?)"`)
	LinkTextRegex     = regexp.MustCompile(`>([\w\W]+?)</a>`)
	TextRegex         = regexp.MustCompile(`>([^<][\w\W]+?)</`)
	MetaRegex         = regexp.MustCompile(`<meta[\w\W]+?>`)
)

type Link struct {
	Url     string `json:"url"`
	Content string `json:"content"`
	Text    string `json:"text"`
}

func jsUpdateInstaller(w http.ResponseWriter, r *http.Request) {
	// 接收一个上传的文件然后运行更新器
	// 并然后重启服务器
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	file, fh, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "解析文件出错:"+err.Error(), http.StatusBadRequest)
		return
	}

	// 读取文件内容

	gs.HOME.PathJoin(".config", "node-x-js").Mkdir()
	fname := gs.Str(fh.Filename)
	if fname.EndsWith(".js") {
		if gs.Str(fh.Filename).In("/") {
			fname = fname.Basename()
		}
		configPath := gs.HOME.PathJoin(".config", "node-x-js").PathJoin(fname.String())
		// 保存上传的文件
		_, tmpFile, err := configPath.OpenFile(gs.O_NEW_WRITE)
		// tmpFileStat, err := tmpFile.Stat()
		if err != nil {
			http.Error(w, "Failed to get temporary file information", http.StatusInternalServerError)
			file.Close()
			return
		}
		_, err = io.Copy(tmpFile, file)
		if err != nil {
			http.Error(w, "Failed to write to temporary file", http.StatusInternalServerError)
			return
		}
		file.Close()
		tmpFile.Close()

		// 运行更新器
		LoadALlConfig()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "upload :" + fname.String(),
			"end":    time.Now().Format("2006-01-02 15:04:05"),
		})
	} else {
		json.NewEncoder(w).Encode(map[string]string{
			"status": "upload failed",
			"end":    time.Now().Format("2006-01-02 15:04:05"),
		})
	}

}

func configupdateInstaller(w http.ResponseWriter, r *http.Request) {
	// 接收一个上传的文件然后运行更新器
	// 并然后重启服务器
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "解析文件出错:"+err.Error(), http.StatusBadRequest)
		return
	}

	// 读取文件内容

	gs.HOME.PathJoin(".config").Mkdir()
	configPath := gs.HOME.PathJoin(".config").PathJoin("node-x.ini")

	// 保存上传的文件
	_, tmpFile, err := configPath.OpenFile(gs.O_NEW_WRITE)
	// tmpFileStat, err := tmpFile.Stat()
	if err != nil {
		http.Error(w, "Failed to get temporary file information", http.StatusInternalServerError)
		file.Close()
		return
	}

	_, err = io.Copy(tmpFile, file)
	if err != nil {
		http.Error(w, "Failed to write to temporary file", http.StatusInternalServerError)
		return
	}
	file.Close()
	tmpFile.Close()

	// 运行更新器
	LoadALlConfig()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "config",
		"end":    time.Now().Format("2006-01-02 15:04:05"),
	})
}

func upgradeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pwdData := gs.Dict[string]{}

	err := json.NewDecoder(r.Body).Decode(&pwdData)
	if err != nil {
		http.Error(w, "解析文件出错:"+err.Error(), http.StatusBadRequest)
		return
	}
	pwd := pwdData["pwd"]
	version := pwdData["version"]
	if pwd == "H3ll0" {
		uu := "https://github.com/Qingluan/node-x/releases/download/" + version + "/node-x"
		updater := DownloadURL(uu)
		if updater == "" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status": "version not exists!",
				"end":    time.Now().Format("2006-01-02 15:04:05"),
			})
			return
		}
		stat, err := os.Stat(updater.String())
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status": "url not file :" + err.Error(),
				"end":    time.Now().Format("2006-01-02 15:04:05"),
			})
			return
		}

		if stat.Size() < 1024*1024 {
			updater.Rm()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status": "url not file : too small",
				"end":    time.Now().Format("2006-01-02 15:04:05"),
			})
		}
		executable, err := os.Executable()
		if err != nil {
			http.Error(w, "Failed to get current executable path", http.StatusInternalServerError)
			// tmpFile.Close()

			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "upgrading",
			"end":    time.Now().Format("2006-01-02 15:04:05"),
		})

		// 运行更新器
		utils.Uprade(updater.String(), executable)
		// 退出自己
		os.Exit(0)
	} else {
		http.Error(w, "password wrong!", http.StatusUnauthorized)
		return
	}

}

func updateInstaller(w http.ResponseWriter, r *http.Request) {
	// 接收一个上传的文件然后运行更新器
	// 并然后重启服务器
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	updater := gs.Str("")

	if err != nil {
		req, err2 := utils.RFromJsonReader(r.Body)
		if err2 != nil {
			http.Error(w, "解析文件出错:"+err.Error(), http.StatusBadRequest)
			return
		}
		if req.URL != "" && req.Headers != nil && req.Headers["pwd"] != "" && req.Headers["pwd"] == "H3ll0" {
			updater = DownloadURL(req.URL)
			if updater == "" {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"status": "url not file url",
					"end":    time.Now().Format("2006-01-02 15:04:05"),
				})
				return
			}
			stat, err := os.Stat(updater.String())
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"status": "url not file :" + err.Error(),
					"end":    time.Now().Format("2006-01-02 15:04:05"),
				})
				return
			}

			if stat.Size() < 1024*1024 {
				updater.Rm()
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"status": "url not file : too small",
					"end":    time.Now().Format("2006-01-02 15:04:05"),
				})
			}
			return
		}
	} else {
		if pwd := r.FormValue("pwd"); pwd != "H3ll0" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status": "upload failed wrong password failed",
				"end":    time.Now().Format("2006-01-02 15:04:05"),
			})
			return
		}
		updater = gs.TMP.PathJoin("updated-installer")
		if updater.IsExists() {
			updater.Rm()
		}

		// 保存上传的文件
		_, tmpFile, err := updater.OpenFile(gs.O_NEW_WRITE)
		// tmpFileStat, err := tmpFile.Stat()
		if err != nil {
			http.Error(w, "Failed to get temporary file information", http.StatusInternalServerError)
			file.Close()
			return
		}

		_, err = io.Copy(tmpFile, file)
		if err != nil {
			http.Error(w, "Failed to write to temporary file", http.StatusInternalServerError)
			return
		}
		file.Close()
		tmpFile.Close()
	}

	// 读取文件内容

	// 获取当前可执行文件的路径
	executable, err := os.Executable()
	if err != nil {
		http.Error(w, "Failed to get current executable path", http.StatusInternalServerError)
		file.Close()
		// tmpFile.Close()

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "upgrading",
		"end":    time.Now().Format("2006-01-02 15:04:05"),
	})

	// 运行更新器
	utils.Uprade(updater.String(), executable)
	// 退出自己
	os.Exit(0)
}

func DownloadURL(u string) gs.Str {
	http.DefaultClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	if res, err := http.Get(u); err == nil {
		t := gs.TMP.PathJoin("updated-installer")
		if t.IsExists() {
			t.Rm()
		}
		if f, err := os.OpenFile(t.String(), os.O_CREATE|os.O_WRONLY, os.ModePerm); err == nil {
			io.Copy(f, res.Body)
			f.Close()
			res.Body.Close()
			return t
		}
	}
	return ""
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	req, err := utils.RFromJsonReader(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	urls := gs.List[string](req.URLS)
	if !urls.In(req.URL) {
		urls = append(urls, req.URL)
	}
	if len(urls) > 0 {

		lock := sync.WaitGroup{}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}
		for no, u := range urls {
			lock.Add(1)
			noRole := 0
			if len(NODES) > 0 {
				noRole = no % (len(NODES) + 1)

			}
			if noRole == 0 {
				go func(url string) {
					defer lock.Done()
					explorer.OpenNoScreen(url, func(screenPath string, page playwright.Page, res playwright.Response) {
						// execute the script
						title, _ := page.Title()
						defer page.Close()

						if req.Script != "" {
							result, err := page.Evaluate(req.Script)
							if err != nil {
								json.NewEncoder(w).Encode(gs.Dict[any]{
									"url": url,
									"err": err.Error(),
								})
								flusher.Flush()
								return
							}

							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url":   url,
								"title": title,
								"page":  result,
							})

						} else {
							if req.Screenshot {

								json.NewEncoder(w).Encode(gs.Dict[any]{
									"url":        url,
									"title":      title,
									"screenshot": screenPath,
								})
								flusher.Flush() // 刷新数据到客户端

							} else {
								html, err := page.Content()
								if err != nil {
									return
								}
								json.NewEncoder(w).Encode(gs.Dict[any]{
									"url":   url,
									"title": title,
									"page":  html,
								})
								flusher.Flush() // 刷新数据到客户端
							}
						}

					})
				}(u)
			} else {
				go func(url string) {
					defer lock.Done()

					ip := NODES.Keys().Sort(func(l, r gs.Str) bool {
						if l[0] < r[0] {
							return true
						} else if l[0] == r[0] {
							if len(l) > 1 && len(r) > 1 {
								return l[1] < r[1]
							} else {
								return len(l) < len(r)
							}
						} else {
							return false
						}
					})[noRole-1]
					d := &utils.R{
						URL:     url,
						Headers: make(map[string]string),
					}
					for k, v := range req.Headers {
						d.Headers[k] = v
					}
					d.Proxy = req.Proxy
					d.Screenshot = req.Screenshot
					d.LoadImage = req.LoadImage
					d.Script = req.Script
					data, _ := json.Marshal(d)
					RedirectToChildren(w, data, ip.String()+"/v1/web")
				}(u)
			}

		}
		time.Sleep(1 * time.Second)
		lock.Wait()

	} else {
		explorer.OpenNoScreen(req.URL, func(screenPath string, page playwright.Page, res playwright.Response) {
			// execute the script
			w.Header().Set("Content-Type", "application/json")
			title, _ := page.Title()
			defer page.Close()
			if req.Script != "" {
				result, err := page.Evaluate(req.Script)
				if err != nil {
					json.NewEncoder(w).Encode(gs.Dict[any]{
						"url": req.URL,
						"err": err.Error(),
					})
					// flusher.Flush()
					return
				}

				json.NewEncoder(w).Encode(gs.Dict[any]{
					"url":   req.URL,
					"title": title,
					"page":  result,
				})

			} else {
				if req.Screenshot {

					json.NewEncoder(w).Encode(gs.Dict[any]{
						"url":        req.URL,
						"title":      title,
						"screenshot": screenPath,
					})

				} else {
					html, err := page.Content()
					if err != nil {
						http.Error(w, "Failed to load URL :"+req.URL, http.StatusInternalServerError)
						return
					}
					json.NewEncoder(w).Encode(gs.Dict[any]{
						"url":   req.URL,
						"title": title,
						"page":  html,
					})
				}
			}
		})
		if err, ok := explorer.PagesErrors[req.URL]; ok {
			if err != nil {
				http.Error(w, "Failed to load URL :"+req.URL, http.StatusInternalServerError)
				return
			}
		}

	}

}

func webNewsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	req, err := utils.RFromJsonReader(r.Body)
	if err != nil {

		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	http.DefaultClient = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flush, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	if len(req.URLS) == 0 && req.URL != "" {
		req.URLS = append(req.URLS, req.URL)
	}
	wait := sync.WaitGroup{}
	for _, u := range req.URLS {
		wait.Add(1)
		fmt.Println("spide:", u)
		time.Sleep(10 * time.Millisecond)
		go func(url string) {
			// get base url from url.
			// baseURL := strings.Join(strings.Split(url, "/")[:3], "/")
			defer wait.Done()
			if res, err := http.Get(u); err != nil {
				be := bytes.NewBuffer([]byte{})
				json.NewEncoder(be).Encode(gs.Dict[any]{
					"url":    u,
					"status": -1,
					"err":    err.Error(),
				})
				w.Write([]byte("data: " + be.String()))
				flush.Flush()
			} else {
				be := bytes.NewBuffer([]byte{})
				buf, err := io.ReadAll(res.Body)
				if err != nil {
					json.NewEncoder(be).Encode(gs.Dict[any]{
						"url":    u,
						"status": res.Status,
						"err":    err.Error(),
					})
					w.Write([]byte("data: " + be.String()))
					flush.Flush()
					return
				}
				res.Body.Close()
				metas := MetaRegex.FindAllString(string(buf), -1)
				tags, err := html.Parse(bytes.NewBuffer([]byte("<html><head>" + strings.Join(metas, "\n") + "</head><body></body><html>")))
				if err != nil {
					if err != nil {
						json.NewEncoder(be).Encode(gs.Dict[any]{
							"url":    u,
							"status": res.Status,
							"err":    err.Error(),
						})
						w.Write([]byte("data: " + be.String()))
						flush.Flush()
						return
					}
				}
				var f func(*html.Node)
				attrs := gs.Dict[gs.List[string]]{}
				f = func(n *html.Node) {
					if n.Type == html.ElementNode {
						if n.Data == "meta" {
							for _, attr := range n.Attr {
								if e, ok := attrs[attr.Key]; ok {
									if attr.Val != e[len(e)-1] {
										e = e.Add(attr.Val)
									}
									attrs[attr.Key] = e
								} else {
									attrs[attr.Key] = gs.List[string]{attr.Val}
								}
							}
						}
					}
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						f(c)
					}
				}
				f(tags)
				nojs := NoJavascriptRegex.ReplaceAllString(string(buf), "")
				noiframe := NoIframe.ReplaceAllString(nojs, "")
				nocss := NocssRegex.ReplaceAllString(noiframe, "")

				nosvg := NoSVG.ReplaceAllString(nocss, "")

				domDocTest := html.NewTokenizer(strings.NewReader(nosvg))
				previousStartTokenTest := domDocTest.Token()
				texts := []string{}
			loopDomTest:
				for {
					tt := domDocTest.Next()
					switch {
					case tt == html.ErrorToken:
						break loopDomTest // End of the document,  done
					case tt == html.StartTagToken:
						previousStartTokenTest = domDocTest.Token()
					case tt == html.TextToken:
						if previousStartTokenTest.Data == "script" {
							continue
						}
						TxtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
						if len(TxtContent) > 5 {
							texts = append(texts, TxtContent)
						}
					}
				}

				json.NewEncoder(be).Encode(gs.Dict[any]{
					"url":    url,
					"status": res.Status,
					// "body":    nosvg,
					"content": strings.Join(texts, "\n"),
					"meta":    attrs,
				})
				w.Write([]byte("data: " + be.String()))
				flush.Flush()
				return
			}
		}(u)
	}
	wait.Wait()

}

func weblinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	req, err := utils.RFromJsonReader(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	http.DefaultClient = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flush, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	if len(req.URLS) == 0 && req.URL != "" {
		req.URLS = append(req.URLS, req.URL)
	}
	waitg := sync.WaitGroup{}
	for _, u := range req.URLS {
		waitg.Add(1)
		go func(url string) {
			// get base url from url.
			baseURL := strings.Join(strings.Split(url, "/")[:3], "/")

			defer waitg.Done()
			if res, err := http.Get(u); err != nil {
				be := bytes.NewBuffer([]byte{})
				json.NewEncoder(be).Encode(gs.Dict[any]{
					"url":    u,
					"status": -1,
					"err":    err.Error(),
				})
				w.Write([]byte("data: " + be.String()))
				flush.Flush()
			} else {
				be := bytes.NewBuffer([]byte{})
				buf, err := io.ReadAll(res.Body)
				if err != nil {
					json.NewEncoder(be).Encode(gs.Dict[any]{
						"url":    u,
						"status": res.Status,
						"err":    err.Error(),
					})
					w.Write([]byte("data: " + be.String()))
					flush.Flush()
					return
				}
				res.Body.Close()
				noiframe := NoIframe.ReplaceAllString(string(buf), "")
				nojs := NoJavascriptRegex.ReplaceAllString(noiframe, "")
				nocss := NocssRegex.ReplaceAllString(nojs, "")
				nosvg := NoSVG.ReplaceAllString(nocss, "")
				links := LinkRegex.FindAllString(nosvg, -1)
				ls := gs.List[Link]{}
				for _, llink := range links {
					href := ""
					title := ""
					fs := HrefRegex.FindStringSubmatch(llink)
					if len(fs) > 0 {
						href = fs[1]
					}
					fs = LinkTextRegex.FindStringSubmatch(llink)
					if len(fs) > 0 {
						title = fs[1]
					}

					title = strings.TrimSpace(title)
					if strings.HasPrefix(href, "/") {
						href = baseURL + href
					} else if !strings.HasPrefix(href, "http") {
						href = url + href
					} else if strings.HasPrefix(href, "://") {
						href = "http" + href
					}

					if !strings.HasPrefix(href, "http") {
						continue
					}
					if len(strings.TrimPrefix(href, baseURL)) < 6 {
						continue
					}
					if title == "" {
						continue
					}

					uri := strings.Join(strings.Split(href, "/")[3:], "/")
					if len(uri) < 3 {
						continue
					}
					fs = strings.Split(uri, ".")
					ex := fs[len(fs)-1]
					switch ex {
					case "jpg", "png", "mp4", "jpeg", "gif":
						continue
					}
					puretext := title
					if strings.Contains(title, ">") && strings.Contains(title, "<") {
						puretext = ""
						fss := TextRegex.FindAllStringSubmatch(title, -1)
						for _, fi := range fss {
							for _, ii := range fi {
								if !strings.Contains(ii, "<") && !strings.Contains(ii, ">") {
									puretext += "\t" + strings.TrimSpace(ii)
								}
							}
						}
					}
					ls = append(ls, Link{
						Url:     href,
						Content: title,
						Text:    strings.TrimSpace(puretext),
					})
				}
				json.NewEncoder(be).Encode(gs.Dict[any]{
					"url":    u,
					"status": res.Status,
					"body":   nosvg,
					"links":  ls,
				})
				w.Write([]byte("data: " + be.String()))
				flush.Flush()
			}
		}(u)

	}
	waitg.Wait()

}

func webTextHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	req, err := utils.RFromJsonReader(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	urls := gs.List[string](req.URLS)
	if !urls.In(req.URL) {
		urls = append(urls, req.URL)
	}
	if len(urls) > 0 {

		lock := sync.WaitGroup{}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}
		// replys := []gs.Dict[any]{}
		for no, u := range urls {
			lock.Add(1)
			noRole := 0
			if len(NODES) > 0 {
				noRole = no % (len(NODES) + 1)

			}
			if noRole == 0 {
				go func(url string) {
					defer lock.Done()

					explorer.OpenNoScreen(url, func(screenPath string, page playwright.Page, res playwright.Response) {
						// execute the script
						defer page.Close()
						title, _ := page.Title()
						if TEXT_JS != "" {
							result, err := page.Evaluate(TEXT_JS)
							if err != nil {
								json.NewEncoder(w).Encode(gs.Dict[any]{
									"url": url,
									"err": err.Error(),
								})
								flusher.Flush()
								return
							}

							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url":   url,
								"title": title,
								"page":  result,
							})
							flusher.Flush()

						} else {

							ts, err := page.Locator("p").AllTextContents()
							if err != nil {
								json.NewEncoder(w).Encode(gs.Dict[any]{
									"url": url,
									"err": err.Error(),
								})
								flusher.Flush()
								return
							}
							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url":   url,
								"title": title,
								"page":  strings.Join(ts, "\n"),
							})
							flusher.Flush()

						}

					})
				}(u)
			} else {

				go func(url string) {
					defer lock.Done()

					ip := NODES.Keys().Sort(func(l, r gs.Str) bool {
						if l[0] < r[0] {
							return true
						} else if l[0] == r[0] {
							if len(l) > 1 && len(r) > 1 {
								return l[1] < r[1]
							} else {
								return len(l) < len(r)
							}
						} else {
							return false
						}
					})[noRole-1]
					d := &utils.R{
						URL:     url,
						Headers: make(map[string]string),
					}
					for k, v := range req.Headers {
						d.Headers[k] = v
					}
					d.Proxy = req.Proxy
					d.Screenshot = req.Screenshot
					d.LoadImage = req.LoadImage
					d.Script = req.Script
					data, _ := json.Marshal(d)
					RedirectToChildren(w, data, ip.String()+"/v1/text")
				}(u)
			}

		}
		time.Sleep(1 * time.Second)
		lock.Wait()

	} else {
		explorer.OpenNoScreen(req.URL, func(screenPath string, page playwright.Page, res playwright.Response) {
			// execute the script
			w.Header().Set("Content-Type", "application/json")
			title, _ := page.Title()
			defer page.Close()
			if TEXT_JS != "" {
				result, err := page.Evaluate(TEXT_JS)
				if err != nil {
					http.Error(w, "Failed to load URL :"+req.URL, http.StatusInternalServerError)
					return
				}

				json.NewEncoder(w).Encode(gs.Dict[any]{
					"url":   req.URL,
					"title": title,
					"page":  result,
				})

			} else {

				ts, err := page.Locator("p").AllTextContents()
				if err != nil {
					json.NewEncoder(w).Encode(gs.Dict[any]{
						"url": req.URL,
						"err": err.Error(),
					})
					// flusher.Flush()
					return
				}
				json.NewEncoder(w).Encode(gs.Dict[any]{
					"url":   req.URL,
					"title": title,
					"page":  strings.Join(ts, "\n"),
				})

			}
		})
		if err, ok := explorer.PagesErrors[req.URL]; ok {
			if err != nil {
				http.Error(w, "Failed to load URL :"+req.URL, http.StatusInternalServerError)
				return
			}
		}

	}

}

func rawHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	html := r.FormValue("html")
	if html == "" {
		http.Error(w, "Missing HTML content", http.StatusBadRequest)
		return
	}

	// // data, err := explorer.EvaluateScript(html)
	// if err != nil {
	// 	http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(data)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	explorer.Browser.Contexts()

	json.NewEncoder(w).Encode(map[string]string{
		"status": "alive",
		"start":  RUNNING_AT,
	})
}

func installHandler(w http.ResponseWriter, r *http.Request) {
	packages := strings.Split(r.FormValue("packages"), ",")
	if len(packages) == 0 {
		http.Error(w, "Missing packages", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("sudo", append([]string{"apt-get", "install", "-y"}, packages...)...)
	err := cmd.Run()
	if err != nil {
		http.Error(w, "Failed to install packages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func downloadImgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("path")
	U := gs.Str(url)
	if U.Basename().EndsWith(".png") && !U.In(";") && !U.In("$") && !U.In("!") && !U.In("`") {
		if U.IsExists() {
			http.ServeFile(w, r, U.String())
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, "not found", http.StatusNotFound)
}

func searcherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	req, err := utils.SearcherFromReader(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	searchedUrl := gs.Str("")
	if req.URL == "" {
		searchKey, Scripts := LoadSearchEngine(req.Name)
		if !gs.Str(searchKey).In("${KEY}") {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(gs.Dict[any]{
				"name":  req.Name,
				"query": req.Query,
				"err":   "no ${KEY} in " + searchKey,
			})
			return
		}
		if req.Script == "" {
			req.Script = Scripts
		}
		searchedUrl = gs.Str(searchKey).Replace("${KEY}", req.Query)
	} else {
		if !gs.Str(req.URL).In("${KEY}") {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(gs.Dict[any]{
				"name":  req.Name,
				"query": req.Query,
				"err":   "no ${KEY} in " + req.URL,
			})
			return
		}
		searchedUrl = gs.Str(req.URL).Replace("${KEY}", req.Query)
	}

	if req.Script == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gs.Dict[any]{
			"name":  req.Name,
			"query": req.Query,
			"err":   "no script",
		})
		return
	}
	if searchedUrl == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gs.Dict[any]{
			"name":  req.Name,
			"query": req.Query,
			"err":   "no url",
		})
		return
	}

	explorer.OpenNoScreen(searchedUrl.String(), func(sscreenPath string, spage playwright.Page, searchres playwright.Response) {
		// execute the script
		// w.Header().Set("Content-Type", "application/json")
		// title, _ := page.Title()
		defer spage.Close()
		fmt.Println("script:", req.Script)
		result, err := spage.Evaluate(req.Script)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(gs.Dict[any]{
				"name":  req.Name,
				"query": req.Query,
				"err":   err.Error(),
			})
			return
		}
		// fmt.Println("result:", result)

		if items, err := utils.SearchItemsFromJson([]byte(result.(string))); err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(gs.Dict[any]{
				"name":  req.Name,
				"query": req.Query,
				"err":   err.Error(),
			})
			return
		} else {
			lock := sync.WaitGroup{}
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")
			flusher, ok := w.(http.Flusher)
			if !ok {
				http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
				return
			}
			for _, item := range items {
				// fmt.Println("item", item)
				lock.Add(1)
				go func(item *utils.SearchItem) {
					defer lock.Done()
					if item.Url != "" {
						explorer.OpenNoScreen(item.Url, func(screenPath string, page playwright.Page, res playwright.Response) {
							defer page.Close()
							title, _ := page.Title()
							itemResult, err := page.Evaluate(TEXT_JS)
							if err != nil {

								json.NewEncoder(w).Encode(gs.Dict[any]{
									"title":      title,
									"url":        item.Url,
									"err":        err.Error(),
									"screenshot": screenPath,
								})
								flusher.Flush()
								return
							}
							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url":        item.Url,
								"title":      title,
								"screenshot": screenPath,
								"page":       itemResult,
							})
							flusher.Flush()

						})
						if err, ok := explorer.PagesErrors[item.Url]; ok && err != nil {
							gs.Str(err.Error()).Color("r").Println(item.Url)
							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url": item.Url,
								"err": err.Error(),
							})
							flusher.Flush()
							return
						}
					}

				}(item)

			}
			time.Sleep(1 * time.Second)
			lock.Wait()

		}

	})
	if err, ok := explorer.PagesErrors[searchedUrl.String()]; ok {
		if err != nil {
			http.Error(w, "Failed to load URL :"+req.URL, http.StatusInternalServerError)
			return
		}
	}
}

func reciveconnect(w http.ResponseWriter, req *http.Request) {
	conRes, err := utils.ConnectResponseFromReader(req.Body)
	if err != nil {
		ReplyErr(err, w)
		return
	}
	if conRes.ID != "" {
		ip := req.RemoteAddr
		if r := gs.Str(req.RemoteAddr); r.In(":") {
			ip = r.Split(":")[0].String()
		}
		ip += ":" + fmt.Sprint(conRes.Port)
		SetConfig("role", ip, conRes.ID)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gs.Dict[any]{
			"id":     conRes.ID,
			"status": "alive",
			"key":    MY_ID,
		})
	} else {
		ReplyErr(errors.New("id not found"), w)
	}

}

func connectTO(w http.ResponseWriter, req *http.Request) {
	c, err := utils.ConnectRequestFromReader(req.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gs.Dict[any]{
			"err": err.Error(),
		})
		return
	}
	if Check(c.URL) {
		SetConfig("", "up", c.URL)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gs.Dict[any]{
			"Connect": c.URL,
		})
	}
}

func ReplyErr(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gs.Dict[any]{
			"err": err.Error(),
		})
		return
	}
}

func webBurpHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		// Handle POST request
		http.Error(w, "Method unsupported!", http.StatusInternalServerError)

		return
	} else {
		// Handle GET request

		w.Header().Set("Content-Type", "application/json")
		return
	}
}

func logHandler(w http.ResponseWriter, req *http.Request) {
	LOGFILE := "/tmp/node-x.log"
	logfile, err := os.Open(LOGFILE)
	if err != nil {
		http.Error(w, "Failed to open log file", http.StatusInternalServerError)
		return
	}
	defer logfile.Close()

	// event stream
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	linebuffer := bufio.NewReader(logfile)
	for {
		line, _, err := linebuffer.ReadLine()
		if err != nil {
			break
		}
		fmt.Fprintf(w, "data: %s\n", strings.TrimSpace(string(line)))
		flusher.Flush()
	}
}

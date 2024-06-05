package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"node-x/utils"

	"gitee.com/dark.H/gs"
	"github.com/playwright-community/playwright-go"
)

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

func updateInstaller(w http.ResponseWriter, r *http.Request) {
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
	if pwd := r.FormValue("pwd"); pwd != "H3ll0" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "upload failed wrong password failed",
			"end":    time.Now().Format("2006-01-02 15:04:05"),
		})
		return
	}

	// 读取文件内容

	updater := gs.TMP.PathJoin("updated-installer")
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

	// 获取当前可执行文件的路径
	executable, err := os.Executable()
	if err != nil {
		http.Error(w, "Failed to get current executable path", http.StatusInternalServerError)
		file.Close()
		tmpFile.Close()

		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "upgrading",
		"end":    time.Now().Format("2006-01-02 15:04:05"),
	})

	// 运行更新器
	utils.Uprade(updater.String(), executable)

	// 退出自己
	os.Exit(0)
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

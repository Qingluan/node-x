package main

import (
	"encoding/json"
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

var ()

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
	LoadConfig()

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

	// 运行更新器
	utils.Uprade(updater.String(), executable)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "upgrading",
		"end":    time.Now().Format("2006-01-02 15:04:05"),
	})
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
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Streaming", "true")
		// replys := []gs.Dict[any]{}
		for _, u := range urls {
			lock.Add(1)

			go func(url string) {
				defer lock.Done()
				explorer.Open(url, func(screenPath string, page playwright.Page, res playwright.Response) {
					// execute the script
					defer page.Close()

					if req.Script != "" {
						result, err := page.Evaluate(req.Script)
						if err != nil {
							return
						}

						json.NewEncoder(w).Encode(gs.Dict[any]{
							"url":  url,
							"page": result,
						})

					} else {
						if req.Screenshot {

							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url":        url,
								"screenshot": screenPath,
							})

						} else {
							html, err := page.Content()
							if err != nil {
								return
							}
							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url":  url,
								"page": html,
							})
						}
					}

				})
			}(u)

		}
		time.Sleep(1 * time.Second)
		lock.Wait()

	} else {
		explorer.Open(req.URL, func(screenPath string, page playwright.Page, res playwright.Response) {
			// execute the script
			w.Header().Set("Content-Type", "application/json")
			defer page.Close()
			if req.Script != "" {
				result, err := page.Evaluate(req.Script)
				if err != nil {
					http.Error(w, "Failed to load URL :"+req.URL, http.StatusInternalServerError)
					return
				}

				json.NewEncoder(w).Encode(gs.Dict[any]{
					"url":  req.URL,
					"page": result,
				})

			} else {
				if req.Screenshot {

					json.NewEncoder(w).Encode(gs.Dict[any]{
						"url":        req.URL,
						"screenshot": screenPath,
					})

				} else {
					html, err := page.Content()
					if err != nil {
						http.Error(w, "Failed to load URL :"+req.URL, http.StatusInternalServerError)
						return
					}
					json.NewEncoder(w).Encode(gs.Dict[any]{
						"url":  req.URL,
						"page": html,
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

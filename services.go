package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"

	"node-x/utils"

	"gitee.com/dark.H/gs"
	"github.com/PuerkitoBio/goquery"
	"github.com/playwright-community/playwright-go"
	"golang.org/x/net/html"
)

var (
	NoJavascriptRegex = regexp.MustCompile(`<script[\w\W]+?</script>`)
	NocssRegex        = regexp.MustCompile(`<style[\w\W]+?</style>`)
	NoSVG             = regexp.MustCompile(`<svg[\w\W]+?</svg>`)
	NoIframe          = regexp.MustCompile(`<iframe[\w\W]+?</iframe>`)
	LinkRegex         = regexp.MustCompile(`<a [\w\W]+?</a>`)
	HrefRegex         = regexp.MustCompile(`href="([\w\W]+?)"`)
	LinkTextRegex     = regexp.MustCompile(`>([\w\W]+?)</a>`)
	TextRegex         = regexp.MustCompile(`>([^<][\w\W]+?)</`)
	MetaRegex         = regexp.MustCompile(`<meta[\w\W]+?>`)
	// LiRegex           = regexp.MustCompile(`<li [\w\W]+?</li>`)
	ULRegex     = regexp.MustCompile(`<ul[\w\W]+?</ul>`)
	TimeRegex   = regexp.MustCompile(`<time[\w\W]+?</time>`)
	TitleRegex  = regexp.MustCompile(`<title[\w\W]+?</title>`)
	FooterRegex = regexp.MustCompile(`<footer[\w\W]+?</footer>`)
	FootRegex   = regexp.MustCompile(`<foot[\w\W]+?</foot>`)
	NAMES       = []string{
		"head",
		"script",
		"foot",
		"footer",
		"iframe",
		"link",
		"style",
		"svg",
		"ul",
		"a",
	}
)

func removeElement(n *html.Node, tag string) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			if strings.ToLower(c.Data) == strings.ToLower(tag) {
				n.RemoveChild(c)
			}

			return // Assuming only one element to remove
		}
		removeElement(c, tag)
	}
}
func extractMeta(raw string) (gs.Dict[gs.List[string]], error) {
	attrs := gs.Dict[gs.List[string]]{}
	metas := MetaRegex.FindAllString(raw, -1)
	tags, err := html.Parse(bytes.NewBuffer([]byte("<html><head>" + strings.Join(metas, "\n") + "</head><body></body><html>")))
	if err != nil {
		if err != nil {

			return attrs, err
		}
	}
	var f func(*html.Node)

	f = func(n *html.Node) {
		key := ""
		val := ""
		if n.Type == html.ElementNode {
			if n.Data == "meta" {
				for _, attr := range n.Attr {
					if attr.Key == "name" {
						key = strings.TrimSpace(attr.Val)
					} else if attr.Key == "property" {
						key = strings.TrimSpace(attr.Val)
					} else if attr.Key == "content" {
						val = strings.TrimSpace(attr.Val)
					}
					if key != "" && val != "" {
						if e, ok := attrs[key]; ok {
							if val != e[len(e)-1] {
								e = e.Add(val)
							}
							attrs[key] = e
						} else {
							attrs[key] = gs.List[string]{val}
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(tags)
	return attrs, nil
}

func filter_garbage_old(raw string) string {
	nocss := NocssRegex.ReplaceAllString(raw, "")
	reader := strings.NewReader(nocss)
	doc, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}
	removeElement(doc, "head")
	removeElement(doc, "style")
	removeElement(doc, "script")
	removeElement(doc, "link")
	removeElement(doc, "iframe")
	removeElement(doc, "footer")
	removeElement(doc, "foot")

	removeElement(doc, "svg")
	removeElement(doc, "ul")
	removeElement(doc, "a")

	var buf bytes.Buffer
	html.Render(&buf, doc)
	return buf.String()
}

func filter_garbage(raw string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
	if err != nil {
		log.Fatal(err)
	}
	// 使用 CSS 选择器选择要删除的节点，这里选择 class 为 "remove" 的 p 元素
	for _, name := range NAMES {
		doc.Find(name).Each(func(i int, s *goquery.Selection) {
			// 删除选中的节点
			s.Remove()
		})
	}
	// 输出修改后的 HTML 内容
	htmlOut, err := doc.Html()

	return htmlOut
}

type Link struct {
	Source  string `json:"source"`
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
							realRes := ""
							switch req.Output {
							case "md", "markdown":
								realRes = utils.HTMLToMarkdown(result.(string))

							case "google":
								realRes = utils.HTMLToMarkdown(result.(string), utils.GoogleSearchOption)
							case "news":
								realRes = utils.HTMLToMarkdown(result.(string), utils.NewsOption)
							case "screen", "screenshot", "shot":
								if buf, err := page.Screenshot(); err == nil {
									tmpImgBase64 := base64.StdEncoding.EncodeToString(buf)
									realRes = "data:image/png;base64," + tmpImgBase64
								} else {
									realRes = "screenshot error: " + err.Error()
								}
							default:
								realRes = result.(string)
							}
							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url":   url,
								"title": title,
								"page":  realRes,
							})
							flusher.Flush() // 刷新数据到客户端

						} else {
							result, err := page.Content()
							if err != nil {
								return
							}
							realRes := ""
							switch req.Output {
							case "md", "markdown":
								realRes = utils.HTMLToMarkdown(result)

							case "google":
								realRes = utils.HTMLToMarkdown(result, utils.GoogleSearchOption)
							case "news":
								realRes = utils.HTMLToMarkdown(result, utils.NewsOption)
							case "screen", "screenshot", "shot":
								if buf, err := page.Screenshot(); err == nil {
									tmpImgBase64 := base64.StdEncoding.EncodeToString(buf)
									realRes = "data:image/png;base64," + tmpImgBase64
								} else {
									realRes = "screenshot error: " + err.Error()
								}
							default:
								realRes = result
							}

							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url":   url,
								"title": title,
								"page":  realRes,
							})
							flusher.Flush() // 刷新数据到客户端

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
				realRes := ""
				switch req.Output {
				case "md", "markdown":
					realRes = utils.HTMLToMarkdown(result.(string))

				case "google":
					realRes = utils.HTMLToMarkdown(result.(string), utils.GoogleSearchOption)
				case "news":
					realRes = utils.HTMLToMarkdown(result.(string), utils.NewsOption)
				case "screen", "screenshot", "shot":
					if buf, err := page.Screenshot(); err == nil {
						tmpImgBase64 := base64.StdEncoding.EncodeToString(buf)
						realRes = "data:image/png;base64," + tmpImgBase64
					} else {
						realRes = "screenshot error: " + err.Error()
					}
				default:
					realRes = result.(string)
				}
				json.NewEncoder(w).Encode(gs.Dict[any]{
					"url":   req.URL,
					"title": title,
					"page":  realRes,
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

func IsChinese(r rune) bool {
	return r >= 0x4e00 && r <= 0x9fff || r >= 0xf900 && r <= 0xfaff ||
		r >= 0xff00 && r <= 0xffef || r >= 0x3400 && r <= 0x4dbf ||
		r >= 0x20000 && r <= 0x2a6df || r >= 0x2a700 && r <= 0x2b73f ||
		r >= 0x2b740 && r <= 0x2b81f || r >= 0x2b820 && r <= 0x2ceaf ||
		r >= 0xF900
}

func webTestHandler(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")

	if len(req.URLS) == 0 && req.URL != "" {
		req.URLS = append(req.URLS, req.URL)
	}
	client := http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	r, erru := http.NewRequest("GET", req.URLS[0], nil)
	if erru != nil {
		fmt.Println("[Fatal]err in new request:", erru)
		return
	}
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	if res, err := client.Do(r); err != nil {
		fmt.Println(req.URLS[0], ",err:", err)
	} else {
		buf, err := io.ReadAll(res.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(buf)
	}
	// wait.Wait()
	// json.NewEncoder(w).Encode(all_resplys)
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
	w.Header().Set("Content-Type", "application/json")

	if len(req.URLS) == 0 && req.URL != "" {
		req.URLS = append(req.URLS, req.URL)
	}
	wait := sync.WaitGroup{}
	all_resplys := gs.List[gs.Dict[any]]{}
	for _, u := range req.URLS {
		wait.Add(1)
		fmt.Println("spide :", u)
		time.Sleep(10 * time.Millisecond)
		go func(url string) {
			// get base url from url.
			// baseURL := strings.Join(strings.Split(url, "/")[:3], "/")
			defer wait.Done()
			client := http.Client{
				Timeout: 60 * time.Second,
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}
			req, erru := http.NewRequest("GET", url, nil)
			if erru != nil {
				fmt.Println("[Fatal]err in new request:", erru)
				return
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
			if res, err := client.Do(req); err != nil {
				fmt.Println(url, ",err:", err)
			} else {
				// be := bytes.NewBuffer([]byte{})
				title := ""
				buf, err := io.ReadAll(res.Body)
				if err != nil {
					fmt.Println(url, ",err:", err)
					return
				}
				res.Body.Close()
				attrs, err := extractMeta(string(buf))
				if err != nil {
					fmt.Println(url, ",err:", err)
					return
				}
				title = TitleRegex.FindString(string(buf))

				fsf := strings.SplitN(title, ">", 2)
				if len(fsf) == 2 {
					title = fsf[1]
				}
				fsf = strings.SplitN(title, "<", 2)
				if len(fsf) == 2 {
					title = fsf[0]
				}

				// noiframe := NoIframe.ReplaceAllString(string(buf), "")
				// nosvg := NoSVG.ReplaceAllString(noiframe, "")
				// nocss := NocssRegex.ReplaceAllString(nosvg, "")
				// nojs := NoJavascriptRegex.ReplaceAllString(nocss, "")
				// noul := ULRegex.ReplaceAllString(nojs, "")
				// nolink := LinkRegex.ReplaceAllString(noul, "")
				// nofooter := FooterRegex.ReplaceAllString(nolink, "")
				// nofoot := FootRegex.ReplaceAllString(nofooter, "")

				nofoot := filter_garbage(string(buf))
				backupTime := []string{}
				for _, r := range TimeRegex.FindAllString(nofoot, -1) {
					fs := strings.Split(r, ">")
					if len(fs) > 1 {
						fs2 := strings.Split(fs[1], "</time")
						backupTime = append(backupTime, fs2[0])
					}
				}
				domDocTest := html.NewTokenizer(strings.NewReader(nofoot))
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
						if previousStartTokenTest.Data == "li" {
							continue
						}
						if previousStartTokenTest.Data == "a" {
							continue
						}
						if previousStartTokenTest.Data == "style" {
							continue
						}

						if previousStartTokenTest.Data == "title" {
							if title == "" {
								title = strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
							}
							continue
						}

						TxtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
						// tl := len(TxtContent)
						cl := 0
						isch := false
						other := 0
						for _, c := range TxtContent {
							if IsChinese(c) {
								cl++
							} else {
								other++
							}
						}
						if cl*2 > other {
							isch = true
						}
						if isch {
							if cl < 8 {
								continue
							}
						} else {
							words := strings.Fields(TxtContent)
							if len(words) < 10 {
								continue
							}
						}

						// fmt.Println("isch:", isch, "cl:", cl, "ol:", other, "tl:", tl, "txt:", TxtContent)

						texts = append(texts, TxtContent)
					}
				}

				all_resplys = append(all_resplys, gs.Dict[any]{
					"url":         url,
					"status":      res.Status,
					"title":       title,
					"content":     strings.Join(texts, "\n"),
					"meta":        attrs,
					"backup_time": backupTime,
				})

				// rawbuf := []byte("data: " + be.String())
				// n, err := w.Write(rawbuf)
				// if err != nil {
				// 	fmt.Println("reply line err:", err)
				// 	return
				// }
				// buf2 := rawbuf[n:]
				// for n != len(rawbuf) {
				// 	nt, err := w.Write(rawbuf[n:])
				// 	if err != nil {
				// 		fmt.Println("reply continue line err:", err)
				// 		return
				// 	}
				// 	n += nt
				// }
				// flush.Flush()
				return
			}
		}(u)
	}
	wait.Wait()

	json.NewEncoder(w).Encode(all_resplys)

}

func webNewsStreamHandler(w http.ResponseWriter, r *http.Request) {
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
			client := http.Client{
				Timeout: 60 * time.Second,
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
			if res, err := client.Do(req); err != nil {
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
				attrs, err := extractMeta(string(buf))
				if err != nil {
					fmt.Println(url, ",err:", err)
					return
				}
				nofoot := filter_garbage(string(buf))
				backupTime := []string{}
				for _, r := range TimeRegex.FindAllString(nofoot, -1) {
					fs := strings.Split(r, ">")
					if len(fs) > 1 {
						fs2 := strings.Split(fs[1], "</time")
						backupTime = append(backupTime, fs2[0])
					}
				}
				domDocTest := html.NewTokenizer(strings.NewReader(nofoot))
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
						if previousStartTokenTest.Data == "li" {
							continue
						}
						if previousStartTokenTest.Data == "a" {
							continue
						}
						TxtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
						// tl := len(TxtContent)
						cl := 0
						isch := false
						other := 0
						for _, c := range TxtContent {
							if IsChinese(c) {
								cl++
							} else {
								other++
							}
						}
						if cl*2 > other {
							isch = true
						}
						if isch {
							if cl < 8 {
								continue
							}
						} else {
							words := strings.Fields(TxtContent)
							if len(words) < 10 {
								continue
							}
						}

						// fmt.Println("isch:", isch, "cl:", cl, "ol:", other, "tl:", tl, "txt:", TxtContent)

						texts = append(texts, TxtContent)
					}
				}

				json.NewEncoder(be).Encode(gs.Dict[any]{
					"url":    url,
					"status": res.Status,
					// "body":    nosvg,
					"backup_time": backupTime,
					"content":     strings.Join(texts, "\n"),
					"meta":        attrs,
				})

				rawbuf := []byte("data: " + be.String())
				n, err := w.Write(rawbuf)
				if err != nil {
					fmt.Println("reply line err:", err)
					return
				}
				// buf2 := rawbuf[n:]
				for n != len(rawbuf) {
					nt, err := w.Write(rawbuf[n:])
					if err != nil {
						fmt.Println("reply continue line err:", err)
						return
					}
					n += nt
				}

				flush.Flush()
				return
			}
		}(u)
	}
	wait.Wait()

}

func IsOverDomain(uri string, domain string) bool {

	if strings.Contains(domain, "://") {
		domain = strings.Split(domain, "/")[2]
	}
	if strings.Contains(uri, "://") {
		uri = strings.Split(uri, "/")[2]
	}

	fs := strings.Split(domain, ".")
	baseDo := strings.Join(fs[len(fs)-2:], ".")
	fs2 := strings.Split(uri, ".")
	baseUr := strings.Join(fs2[len(fs2)-2:], ".")
	return strings.Contains(baseDo, baseUr)
}

func webChannelHandler(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Connection", "keep-alive")

	if len(req.URLS) == 0 && req.URL != "" {
		req.URLS = append(req.URLS, req.URL)
	}
	waitg := sync.WaitGroup{}
	allReplyResponse := gs.List[Link]{}
	for _, u := range req.URLS {
		waitg.Add(1)
		go func(url string) {
			defer waitg.Done()
			// get base url from url.
			url_fs := strings.Split(url, "/")
			if len(url_fs) < 3 {
				return
			}
			baseURL := strings.Join(url_fs[:3], "/")

			client := http.Client{
				Timeout: 60 * time.Second,
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
			if res, err := client.Do(req); err != nil {
				fmt.Println(url, " ,err:", err)
			} else {
				buf, err := io.ReadAll(res.Body)
				if err != nil {
					fmt.Println(url, " ,err:", err)
					return
				}
				res.Body.Close()
				noiframe := NoIframe.ReplaceAllString(string(buf), "")
				nojs := NoJavascriptRegex.ReplaceAllString(noiframe, "")
				nocss := NocssRegex.ReplaceAllString(nojs, "")
				nosvg := NoSVG.ReplaceAllString(nocss, "")
				links := LinkRegex.FindAllString(nosvg, -1)
				// ls := gs.List[Link]{}
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
					e := allReplyResponse.Any(func(no int, i Link) bool {
						return i.Url == href
					})
					if e.Url != "" {
						continue
					}
					href_fs := strings.Split(href, "/")
					if len(href_fs) < 3 {
						continue
					}
					uri := strings.Join(href_fs[3:], "/")
					if len(uri) < 3 {
						continue
					}
					fs = strings.Split(uri, "/")
					lastFile := fs[len(fs)-1]
					fs = strings.Split(lastFile, ".")
					ex := fs[len(fs)-1]
					if strings.Contains(ex, "?") {
						ex = strings.Split(ex, "?")[0]
					}
					switch ex {
					case "jpg", "png", "icon", "svg", "ico", "raw", "mp4", "jpeg", "gif", "pdf", "docx", "doc", "xlsx", "xls", "zip", "rar", "avi", "mp3", "mkv":
						continue
					}
					if !IsOverDomain(href, url) {
						continue
					}
					if len(strings.TrimPrefix(href, baseURL)) > 20 {
						continue
					}
					if title == "" {
						continue
					}
					if len(title) > 21 {
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

					// fmt.Println("success href:", href)
					allReplyResponse = append(allReplyResponse, Link{
						Source:  url,
						Url:     href,
						Content: title,
						Text:    strings.TrimSpace(puretext),
					})

				}

			}
		}(u)

	}
	waitg.Wait()

	json.NewEncoder(w).Encode(gs.Dict[any]{
		"links": allReplyResponse,
	})

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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Connection", "keep-alive")

	if len(req.URLS) == 0 && req.URL != "" {
		req.URLS = append(req.URLS, req.URL)
	}
	waitg := sync.WaitGroup{}
	allReplyResponse := gs.List[Link]{}
	for _, u := range req.URLS {
		waitg.Add(1)
		go func(url string) {
			// get base url from url.
			defer waitg.Done()
			url_fs := strings.Split(url, "/")
			if len(url_fs) < 3 {
				return
			}
			baseURL := strings.Join(url_fs[:3], "/")

			client := http.Client{
				Timeout: 60 * time.Second,
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
			if res, err := client.Do(req); err != nil {
				fmt.Println(url, " ,err:", err)
			} else {
				buf, err := io.ReadAll(res.Body)
				if err != nil {
					fmt.Println(url, " ,err:", err)
					return
				}
				res.Body.Close()
				noiframe := NoIframe.ReplaceAllString(string(buf), "")
				nojs := NoJavascriptRegex.ReplaceAllString(noiframe, "")
				nocss := NocssRegex.ReplaceAllString(nojs, "")
				nosvg := NoSVG.ReplaceAllString(nocss, "")
				links := LinkRegex.FindAllString(nosvg, -1)
				// ls := gs.List[Link]{}
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
					href_fs := strings.Split(href, "/")
					if len(href_fs) < 3 {
						continue
					}
					uri := strings.Join(href_fs[3:], "/")
					if len(uri) < 3 {
						continue
					}
					fs = strings.Split(uri, "/")
					lastFile := fs[len(fs)-1]
					fs = strings.Split(lastFile, ".")
					ex := fs[len(fs)-1]
					if strings.Contains(ex, "?") {
						ex = strings.Split(ex, "?")[0]
					}
					switch ex {
					case "jpg", "png", "icon", "svg", "ico", "raw", "mp4", "jpeg", "gif", "pdf", "docx", "doc", "xlsx", "xls", "zip", "rar":
						continue
					}
					if !IsOverDomain(href, url) {
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
					// fmt.Println("success href:", href)
					allReplyResponse = append(allReplyResponse, Link{
						Source:  url,
						Url:     href,
						Content: title,
						Text:    strings.TrimSpace(puretext),
					})

				}

			}
		}(u)

	}
	waitg.Wait()

	json.NewEncoder(w).Encode(gs.Dict[any]{
		"links": allReplyResponse,
	})

}

func weblinkStreamHandler(w http.ResponseWriter, r *http.Request) {
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
			defer waitg.Done()
			// get base url from url.
			url_fs := strings.Split(url, "/")
			if len(url_fs) < 3 {
				return
			}
			baseURL := strings.Join(url_fs[:3], "/")

			client := http.Client{
				Timeout: 60 * time.Second,
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
			if res, err := client.Do(req); err != nil {
				be := bytes.NewBuffer([]byte{})
				json.NewEncoder(be).Encode(gs.Dict[any]{
					"url":    u,
					"status": -1,
					"err":    err.Error(),
				})
				fmt.Println("err:", be.String())
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
					fmt.Println("err:", be.String())
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
					href_fs := strings.Split(href, "/")
					if len(href_fs) < 3 {
						continue
					}
					uri := strings.Join(href_fs[3:], "/")
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
					fmt.Println("success href:", href)
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
				rawbuf := []byte("data: " + be.String())
				n, err := w.Write(rawbuf)
				if err != nil {
					fmt.Println("reply line err:", err)
					return
				}
				// buf2 := rawbuf[n:]
				for n != len(rawbuf) {
					nt, err := w.Write(rawbuf[n:])
					if err != nil {
						fmt.Println("reply continue line err:", err)
						return
					}
					n += nt
				}
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
		res := utils.HTMLToMarkdown(result.(string), utils.GoogleSearchOption)
		// fmt.Println("result:", result)

		if items, err := utils.SearchItemMarkdownToJson(res); err != nil {
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
							// itemResult, err := page.Evaluate(TEXT_JS)
							result, err := page.Content()

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
							realRes := utils.HTMLToMarkdown(result)

							json.NewEncoder(w).Encode(gs.Dict[any]{
								"url":        item.Url,
								"title":      title,
								"screenshot": screenPath,
								"page":       realRes,
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

package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"gitee.com/dark.H/gs"
	"github.com/playwright-community/playwright-go"
)

type Options struct {
	UserAgent string
	LoadImage bool
	Proxy     string
}

type Console struct {
	// context playwright.BrowserContext
	Browser     playwright.Browser
	Pages       gs.Dict[playwright.Page]
	PagesScreen gs.Dict[gs.Str]
	BigError    error
	PagesErrors gs.Dict[error]
}

func NewBrowserContext() (*Console, error) {
	con := new(Console)
	pw, err := playwright.Run()
	if err != nil {
		playwright.Install()
		pw, err = playwright.Run()
		if err != nil {
			return nil, err
		}

	}
	// con.context = pw
	browser, err := pw.Chromium.Launch()
	if err != nil {
		return nil, err
	}
	con.Browser = browser
	con.Pages = make(gs.Dict[playwright.Page])
	con.PagesErrors = make(gs.Dict[error])
	con.PagesScreen = make(gs.Dict[gs.Str])
	return con, nil
}

func (bc *Console) Open(url string, after func(screenPath string, page playwright.Page, res playwright.Response), options ...Options) *Console {
	if bc.Browser.IsConnected() {
		var page playwright.Page
		var err error
		loadingImg := false
		if options != nil {
			op := playwright.BrowserNewContextOptions{}
			if options[0].UserAgent != "" {
				op.UserAgent = &options[0].UserAgent
			}
			if options[0].Proxy != "" {
				op.Proxy = &playwright.Proxy{
					Server: options[0].Proxy,
				}
			}
			if options[0].LoadImage {
				loadingImg = true

			}
			browserContext, err := bc.Browser.NewContext(op)
			if err != nil {
				bc.BigError = err
				return bc
			}
			defer browserContext.Close()
			page, err = browserContext.NewPage()
			if err != nil {
				bc.BigError = err
				return bc
			}
		} else {
			page, err = bc.Browser.NewPage(playwright.BrowserNewPageOptions{})
			if err != nil {
				bc.BigError = err
				return bc
			}
		}
		defer page.Close()
		gs.Str("open:" + url).Color("g").Println("page")
		st := time.Now()
		timeOut := float64(50000)
		res, err := page.Goto(url, playwright.PageGotoOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
			Timeout:   &timeOut,
		})
		if !loadingImg {
			page.Route("**/*.(png|jpg|jpeg|gif|svg|jpg)", func(route playwright.Route) {
				// if route.Request().ResourceType() == "image" {
				route.Abort()
				// }
			})
		}

		// bc.pages.Set(url, page)
		if err != nil {
			bc.PagesErrors[url] = err
			return bc
		}

		id := gs.TMP.RandStr(16)
		screenPath := gs.TMP.PathJoin(string(id.Add(".png"))).String()

		usd := time.Since(st)
		gs.Str("[open:" + usd.String() + "] open:" + url).Color("g").Println("page")
		time.Sleep(time.Second * 1)
		buf, err := page.Screenshot(playwright.PageScreenshotOptions{
			Path: &screenPath,
		})
		if err != nil {
			bc.PagesErrors[url] = err
			return bc
		} else {
			fmt.Println("screenshot: ", len(buf))
		}
		gs.Str("screenshot: " + screenPath).Color("g").Println("page")
		bc.PagesScreen[url] = gs.Str(screenPath)

		bc.Pages[url] = page
		// res.
		after(screenPath, page, res)
	}
	return bc
}

func (bc *Console) OpenNoScreen(url string, after func(screenPath string, page playwright.Page, res playwright.Response), options ...Options) *Console {
	if bc.Browser.IsConnected() {
		var page playwright.Page
		var err error
		loadingImg := false
		if options != nil {
			op := playwright.BrowserNewContextOptions{}
			if options[0].UserAgent != "" {
				op.UserAgent = &options[0].UserAgent
			}
			if options[0].Proxy != "" {
				op.Proxy = &playwright.Proxy{
					Server: options[0].Proxy,
				}
			}
			if options[0].LoadImage {
				loadingImg = true

			}
			browserContext, err := bc.Browser.NewContext(op)
			if err != nil {
				bc.BigError = err
				return bc
			}
			defer browserContext.Close()
			page, err = browserContext.NewPage()
			if err != nil {
				bc.BigError = err
				return bc
			}
		} else {
			page, err = bc.Browser.NewPage(playwright.BrowserNewPageOptions{})
			if err != nil {
				bc.BigError = err
				return bc
			}
		}
		defer page.Close()
		// st := time.Now()
		timeOut := float64(50000)
		res, err := page.Goto(url, playwright.PageGotoOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
			Timeout:   &timeOut,
		})
		if !loadingImg {
			page.Route("**/*.(png|jpg|jpeg|gif|svg|jpg)", func(route playwright.Route) {
				// if route.Request().ResourceType() == "image" {
				route.Abort()
				// }
			})
		}

		// bc.pages.Set(url, page)
		if err != nil {
			bc.PagesErrors[url] = err
			// return bc
		}
		// bc.Pages[url] = page
		// res.
		gs.Str("open:" + url).Color("g").Println("page")
		after("", page, res)
	}
	return bc
}

func (bc *Console) WithPage(url string, options ...Options) (page playwright.Page, err error) {
	if bc.Browser.IsConnected() {

		if options != nil {
			op := playwright.BrowserNewContextOptions{}
			if options[0].UserAgent != "" {
				op.UserAgent = &options[0].UserAgent
			}
			if options[0].Proxy != "" {
				op.Proxy = &playwright.Proxy{
					Server: options[0].Proxy,
				}
			}
			browserContext, err := bc.Browser.NewContext(op)
			if err != nil {
				bc.BigError = err
				return nil, err
			}
			page, err = browserContext.NewPage()
			if err != nil {
				bc.BigError = err
				return nil, err
			}
		} else {
			ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"
			op := playwright.BrowserNewContextOptions{
				UserAgent: &ua,
			}
			browserContext, err := bc.Browser.NewContext(op)
			if err != nil {
				bc.BigError = err
				return nil, err
			}
			page, err = browserContext.NewPage()
			if err != nil {
				bc.BigError = err
				return nil, err
			}
		}

		gs.Str("open:" + url).Color("g").Println("page")
		// st := time.Now()
		_, err = page.Goto(url, playwright.PageGotoOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		})
		// bc.pages.Set(url, page)
		if err != nil {
			bc.PagesErrors[url] = err
			return nil, err
		}
	}
	return
}

func (bc *Console) ScreenPage(page playwright.Page) (shotPath string) {
	id := gs.TMP.RandStr(16)
	screenPath := gs.TMP.PathJoin(string(id.Add(".png"))).String()
	// usd := time.Since(st)
	// gs.Str("[open:" + usd.String() + "] open:" + url).Color("g").Println("page")
	time.Sleep(time.Second * 1)
	page.Screenshot(playwright.PageScreenshotOptions{
		Path: &screenPath,
	})
	shotPath = screenPath
	return
}

func (bc *Console) ActionChains(url string, chains string, reply ...func(msg string)) (shotPaths []string) {
	bc.Open(url, func(screenPath string, page playwright.Page, res playwright.Response) {
		// before action chains if exists
		// shotPath = screenPath
		shotPaths = append(shotPaths, screenPath)
		actionsChains := gs.Str(chains).Trim().Split("|")
		for _, actionChain := range actionsChains {
			time.Sleep(300 * time.Millisecond)
			if !actionChain.In(":") {
				continue
			}

			fs := actionChain.Trim().Split(":")

			oper := fs[0].Trim()
			css := fs[1].Trim()
			switch strings.ToLower(oper.Str()) {
			case "c":
				page.Locator(css.Str()).Click()
			case "click":
				page.Locator(css.Str()).Click()
			case "btn":
				page.Locator(css.Str()).Click()
			case "fill":
				if len(fs) == 3 {
					inputValue := fs[2].Trim()
					page.Locator(css.Str()).Fill(inputValue.String())

				} else {
					if reply != nil {
						reply[0](fmt.Sprintf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain))
					}
					bc.PagesErrors[url] = fmt.Errorf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain)
					return
				}
			case "input":
				if len(fs) == 3 {
					inputValue := fs[2].Trim()
					page.Locator(css.Str()).Fill(inputValue.String())

				} else {
					bc.PagesErrors[url] = fmt.Errorf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain)
					return
				}
			case "i":
				if len(fs) == 3 {
					inputValue := fs[2].Trim()
					page.Locator(css.Str()).Fill(inputValue.String())

				} else {
					bc.PagesErrors[url] = fmt.Errorf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain)
					return
				}
			case "press":
				if len(fs) == 3 {
					inputValue := fs[2].Trim()
					page.Locator(css.Str()).Press(inputValue.String())
				}
			case "scroll":
				page.Locator(css.Trim().Str()).ScrollIntoViewIfNeeded()
			case "screenshot":
				for i := 0; i < 10; i++ {
					pi := gs.TMP.PathJoin(gs.TMP.RandStr(16).String() + ".png").String()
					page.Screenshot(playwright.PageScreenshotOptions{
						Path: &pi,
					})
					time.Sleep(500 * time.Millisecond)
					shotPaths = append(shotPaths, pi)
				}
			case "wait":
				waitSec, err := strconv.Atoi(css.Str())
				if err != nil {
					time.Sleep(time.Second * 1)
				} else {
					time.Sleep(time.Second * time.Duration(waitSec))
				}

			}

			gs.Str("action: " + actionChain).Color("g").Println("action")
		}
		pi := gs.TMP.PathJoin(gs.TMP.RandStr(16).String() + ".png").String()
		page.Screenshot(playwright.PageScreenshotOptions{
			Path: &pi,
		})
		time.Sleep(500 * time.Millisecond)
		shotPaths = append(shotPaths, pi)
	})
	return
}

func (bc *Console) ActionChainsWithPage(page playwright.Page, chains string, reply ...func(msg string)) (shotPath []string) {
	// before action chains if exists
	screenPath := gs.TMP.RandStr(16).String() + ".png"
	// shotPath = screenPath
	actionsChains := gs.Str(chains).Trim().Split("|")
	for _, actionChain := range actionsChains {
		time.Sleep(200 * time.Millisecond)
		if !actionChain.In(":") {
			continue
		}

		fs := actionChain.Trim().Split(":")

		oper := fs[0].Trim()
		css := gs.Str("")
		inputValue := gs.Str("")
		if len(fs) > 1 {
			css = fs[1].Trim()
		}
		if len(fs) > 2 {
			inputValue = fs[2].Trim()
		}

		switch strings.ToLower(oper.Str()) {
		case "c":
			page.Locator(css.Str()).Click()
		case "click":
			page.Locator(css.Str()).Click()
		case "btn":
			page.Locator(css.Str()).Click()
		case "fill":
			if len(fs) == 3 {
				// inputValue := fs[2].Trim()
				gs.Str("fill: '" + inputValue.String() + "'").Color("g").Println("action")
				page.Locator(css.Str()).Fill(inputValue.String())

			} else {
				if reply != nil {
					reply[0](fmt.Sprintf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain))
				}
				// bc.PagesErrors[url] = fmt.Errorf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain)
				return
			}
		case "input":
			if len(fs) == 3 {
				// inputValue := fs[2].Trim()
				gs.Str("fill: '" + inputValue.String() + "'").Color("g").Println("action")
				page.Locator(css.Str()).Fill(inputValue.String())

			} else {
				reply[0](fmt.Sprintf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain))
				return
			}
		case "i":
			if len(fs) == 3 {
				// inputValue := fs[2].Trim()
				gs.Str("fill: '" + inputValue.String() + "'").Color("g").Println("action")
				page.Locator(css.Str()).Fill(inputValue.String())

			} else {
				reply[0](fmt.Sprintf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain))
				return
			}
		case "press":
			if len(fs) == 3 {
				// inputValue := fs[2].Trim()
				page.Locator(css.Str()).Press(inputValue.String())
			}
		case "scroll":
			page.Locator(css.Trim().Str()).ScrollIntoViewIfNeeded()
		case "screenshot":
			for i := 0; i < 5; i++ {
				pi := gs.TMP.PathJoin(gs.TMP.RandStr(16).String() + ".png").String()
				page.Screenshot(playwright.PageScreenshotOptions{
					Path: &pi,
				})
				time.Sleep(100 * time.Millisecond)
				shotPath = append(shotPath, pi)
			}
		case "s":
			for i := 0; i < 5; i++ {
				pi := gs.TMP.PathJoin(gs.TMP.RandStr(16).String() + ".png").String()
				page.Screenshot(playwright.PageScreenshotOptions{
					Path: &pi,
				})
				time.Sleep(100 * time.Millisecond)
				shotPath = append(shotPath, pi)
			}

		case "wait":
			waitSec, err := strconv.Atoi(css.Str())
			if err != nil {
				time.Sleep(time.Second * 1)
			} else {
				time.Sleep(time.Second * time.Duration(waitSec))
			}

		}

		gs.Str("action: " + actionChain).Color("g").Println("action")
	}
	shotPath = append(shotPath, screenPath)
	page.Screenshot(playwright.PageScreenshotOptions{
		Path: &screenPath,
	})

	return

}

func (bc *Console) Account(accountTemplate *AccountModel, account string, valid func(ok bool, result string)) (shotPath []string) {
	if accountTemplate.Before == "CURL" && gs.Str(accountTemplate.AccountCss).In("${KEY}") {
		data := gs.Str(accountTemplate.AccountCss).Replace("${KEY}", account).String()
		if res, err := http.Post(accountTemplate.Url, "application/json", strings.NewReader(data)); err != nil {
			valid(false, err.Error())
		} else {
			buf, err := io.ReadAll(res.Body)
			if err == nil {
				ds := make(gs.Dict[any])
				if err = json.Unmarshal(buf, &ds); err == nil {
					msg := ""
					for k := range ds {
						if len(msg) > 0 {
							msg += "\n"
						}
						msg += fmt.Sprintf("%s: %v ", k, ds[k])
					}
					valid(true, msg)
					return
				}
				valid(true, string(buf))
			} else {
				valid(false, string(err.Error()))
			}

		}
	} else {
		bc.Open(accountTemplate.Url, func(screenPath string, page playwright.Page, res playwright.Response) {
			// before action chains if exists
			// img = screenPath
			shotPath = append(shotPath, screenPath)
			actionsChains := accountTemplate.ActionChains()
			// actionsChains := gs.Str(chains).Trim().Split("|")
			for _, actionChain := range actionsChains {
				time.Sleep(1 * time.Second)
				if !actionChain.In(":") {
					continue
				}

				fs := actionChain.Trim().Split(":")

				oper := fs[0].Trim()
				css := gs.Str("")
				inputValue := gs.Str("")
				if len(fs) > 1 {
					css = fs[1].Trim()
				}
				if len(fs) > 2 {
					inputValue = fs[2].Trim()
				}

				switch strings.ToLower(oper.Str()) {
				case "c":
					page.Locator(css.Str()).Click()
				case "click":
					page.Locator(css.Str()).Click()
				case "btn":
					page.Locator(css.Str()).Click()
				case "fill":
					if len(fs) == 3 {
						// inputValue := fs[2].Trim()
						gs.Str("fill: '" + inputValue.String() + "'").Color("g").Println("action")
						page.Locator(css.Str()).Fill(inputValue.String())

					} else {
						// if reply != nil {
						// 	reply[0](fmt.Sprintf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain))
						// }
						bc.PagesErrors[accountTemplate.Url] = fmt.Errorf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain)
						return
					}
				case "input":
					if len(fs) == 3 {
						// inputValue := fs[2].Trim()
						gs.Str("fill: '" + inputValue.String() + "'").Color("g").Println("action")
						page.Locator(css.Str()).Fill(inputValue.String())

					} else {
						// reply[0](fmt.Sprintf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain))
						return
					}
				case "i":
					if len(fs) == 3 {
						// inputValue := fs[2].Trim()
						gs.Str("fill: '" + inputValue.String() + "'").Color("g").Println("action")
						page.Locator(css.Str()).Fill(inputValue.String())

					} else {
						// reply[0](fmt.Sprintf("fill error lacked fill value : %s [exm: 'fill:css:value' ]", actionChain))
						return
					}
				case "press":
					if len(fs) == 3 {
						// inputValue := fs[2].Trim()
						page.Locator(css.Str()).Press(inputValue.String())
					}
				case "scroll":
					page.Locator(css.Trim().Str()).ScrollIntoViewIfNeeded()
				case "screenshot":
					for i := 0; i < 10; i++ {
						pi := gs.TMP.PathJoin(gs.TMP.RandStr(16).String() + ".png").String()
						page.Screenshot(playwright.PageScreenshotOptions{
							Path: &pi,
						})
						time.Sleep(500 * time.Millisecond)
						shotPath = append(shotPath, pi)
					}

				case "wait":
					waitSec, err := strconv.Atoi(css.Str())
					if err != nil {
						time.Sleep(time.Second * 1)
					} else {
						time.Sleep(time.Second * time.Duration(waitSec))
					}

				}

				gs.Str("action: " + actionChain).Color("g").Println("action")
			}
			pend := gs.TMP.PathJoin(string(gs.TMP.RandStr(16)) + ".png").String()
			page.Screenshot(playwright.PageScreenshotOptions{
				Path: &pend,
			})
			shotPath = append(shotPath, pend)

			// focus input by css
			gs.Str(accountTemplate.AccountCss + " : " + account).Color("g").Println("action css")
			loc := page.Locator(accountTemplate.AccountCss)
			loc.Fill(accountTemplate.PhoneCode + account)
			// time.Sleep(time.Second * 1)
			time.Sleep(time.Second * 3)
			loc.Press("Enter")
			time.Sleep(time.Second * 3)
			pend = gs.TMP.PathJoin(string(gs.TMP.RandStr(16)) + ".png").String()
			page.Screenshot(playwright.PageScreenshotOptions{
				Path: &pend,
			})
			shotPath = append(shotPath, pend)

			gs.Str("account type:" + screenPath).Color("g").Println(accountTemplate.Name)
			reverse := false
			valids := accountTemplate.Valid
			if gs.Str(valids).Trim().StartsWith("!") {
				reverse = true
				valids = valids[1:]
			}
			time.Sleep(3 * time.Second)
			pend = gs.TMP.PathJoin(string(gs.TMP.RandStr(16)) + ".png").String()
			page.Screenshot(playwright.PageScreenshotOptions{
				Path: &pend,
			})
			gs.Str("Finally: " + screenPath).Color("g").Println(accountTemplate.Name)
			if fs, err := page.Locator(valids).All(); err != nil {

				valid(false, err.Error())
			} else {
				if len(fs) > 0 {
					if reverse {
						valid(false, "不存在")
					} else {
						valid(true, "存在")
					}
				} else {
					if reverse {
						valid(true, "存在")
					} else {
						valid(false, "不存在")
					}
				}
			}
			if locStr, err := page.Locator("html").InnerHTML(); err == nil {
				gs.Str(locStr).ToFile(screenPath + ".html")
				gs.Str("html: " + screenPath + ".html").Color("g").Println()
			}

		})
	}
	if bc.PagesErrors[accountTemplate.Url] != nil {
		fmt.Println("error: ", bc.PagesErrors[accountTemplate.Url])
	}
	return
}

func (bc *Console) Close() {
	bc.Browser.Close()
}

func (bc *Console) Check(account string, confPath string, notify func(name string, result bool, resultStr string, err error)) map[string]bool {

	models, err := ReadAccountModels(confPath)
	if err != nil {
		fmt.Println("error: ", err)
		return nil
	}
	wait := sync.WaitGroup{}
	fmt.Println("check types:", len(models))
	result := make(map[string]bool)
	for _, model := range models {
		wait.Add(1)
		go func(w *sync.WaitGroup, m *AccountModel) {
			defer w.Done()
			fmt.Println("check:", m.Name)
			bc.Account(m, account, func(ok bool, result string) {
				// result[m.Name] = ok
				notify(m.Name, ok, result, nil)
			})
			if err, ok := bc.PagesErrors[m.Url]; ok {
				gs.Str(err.Error()).Color("r").Println(m.Name)
				notify(m.Name, false, "不存在", err)
			}
		}(&wait, model)
	}
	time.Sleep(time.Second * 1)
	wait.Wait()
	return result
}

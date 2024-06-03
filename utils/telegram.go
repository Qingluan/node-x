package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"gitee.com/dark.H/gs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/net/proxy"
)

type TgAuth struct {
	BotToken     string `json:"token"`
	Proxy        string `json:"proxy"`
	GroupTitle   string `json:"group_title"`
	GroupID      int64  `json:"chat_id"`
	WaitSec      int    `json:"wait"`
	exit         bool
	GroupChat    *tgbotapi.Chat
	callBack     func(question string, ifok bool)
	Bot          *tgbotapi.BotAPI
	textcallBack func(user, msg string)
	Debug        bool
	TmpFiles     gs.Dict[string]
}

func NewTgAuth(token, title string) (t *TgAuth) {
	t = new(TgAuth)
	t.BotToken = token
	t.GroupTitle = title
	t.WaitSec = 9
	var err error
	t.Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	t.Bot.Debug = t.Debug
	t.TmpFiles = make(gs.Dict[string])
	return
}

func NewTgAuthWithProxy(token, title, proxyHost string) (t *TgAuth) {
	t = new(TgAuth)
	t.BotToken = token
	t.GroupTitle = title
	t.Proxy = proxyHost
	t.WaitSec = 9
	if gs.Str(proxyHost).StartsWith("socks5://") {
		cc := gs.Str(proxyHost).Split("socks5://")[1].Str()
		proxyDialer, _ := proxy.SOCKS5("tcp", cc, nil, nil)

		t.Bot, _ = tgbotapi.NewBotAPIWithClient(t.BotToken, tgbotapi.APIEndpoint, &http.Client{
			Transport: &http.Transport{
				Dial: proxyDialer.Dial,
			},
		})
	} else {
		t.Bot, _ = tgbotapi.NewBotAPI(t.BotToken)
	}
	t.Bot.Debug = t.Debug
	gs.Str(t.Bot.Token).Color("g").Println("Logined")
	return
}

func NewTgAuthWithProxyDialer(token, title string, proxy proxy.Dialer) (t *TgAuth) {
	t = new(TgAuth)
	t.BotToken = token
	t.GroupTitle = title
	t.WaitSec = 9
	t.Bot, _ = tgbotapi.NewBotAPIWithClient(t.BotToken, tgbotapi.APIEndpoint, &http.Client{
		Transport: &http.Transport{
			Dial: proxy.Dial,
		},
	})
	t.Bot.Debug = t.Debug
	gs.Str(t.Bot.Token).Color("g").Println("Logined")
	return
}

func (auth *TgAuth) WithAGroup() {
	if auth.GroupID != 0 {
		if chat, err := auth.Bot.GetChat(tgbotapi.ChatInfoConfig{
			ChatConfig: tgbotapi.ChatConfig{
				ChatID: auth.GroupID,
				// SuperGroupUsername:"" ,
			},
		}); err == nil {
			gs.Str(chat.Title).Color("g").Println("Find!")
			auth.GroupChat = &chat
			auth.GroupID = chat.ID
		}
	}
}

func (auth *TgAuth) SendFile(fpath string) {
	if gs.Str(fpath).IsExists() && !gs.Str(fpath).IsDir() {
		ufile := tgbotapi.FilePath(fpath)
		doc := tgbotapi.NewDocument(auth.GroupID, ufile)
		doc.AllowSendingWithoutReply = true
		auth.Bot.Send(doc)
	} else {
		auth.Say("not found :" + fpath)
	}
}

func (auth *TgAuth) Exit() {
	auth.exit = true

}

func Downloads(url string, fileName string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// 发送GET请求
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: unexpected status code", resp.StatusCode)
		return
	}

	// 创建一个文件来保存下载的内容
	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer out.Close()

	// 使用io.Copy将响应主体复制到文件
	_, err = io.Copy(out, resp.Body)
}

func (auth *TgAuth) DownloadFile(fileName string) {
	if url, ok := auth.TmpFiles[fileName]; ok {
		auth.Confirm("If Upload:" + fileName)
		auth.OnConfirm(func(question string, ifok bool) {
			if question == "If Upload:"+fileName {
				if ifok {
					auth.Say("开始下载")
					Downloads(url, fileName)
					auth.Say("上传完成")
				}

			}
		})

	}
}

func (auth *TgAuth) UploadFile(filePath string) {
	if strings.Contains(filePath, ".") {
		fs := strings.Split(filePath, ".")
		fileType := fs[len(fs)-1]
		if fileType == "jpg" || fileType == "png" || fileType == "gif" || fileType == "jpeg" {
			auth.UploadImg(filePath)
			return
		}
	}
	if auth.GroupChat != nil && auth.GroupID != 0 {
		file := tgbotapi.FilePath(filePath)
		msg := tgbotapi.NewDocument(auth.GroupID, file)
		auth.Bot.Send(msg)
	}
}

func (auth *TgAuth) UploadImg(path string) {
	if auth.GroupChat != nil && auth.GroupID != 0 {
		file := tgbotapi.FilePath(path)
		msg := tgbotapi.NewPhoto(auth.GroupID, file)
		auth.Bot.Send(msg)
	}
}

func (auth *TgAuth) StartListen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := auth.Bot.GetUpdatesChan(u)

	for update := range updates {
		if auth.exit {
			break
		}
		// update.Message.
		if update.Message != nil && update.Message.Chat.Title == auth.GroupTitle { // If we got a message
			gs.Str("Hit - " + update.Message.Chat.Title).Color("g").Println("info")
			if auth.GroupChat == nil {
				auth.GroupID = update.Message.Chat.ID
				if chat, err := auth.Bot.GetChat(tgbotapi.ChatInfoConfig{
					ChatConfig: tgbotapi.ChatConfig{
						ChatID: auth.GroupID,
						// SuperGroupUsername:"" ,
					},
				}); err == nil {
					gs.Str(chat.Title).Color("g").Println("Find!")
					auth.GroupChat = &chat

					auth.GroupID = chat.ID
				}
			}
			if update.Message.Chat.Title == auth.GroupTitle {
				if auth.textcallBack != nil {
					go auth.textcallBack(update.Message.From.UserName, update.Message.Text)
				}
			}
		} else if update.Poll != nil {
			gs.Str(update.Poll.Question).Color("g").Println()
			if auth.callBack != nil {
				gs.Str(update.Poll.Question + " > call").Color("g").Println()
				go auth.handlerConfirm(update.Poll)
			}
		} else if update.Message != nil {
			fmt.Println(update.Message.Chat)
			if update.Message.Document != nil {
				gs.Str(update.Message.Document.FileName).Color("g").Println()

				file, err := auth.Bot.GetFile(tgbotapi.FileConfig{FileID: update.Message.Document.FileID}) // <-- point of error
				if err == nil {
					url := file.Link(auth.BotToken)
					auth.TmpFiles[update.Message.Document.FileName] = url
					auth.DownloadFile(update.Message.Document.FileName)
				}
			}
		}
	}
}

func (auth *TgAuth) Confirm(question string, answers ...string) {
	if auth.GroupChat != nil && auth.GroupID != 0 {
		msg := tgbotapi.NewPoll(auth.GroupID, question, "Yes", "No")
		msg.CloseDate = auth.WaitSec
		auth.Bot.Send(msg)
		// poolID := msgs.Poll.ID

	}
}

func (auth *TgAuth) Say(text string) {
	if auth.GroupChat != nil && auth.GroupID != 0 {
		_text := ""

		if len(text) > 2048 {
			gs.Str(text).ToFile("test.txt", gs.O_NEW_WRITE)
			qs := []string{}
			gs.Str(text).Split("\n").Every(func(no int, i gs.Str) {
				if i.Trim() != "" {
					ireal := strings.Map(func(r rune) rune {
						if unicode.IsGraphic(r) {
							return r
						}
						return -1
					}, i.Str())
					_text += ireal + "\n"
					if len(_text) > 2000 {
						qs = append(qs, _text)
						fmt.Println(len(_text), _text)
						_text = ""

					}
				}
			})
			if len(_text) > 0 {
				qs = append(qs, _text)
			}

			for _, mm := range qs {
				msg := tgbotapi.NewMessage(auth.GroupID, mm)
				// msg.CloseDate = auth.WaitSec
				fmt.Println("send : ", len(mm))
				_m, err := auth.Bot.Send(msg)
				time.Sleep(500 * time.Microsecond)
				if err != nil {
					gs.Str("can not say !:" + err.Error()).Color("r").Println("telegram")
				} else {
					gs.Str("say :" + _m.Chat.Title).Color("b").Println("telegram")
				}
			}
		} else {
			msg := tgbotapi.NewMessage(auth.GroupID, text)
			// msg.CloseDate = auth.WaitSec
			_m, err := auth.Bot.Send(msg)
			if err != nil {
				gs.Str("can not say !:" + err.Error()).Color("r").Println("telegram")
			} else {
				gs.Str("say :" + _m.Chat.Title).Color("b").Println("telegram")

			}

		}
		// poolID := msgs.Poll.ID

	} else {
		gs.Str("can not found grou to say !" + text).Color("r").Println("telegram")
	}
}

func (auth *TgAuth) OnConfirm(callBack func(question string, ifok bool)) {
	auth.callBack = callBack
}

func (auth *TgAuth) OnText(callBack func(user, text string)) {
	auth.textcallBack = callBack
}

func (auth *TgAuth) handlerConfirm(pool *tgbotapi.Poll) {
	var maxOption tgbotapi.PollOption
	gs.List[tgbotapi.PollOption](pool.Options).Every(func(no int, i tgbotapi.PollOption) {
		if i.VoterCount > maxOption.VoterCount {
			maxOption = i
		}
	})
	if maxOption.Text == "Yes" {
		gs.Str("Y").Println(pool.Question)
		auth.callBack(pool.Question, true)
	} else {
		gs.Str("N").Println(pool.Question)
		auth.callBack(pool.Question, false)
	}
}

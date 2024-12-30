package utils

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type NewsInfo struct {
	PublishTime int64    `json:"publish_time"`
	Summary     string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

func (newsInfo *NewsInfo) Time() string {
	return time.Unix(newsInfo.PublishTime, 0).Format("2006-01-02 15:04:05")
}

type FilterOption struct {
	TextFilter func(tag, text string) bool
	JumpTags   []string
	RemoveTags []string
	Maps       func(doc *goquery.Document, buf *bytes.Buffer)
}

var (
	GoogleSearchOption = &FilterOption{
		RemoveTags: []string{"#sfooter", "[role=navigation]", "#bres", "#searchform", "html>body>h1", "g-menu", ".appbar", "#top_nav"},
		TextFilter: func(tag, text string) bool {

			return text != "Accessibility feedback" && text != "Accessibility help" && text != "Skip to main content"
		},
	}
	NewsOption = &FilterOption{
		RemoveTags: []string{"ul", "nav", "a", "*comment*"},
		TextFilter: func(tag, text string) bool {
			if strings.HasPrefix(tag, "h") {
				// fmt.Println("hhh", len(text), strings.TrimSpace(text))
				return len(text) > 30
			}
			return len(text) > 40
		},
		Maps: func(doc *goquery.Document, buf *bytes.Buffer) {
			title := doc.Find("title").Text()
			if info, err := ExtractNewsInfo(doc); err == nil {
				buf.WriteString(fmt.Sprintf("## %s\n\n", title))
				// from int64 timestample to string
				if len(info.Keywords) > 0 {
					buf.WriteString(fmt.Sprintf("> %s\n", "*"+strings.Join(info.Keywords, "*, *")+"*"))
				}
				buf.WriteString(fmt.Sprintf("> %s\n\n", info.Time()))

			}
		},
	}
)

func parseDate(s string) (time.Time, error) {
	formats := []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05+07:00",
		"2006-01-02 15:04:05",
		"2006-01-02",
		"January 2, 2006",
		"Jan 2, 2006",
		"02 Jan 2006",
		"06/01/02", // 如 "24/01/03" 表示 2024年1月3日
		"01/02/06", // 如 "01/03/24" 表示 2024年1月3日
		"01-02-06",
		"02/01/06",
		"060102", // 无分隔符的日期格式
		"2006/01/02 15:04:05",
		"02-01-2006",
		// 添加更多可能的日期格式
	}
	for _, layout := range formats {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("无法解析日期: %s", s)
}

func ExtractNewsDate(doc *goquery.Document) (int64, error) {
	var date time.Time
	// 从 meta 标签中提取日期
	doc.Find("meta").Each(func(_ int, s *goquery.Selection) {
		content, exists := s.Attr("content")
		if exists {
			// fmt.Println(content)
			t, err := parseDate(content)
			if err == nil {
				date = t
				return // 找到一个日期后停止
			}
		}
	})

	if !date.IsZero() {
		return date.Unix(), nil
	}

	// 从 time 或 datetime 标签中提取日期
	doc.Find("time, [datetime]").Each(func(_ int, s *goquery.Selection) {
		datetime, exists := s.Attr("datetime")
		if exists {
			t, err := parseDate(datetime)
			if err == nil {
				date = t
				return
			}
		}
		// 检查 time 标签的文本内容
		text := s.Text()
		t, err := parseDate(text)
		if err == nil {
			date = t
			return
		}
	})

	if !date.IsZero() {
		return date.Unix(), nil
	}

	// 未找到日期
	return 0, fmt.Errorf("未找到新闻日期")
}

func ExtractNewsInfo(doc *goquery.Document) (info *NewsInfo, err error) {
	info = new(NewsInfo)

	// 提取日期
	var date time.Time
	// 从 meta 标签中提取日期
	doc.Find("head meta").Each(func(_ int, s *goquery.Selection) {
		content, exists := s.Attr("content")
		if exists {
			t, err := parseDate(content)
			if err == nil {
				date = t
				return // 找到一个日期后停止
			}
		}
	})

	if date.IsZero() {
		// 从 time 或 datetime 标签中提取日期
		doc.Find("time, [datetime]").Each(func(_ int, s *goquery.Selection) {
			datetime, exists := s.Attr("datetime")
			if exists {
				t, err := parseDate(datetime)
				if err == nil {
					date = t
					return
				}
			}
			// 检查 time 标签的文本内容
			text := s.Text()
			t, err := parseDate(text)
			if err == nil {
				date = t
				return
			}
		})
	}

	if !date.IsZero() {
		info.PublishTime = date.Unix()
	}

	// 提取摘要
	doc.Find("meta").Each(func(_ int, s *goquery.Selection) {
		content, exists := s.Attr("content")
		name, _ := s.Attr("name")
		property, _ := s.Attr("property")
		if strings.Contains(strings.ToLower(name), "desc") || strings.Contains(strings.ToLower(property), "desc") {
			if exists && content != "" {
				info.Summary = content
				return
			}
		}
	})

	// 提取关键字
	doc.Find("meta").FilterFunction(func(i int, s *goquery.Selection) bool {
		name, _ := s.Attr("name")
		property, _ := s.Attr("property")
		keywords := []string{"keywords", "keyword", "key", "keys"}
		for _, k := range keywords {
			if strings.Contains(strings.ToLower(name), k) || strings.Contains(strings.ToLower(property), k) {
				return true
			}
		}
		return false
	}).Each(func(_ int, s *goquery.Selection) {
		content, exists := s.Attr("content")
		if exists && content != "" {
			keywords := strings.Split(content, ",")
			for _, k := range keywords {
				info.Keywords = append(info.Keywords, strings.TrimSpace(k))
			}
			return
		}
	})

	return info, nil
}

func getDirectText(node *goquery.Selection) string {
	var buf bytes.Buffer
	dd := 0
	node.Contents().Each(func(_ int, s *goquery.Selection) {
		if dd < 1 {
			buf.WriteString(s.Text())
		}
		dd += 1
	})
	return strings.TrimSpace(buf.String())

}

func IsBlockNode(node *goquery.Selection) bool {
	switch node.Nodes[0].Data {
	case "p", "h1", "h2", "h3", "h4", "h5", "h6", "blockquote", "tr", "li":
		return true
	default:
		return false
	}
}

// calculate parent li depth number, return 0 if not li
func ParentLiNum(node *goquery.Selection) int {
	var dd int
	node.ParentsUntil("body").Each(func(_ int, s *goquery.Selection) {
		if s.Is("li") {
			dd += 1
		}
	})
	return dd
}
func IsParentLink(node *goquery.Selection) bool {
	return node.ParentsUntil("body").Is("a")
}

func FirstChildIsH(node *goquery.Selection) (ok bool) {
	used := false
	node.Children().Each(func(i int, s *goquery.Selection) {
		// fmt.Println("first", s.Nodes[0].Data)
		switch s.Nodes[0].Data {

		case "br", "hr":

			// return
		case "h1", "h2", "h3", "h4", "h5", "h6":
			if !used {
				ok = true
				used = true
			}
		default:
			if !used {
				used = true
				ok = false
			}

		}
	})
	return ok
}

func ExtractDate(node *goquery.Selection) string {
	return ""
}

func IsFirstChildImg(node *goquery.Selection) bool {
	return node.Children().First().Is("img")
}

func isValidHref(href string) bool {
	if href == "" || href == "#" || href == "javascript:void(0)" {
		return false
	}
	return true
}

func WriteChildIsH(node *goquery.Selection, buf *bytes.Buffer) {
	used := false
	node.Children().Each(func(i int, s *goquery.Selection) {
		switch s.Nodes[0].Data {

		case "br", "hr":
			// fmt.Println(">>>>>>")
			// return
		case "h1", "h2", "h3", "h4", "h5", "h6":
			if !used {
				// fmt.Println(">>>>>>")
				// ok = true
				switch s.Nodes[0].Data {
				case "h1":
					buf.WriteString("\n\n")
					buf.WriteString("# ")
				case "h2":
					buf.WriteString("\n\n")
					buf.WriteString("## ")
				case "h3":
					buf.WriteString("\n\n")
					buf.WriteString("### ")
				case "h4":
					buf.WriteString("\n\n")
					buf.WriteString("#### ")
				case "h5":
					buf.WriteString("\n\n")
					buf.WriteString("##### ")
				case "h6":
					buf.WriteString("\n\n")
					buf.WriteString("###### ")
				}
				used = true

			}
		default:
			if !used {
				used = true
				// ok = false
			}

		}
	})

}

func IsExcludeTag(node *goquery.Selection, excludeTags ...string) bool {
	tag := node.Nodes[0].Data
	for _, excludeTag := range excludeTags {
		if tag == excludeTag {
			return true
		}
	}
	return false
}

func SearchItemMarkdownToJson(markdownSearch string) (items []*SearchItem, err error) {
	for _, line := range strings.Split(markdownSearch, "\n") {
		if strings.HasPrefix(line, "##") && strings.Contains(line, "(http") && strings.Contains(line, ")") {
			url := strings.Split(strings.Split(line, "(")[1], ")")[0]
			title := strings.TrimSpace(strings.Split(strings.Split(line, "[")[1], "]")[0])
			items = append(items, &SearchItem{
				Title: title,
				Url:   url,
			})
		}
	}
	return
}

func HTMLToMarkdown(html string, options ...*FilterOption) string {

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	// 提取所有<style>标签的内容
	var cssTexts []string
	doc.Find("style").Each(func(i int, s *goquery.Selection) {
		cssTexts = append(cssTexts, s.Text())
	})

	// 提取所有设置display: none的选择器
	var hiddenSelectors []string
	for _, cssText := range cssTexts {
		hiddenSelectors = append(hiddenSelectors, extractHiddenSelectors(cssText)...)
	}

	// 移除被选择器选中的元素
	for _, selector := range hiddenSelectors {
		doc.Find(selector).Remove()
	}
	doc.Find("script").Remove()
	doc.Find("style").Remove()
	doc.Find("footer").Remove()
	for _, option := range options {
		for _, removeTag := range option.RemoveTags {

			doc.Find(removeTag).Each(func(i int, s *goquery.Selection) {
				s.Remove()
			})
		}
	}
	var buf bytes.Buffer
	for _, option := range options {
		if option.Maps != nil {
			option.Maps(doc, &buf)
		}
	}
	var walk func(*goquery.Selection, bool, string, ...string)
	walk = func(selection *goquery.Selection, depth bool, prefixSpace string, extags ...string) {

		// lastDiv := false
		selection.Children().Each(func(i int, s *goquery.Selection) {
			tag := s.Nodes[0].Data
			className := s.AttrOr("class", "")
			idName := s.AttrOr("id", "")
			text := strings.TrimSpace(s.Text())
			for _, option := range options {
				if option.TextFilter != nil && !option.TextFilter(tag, text) {
					return
				}
				if len(option.JumpTags) > 0 && IsExcludeTag(s, option.JumpTags...) {
					return
				}
				for _, removeTag := range option.RemoveTags {
					if strings.HasPrefix(removeTag, "*") && strings.HasSuffix(removeTag, "*") {
						if strings.Contains(className, removeTag[1:len(removeTag)-1]) {
							return
						}
						if strings.Contains(idName, removeTag[1:len(removeTag)-1]) {
							return
						}
					}
				}
			}

			if extags != nil {
				for _, extag := range extags {
					if tag == extag {
						return
					}
				}
			}
			buf.WriteString(prefixSpace)
			switch tag {
			case "h1", "h2", "h3", "h4", "h5", "h6":
				level := tag[1]

				if !IsParentLink(s) {
					buf.WriteString("\n\n")
					buf.WriteString(strings.Repeat("#", int(level-'0')) + " ")
					// fmt.Println("hh:", text, len(text))
					buf.WriteString(text + "\n\n")
				} else {
					buf.WriteString(text)
				}
				walk(s, depth, "")
			case "p":
				buf.WriteString(text + "\n\n")

			case "section":
				buf.WriteString("\n------ \n")
				walk(s, depth, "")
			case "ul", "ol":
				ordered := tag == "ol"
				// prefixNum := 0
				prefixNum := ParentLiNum(s)
				space := strings.Repeat(" ", prefixNum*4)
				if ordered {
					space = fmt.Sprintf("%s%d. ", space, i+1)
				} else {
					space = fmt.Sprintf("%s- ", space)
				}
				walk(s, depth, space)

				buf.WriteString("\n")

			case "table":
				// 处理表头
				thead := s.Find("thead")
				if thead.Size() > 0 {
					thead.Find("tr").Each(func(i int, row *goquery.Selection) {
						row.Find("th").Each(func(j int, th *goquery.Selection) {
							buf.WriteString(th.Text() + " | ")
						})
						buf.WriteString("\n")
						row.Find("th").Each(func(j int, th *goquery.Selection) {
							buf.WriteString("---|")
						})
						buf.WriteString("\n")
					})
				}
				// 处理表体
				tbody := s.Find("tbody")
				if tbody.Size() > 0 {
					tbody.Find("tr").Each(func(i int, row *goquery.Selection) {
						row.Find("td").Each(func(j int, td *goquery.Selection) {
							buf.WriteString(td.Text() + " | ")
						})
						buf.WriteString("\n")
					})
				}
				buf.WriteString("\n")
			case "hr":
				buf.WriteString("\n---\n")

			case "pre", "code":
				// text := s.Text()
				buf.WriteString("```\n" + text + "\n```\n")
			case "blockquote":
				buf.WriteString("> ")
				walk(s, depth, prefixSpace)
			case "a":

				// if href != "#" {

				// text := strings.TrimSpace(s.Text())
				// gs.Str(text).Color("g").Println(s.Children().Size())
				href, _ := s.Attr("href")

				if s.Children().Size() != 0 {
					if FirstChildIsH(s) {
						WriteChildIsH(s, &buf)
						// fmt.Println("is h")
						buf.WriteString(fmt.Sprintf("["))
					} else {
						if isValidHref(href) {
							buf.WriteString(fmt.Sprintf("["))
						}
					}

					walk(s, false, "")
					// fmt.Println(href)
					if isValidHref(href) {
						buf.WriteString(fmt.Sprintf("](%s)\n\n", href))
					}

				} else {
					buf.WriteString(fmt.Sprintf("[%s](%s)\n\n", text, href))
				}

				// buf.WriteString(fmt.Sprintf("[%s](%s)\n\n", gs.Str(text).Color("g"), href))

			case "img":
				src, _ := s.Attr("src")
				alt, _ := s.Attr("alt")
				if !isValidHref(src) {
					return
				}
				if alt == "" {
					alt = "image"
				}
				buf.WriteString(fmt.Sprintf("![%s](%s)", alt, src))
			case "span":
				// tt := s.Text()
				buf.WriteString(text)

				if IsBlockNode(s.Parent()) {
					buf.WriteString("\n")
				}
				walk(s, depth, "")

			default:
				if depth {
					walk(s, depth, "")
				}

			}
			// 跳过没有文本内容的节点
			if text == "" && s.Children().Size() == 0 {
				return
			}
		})
	}

	walk(doc.Selection, true, "")
	return buf.String()
}

func preprocessCSS(cssText string) string {
	// 移除注释
	re := regexp.MustCompile(`/\*.*?\*/`)
	cssText = re.ReplaceAllString(cssText, "")
	// 将多行声明合并成一行
	cssText = strings.ReplaceAll(cssText, "\n", " ")
	return cssText
}

func extractHiddenSelectors(cssText string) []string {
	cssText = preprocessCSS(cssText)
	// 分割规则块，以 "}" 为分隔符
	rules := strings.Split(cssText, "}")
	var selectors []string
	for _, rule := range rules {
		// 检查是否包含 "display: none"
		if strings.Contains(rule, "display: none") {
			// 分割选择器，以 "{" 为分隔符
			parts := strings.Split(rule, "{")
			if len(parts) >= 1 {
				selectorsStr := strings.TrimSpace(parts[0])
				// 分割多个选择器，以 "," 为分隔符
				selList := strings.Split(selectorsStr, ",")
				for _, sel := range selList {
					sel = strings.TrimSpace(sel)
					if sel != "" {
						selectors = append(selectors, sel)
					}
				}
			}
		}
	}
	return selectors
}

// func main() {
// 	htmlbuf, err := os.ReadFile(os.Args[1])
// 	if err != nil {
// 		return
// 	}
// 	markdown := HTMLToMarkdown(string(htmlbuf), NewsOption)
// 	markdown = strings.ReplaceAll(markdown, "\n\n\n\n\n\n\n\n", "\n ------ \n ")
// 	fmt.Println(markdown)
// }

package generator

import (
	"bytes"
	"fmt"
	"os"
	"time"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type articleInfo struct {
	Title      string
	URL        string
	Text       string
	WriterName string
	Date       string
	Comment    bool
}

func GenerateArticle(websiteName, websiteURL, title, url, text, writerName string, date int64, isComment bool) {
	path := "dist/" + url
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			fmt.Println("Create file fail: " + path)
		}
	}
	f, err := os.Create(path + "index.html")
	if err != nil {
		fmt.Println("Create file fail: " + path + "index.html")
		return
	}
	output := blackfriday.Run([]byte(text))
	tmpl := new(bytes.Buffer)
	err = G(tmpl, "article", articleInfo{
		Title:      title,
		URL:        url,
		Text:       string(output),
		WriterName: writerName,
		Date:       time.Unix(date/1000, 0).Format("2006-1-2"),
		Comment:    isComment,
	})
	if err != nil {
		fmt.Println("Execute fail: " + path + "index.html")
		return
	}
	err = G(f, "index", pageIndexParams{
		Tmpl:        tmpl.String(),
		WebsiteName: websiteName,
		WebsiteURL:  websiteURL,
		HeadTitle:   title,
	})
	if err != nil {
		fmt.Println("Execute fail: " + path + "index.html")
		return
	}
	fmt.Println("Execute: " + f.Name())
	f.Close()
}

func RemoveArticle(url string) {
	os.Remove("dist/" + url + "index.html")
}

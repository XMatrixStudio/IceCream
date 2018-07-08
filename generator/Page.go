package generator

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type pageInfo struct {
	PageNum  int
	Page     int
	P1       int
	P2       int
	P3       int
	Articles []ArticleInfo
}

type ArticleInfo struct {
	Title      string
	URL        string
	Date       string
	WriterName string
	Text       string
}

func GeneratePage(websiteName, websiteURL string, websiteArticles, pageNum int, articles []ArticleInfo) {
	var path string
	if pageNum == 1 {
		path = "dist/"
	} else if pageNum > 1 {
		path = "dist/page/" + strconv.Itoa(pageNum) + "/"
	} else {
		return
	}
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
	tmpl := new(bytes.Buffer)
	var p1, p2, p3 int
	page := (websiteArticles + 9) / 10
	if page == 0 {
		page = 1
	}
	if pageNum == 1 {
		p1, p2, p3 = pageNum, pageNum+1, pageNum+2
	} else if pageNum == page {
		p1, p2, p3 = pageNum-2, pageNum-1, pageNum
	} else {
		p1, p2, p3 = pageNum-1, pageNum, pageNum+1
	}
	err = G(tmpl, "page", pageInfo{
		PageNum:  pageNum,
		Page:     page,
		P1:       p1,
		P2:       p2,
		P3:       p3,
		Articles: articles,
	})
	if err != nil {
		fmt.Println("Execute fail: " + path + "index.html")
		return
	}
	err = G(f, "index", pageIndexParams{
		Tmpl:        tmpl.String(),
		WebsiteName: websiteName,
		WebsiteURL:  websiteURL,
		HeadTitle:   websiteName,
	})
	if err != nil {
		fmt.Println("Execute fail: " + path + "index.html")
		return
	}
	fmt.Println("Execute: " + f.Name())
	f.Close()
}

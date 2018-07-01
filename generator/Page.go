package generator

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type pageInfo struct {
	PageNum  int
	Articles []ArticleInfo
}

type ArticleInfo struct {
	Title      string
	URL        string
	Date       string
	WriterName string
	Text       string
}

func GeneratePage(pageNum int, articles []ArticleInfo) {
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
	err = G(tmpl, "page", pageInfo{
		PageNum:  pageNum,
		Articles: articles,
	})
	if err != nil {
		fmt.Println("Execute fail: " + path + "index.html")
		return
	}
	err = G(f, "index", pageIndexParams{
		Tmpl:  tmpl.String(),
		Title: "XMatrix",
	})
	if err != nil {
		fmt.Println("Execute fail: " + path + "index.html")
		return
	}
	fmt.Println("Execute: " + f.Name())
	f.Close()
}

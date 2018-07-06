package generator

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func GenerateArchive(websiteName, websiteURL string, websiteArticles, pageNum int, articles []ArticleInfo) {
	var path string
	if pageNum == 1 {
		path = "dist/archives/"
	} else if pageNum > 1 {
		path = "dist/archives/page/" + strconv.Itoa(pageNum) + "/"
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
	if pageNum == 1 {
		p1, p2, p3 = pageNum, pageNum+1, pageNum+2
	} else if pageNum == (websiteArticles+9)/10 {
		p1, p2, p3 = pageNum-2, pageNum-1, pageNum
	} else {
		p1, p2, p3 = pageNum-1, pageNum, pageNum+1
	}
	err = G(tmpl, "archive", pageInfo{
		PageNum:  pageNum,
		Page:     (websiteArticles + 9) / 10,
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

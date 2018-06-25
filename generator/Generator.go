package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var t *template.Template
var arr []string

func Generate(theme string) {
	t = template.New("index")
	filepath.Walk("themes/"+theme, walkingLayout)
	for _, item := range arr {
		err := t.ExecuteTemplate(os.Stdout, item, nil)
		if err != nil {
			fmt.Println("Execute fail: " + item)
			continue
		}
		fmt.Println("Execute success: " + item)
	}
	fmt.Println(t.DefinedTemplates())
	return
}

func walkingLayout(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	} else if f.IsDir() {
		return nil
	}
	t, err = t.ParseFiles(path)
	if err != nil {
		fmt.Println("Layout fail: " + path)
		return err
	}
	fmt.Println("Layout loading: " + path)
	if !strings.Contains(path, "layout") {
		arr = append(arr, f.Name())
	}
	return nil
}

func generateIndex(theme string) {

}

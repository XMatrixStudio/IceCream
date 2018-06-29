package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var t *template.Template

type pageIndexParams struct {
	Tmpl string
}

func Generate(theme string) {
	t = template.New("index")
	filepath.Walk("themes/default/layouts", walkingLayouts)
	pages, err := walkingPages("themes/default/pages", "")
	if err != nil {
		fmt.Println("Read dir fail: " + "themes/default/pages")
		return
	}
	for tPath, tName := range pages {
		tmpl := new(bytes.Buffer)
		err := t.ExecuteTemplate(tmpl, tName, nil)
		if err != nil {
			fmt.Println("Execute template fail: " + tPath)
			continue
		}
		path := "dist" + tPath
		if _, err := os.Stat(path); err != nil {
			err := os.MkdirAll(path, 0777)
			if err != nil {
				fmt.Println("Create file fail: " + path)
			}
		}
		f, err := os.Create("dist" + tPath + "/index.html")
		if err != nil {
			fmt.Println("Create file fail: " + tName)
			continue
		}
		err = t.ExecuteTemplate(f, "index", pageIndexParams{
			Tmpl: tmpl.String(),
		})
		if err != nil {
			fmt.Println("Execute fail: " + tName)
			fmt.Println(err.Error())
			continue
		}
		f.Close()
		fmt.Println("Execute: " + f.Name())
	}
}

func walkingLayouts(path string, f os.FileInfo, err error) error {
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
	return nil
}

func walkingPages(root, dirPath string) (pages map[string]string, err error) {
	dir, err := ioutil.ReadDir(root + dirPath)
	if err != nil {
		fmt.Println("Read dir fail: " + root + dirPath)
		return
	}
	sep := string(os.PathSeparator)
	for _, f := range dir {
		if f.IsDir() {
			walkingPages(root, dirPath+sep+f.Name())
		} else {
			ok := strings.HasSuffix(f.Name(), ".html")
			if ok {
				t, err = t.ParseFiles(root + dirPath + sep + f.Name())
				if err != nil {
					fmt.Println("Layout page fail: " + root + dirPath + sep + f.Name())
					return
				}
				fmt.Println("Layout page loading: " + root + dirPath + sep + f.Name())
				name := f.Name()[:len(f.Name())-5]
				if pages == nil {
					pages = make(map[string]string)
				}
				pages[dirPath+sep+name] = name
			}
		}
	}
	return
}

func generateIndex(theme string) {

}

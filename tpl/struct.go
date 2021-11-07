package tpl

import (
	"io/ioutil"
	"log"
	"net/url"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type TemplateData struct {
	SiteUrl     string `remark:"站点地址"`
	Title       string `remark:"<title>Title"`
	Keywords    string `remark:"<meta>Keywords逗号隔开"`
	Description string `remark:"<meta>description"`
	HasKatex    bool   `remark:"md中是否解析了katex"`
	HasMermaid  bool   `remark:"md中是否解析了mermaid"`

	CurrentDirs   []TemplateFileItemData `remark:"路径"`
	CurrentFile   string                 `remark:"当前渲染文件(也有可能是目录)"`
	CurrentIsFile bool                   `remark:"是否有渲染文件"`
	Content       []byte                 `remark:"md"`
	MdHtml        string                 `remark:"html"`
	UpperPath     string                 `remark:"上一级连接"`
}

type TemplateFileItemData struct {
	FileExtension string `remark:"文件后缀名"`
	IsFile        bool   `remark:"是否文件"`
	Name          string `remark:"文件或目录名"`
	Href          string `remark:"连接，不带SiteUrl"`
	Icon          string `remark:"Icon"`
}

func GetTemplateData(pathname string, pathnameIsDir bool, mdfolder string, indexs string) (data *TemplateData) {
	data = &TemplateData{
		CurrentDirs: []TemplateFileItemData{},
	}
	pathname, _ = url.QueryUnescape(pathname)
	if pathname == "" {
		pathname = "/"
	}
	_path := strings.ReplaceAll(filepath.Dir(pathname), "\\", "/")
	if _path == "" {
		_path = "/"
	}
	if _path != "/" {
		data.UpperPath = strings.ReplaceAll(filepath.Dir(_path), "\\", "/")
		if !strings.HasSuffix(data.UpperPath, "/") {
			data.UpperPath = data.UpperPath + "/"
		}
	}
	index_array := strings.Split(indexs, ",")
	if indexs == "" {
		index_array = []string{"index.html", "index.md", "README.md"}
	}
	fileFullName := path.Join(mdfolder, pathname)
	listDir := fileFullName
	data.CurrentFile = pathname
	data.CurrentIsFile = !pathnameIsDir

	if !pathnameIsDir {
		listDir = filepath.Dir(fileFullName)
	}
	for _, fi := range Listdir(listDir) {
		if strings.HasPrefix(fi.Name(), ".") || strings.HasPrefix(fi.Name(), "_") {
			continue
		}
		item := TemplateFileItemData{
			Name:          fi.Name(),
			IsFile:        !fi.IsDir(),
			FileExtension: "",
			Href:          strings.ReplaceAll(path.Join(_path, fi.Name()), "\\", "/"),
		}
		if !fi.IsDir() {
			item.FileExtension = path.Ext(fi.Name())
			if pathnameIsDir {
				for _, index := range index_array {
					if index == fi.Name() {
						data.CurrentFile = item.Href
						data.CurrentIsFile = true
					}
				}
			}
		} else {
			item.Href = item.Href + "/"
		}
		if item.FileExtension == "" {
			item.Icon = GetExtensionsIcon(item.Name, fi.IsDir())
		} else {
			item.Icon = GetExtensionsIcon(item.FileExtension, fi.IsDir())
		}
		data.CurrentDirs = append(data.CurrentDirs, item)
	}
	sort.Slice(data.CurrentDirs, func(i, j int) bool {
		// 1. IsFile:升序排序
		if data.CurrentDirs[i].IsFile != data.CurrentDirs[j].IsFile {
			return data.CurrentDirs[j].IsFile
		}
		// 2. Name:升序排序
		return strings.ToLower(data.CurrentDirs[i].Name) < strings.ToLower(data.CurrentDirs[j].Name)
	})
	if data.CurrentIsFile {
		c, err := ioutil.ReadFile(path.Join(mdfolder, data.CurrentFile))
		if err != nil {
			log.Println(err)
			return nil
		}
		data.Content = c
	}
	return
}

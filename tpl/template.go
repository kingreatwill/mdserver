package tpl

import (
	"io"
	"log"
	"text/template"

	"github.com/kingreatwill/README/static"
)

/*
// bytes.Buffer 实现了io.Writer接口
buf := new(bytes.Buffer)
tpl.Execute(buf, map[string]interface{}{
	"Title": "Bob",
})
fmt.Println(buf.String())
*/
func Execute(wr io.Writer, data *TemplateData) {
	// 首先查找static目录；存在就读取对应的模板文件；
	t, err := template.ParseFS(static.TemplatesFiles, "templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpls := Listfile("static/templates")
	if len(tmpls) > 0 {
		t, err = t.ParseFiles(tmpls...)
		if err != nil {
			log.Println(err)
			return
		}
	}
	err = t.Execute(wr, data)
	if err != nil {
		log.Println(err)
		return
	}
}

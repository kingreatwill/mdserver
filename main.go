package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/kingreatwill/README/md"
	"github.com/kingreatwill/README/static"
	"github.com/kingreatwill/README/tpl"
	"github.com/valyala/fasthttp"
)

func main() {
	mdfolder := "E:/code/open"
	m := func(ctx *fasthttp.RequestCtx) {
		_path := string(ctx.Path())
		if strings.HasPrefix(_path, "/static/") {
			send_file(ctx, strings.Replace(_path, "/static/", "", 1), true)
		} else {
			fileFullName := path.Join(mdfolder, _path)
			c_fi, _ := os.Stat(fileFullName)
			if c_fi == nil {
				ctx.Error("file not found.", fasthttp.StatusNotFound)
				return
			}
			if c_fi.IsDir() || strings.HasSuffix(_path, ".md") {
				data := tpl.GetTemplateData(_path, c_fi.IsDir(), mdfolder, "")
				if data == nil {
					ctx.Error("err.", fasthttp.StatusInternalServerError)
					return
				}
				if len(data.Content) > 0 {
					data.MdHtml = string(md.New().Convert(data.Content))
					//data.MdHtml = template.HTML(data.MdHtml)
				}
				body := new(bytes.Buffer)
				tpl.Execute(body, data)
				ctx.SetStatusCode(fasthttp.StatusOK)
				ctx.Response.Header.Set("Content-type", "text/html")
				ctx.SetBody(body.Bytes())
			} else {
				send_file(ctx, fileFullName, false)
			}
		}
	}
	fasthttp.ListenAndServe(":8080", m)
}

func send_file(ctx *fasthttp.RequestCtx, filename string, embed bool) {
	ext := path.Ext(filename)
	var buf []byte
	var err error
	if embed {
		buf, err = static.IconsCssFiles.ReadFile(filename)
	} else {
		buf, err = ioutil.ReadFile(filename)
	}
	if err != nil {
		ctx.Error("file not found.", fasthttp.StatusNotFound)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.Set("Content-type", tpl.GetMime(ext))
	ctx.SetBody(buf)
}

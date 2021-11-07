package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/kingreatwill/README/md"
	"github.com/kingreatwill/README/static"
	"github.com/kingreatwill/README/tpl"
	"github.com/valyala/fasthttp"
)

var (
	directory string
	host      string
	port      int
	crontab   string
	index     string
)

func init() {
	flag.StringVar(&directory, "d", "", "default:current directory.") // -d
	flag.StringVar(&host, "host", "", "host.")
	flag.StringVar(&crontab, "c", "0 *1 * * *", "默认整点执行一次.")
	flag.StringVar(&index, "i", "index.html,index.md,README.md", "指定默认主页.")
	flag.IntVar(&port, "p", 8080, "端口")
}

// 获取当前执行文件绝对路径（go run）
func getPwd() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func main() {
	// 参数解析;
	flag.Parse()
	if directory == "" {
		directory = getPwd()
	}
	addr := fmt.Sprintf("%v:%v", host, port)
	// 路径处理;
	m := func(ctx *fasthttp.RequestCtx) {
		_path := string(ctx.Path())
		if strings.HasPrefix(_path, "/static/") {
			send_file(ctx, strings.Replace(_path, "/static/", "", 1), true)
		} else {
			fileFullName := path.Join(directory, _path)
			c_fi, _ := os.Stat(fileFullName)
			if c_fi == nil {
				ctx.Error("file not found.", fasthttp.StatusNotFound)
				return
			}
			if c_fi.IsDir() || strings.HasSuffix(_path, ".md") {
				data := tpl.GetTemplateData(_path, c_fi.IsDir(), directory, "")
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
	fasthttp.ListenAndServe(addr, m)
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

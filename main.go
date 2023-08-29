package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

// todo: 增加在线人数, 文件访问次数, 显示创建和修改时间, 显示git提交信息和diff, 排除文件夹, 隐藏文件等功能
// https://github.com/go-git/go-git
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
	handler := func(ctx *fasthttp.RequestCtx) {
		_path := string(ctx.Path())
		if strings.HasPrefix(_path, "/static/") {
			send_file(ctx, strings.Replace(_path, "/static/", "", 1), true)
		} else {
			fileFullName := path.Join(directory, _path)
			c_fi, _ := os.Stat(fileFullName)
			if c_fi == nil {
				ctx.Error("file or directory not found.", fasthttp.StatusNotFound)
				return
			}
			if c_fi.Name()[0] == '.' {
				ctx.Error("Warning, don't do anything illegal.", fasthttp.StatusInternalServerError)
				return
			}
			if c_fi.IsDir() || strings.HasSuffix(_path, ".md") {
				data := tpl.GetTemplateData(_path, c_fi.IsDir(), directory, "")
				if data == nil {
					ctx.Error("err.", fasthttp.StatusInternalServerError)
					return
				}
				if len(data.Content) > 0 {
					md.New().Convert(data)
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
	log.Println(fmt.Sprintf("Serving from http://%v", addr))
	s := &fasthttp.Server{
		Handler:        handler,
		ReadBufferSize: 1024 * 1024, // 本地开发，有时会出现BIG cookies；
	}
	if err := s.ListenAndServe(addr); err != nil {
		log.Println(err)
	}
	//fasthttp.ListenAndServe(addr, handler)
}

func send_file(ctx *fasthttp.RequestCtx, filename string, embed bool) {
	ext := path.Ext(filename)
	var buf []byte
	var err error
	if embed {
		if buf, err = ioutil.ReadFile(filename); err != nil {
			buf, err = static.Files.ReadFile(filename)
		}
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

package md

import (
	"bytes"
	"fmt"
	"log"

	mermaid "github.com/abhinav/goldmark-mermaid"
	toc "github.com/abhinav/goldmark-toc"
	"github.com/kingreatwill/README/tpl"
	katex "github.com/kingreatwill/goldmark-katex"

	//mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type mdConvert struct {
	engine goldmark.Markdown
}

func New() *mdConvert {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			katex.KaTeX,
			emoji.Emoji,
			//mathjax.MathJax,
			highlighting.Highlighting,
			&toc.Extender{},
			&mermaid.Extender{},
			meta.Meta,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	return &mdConvert{
		engine: md,
	}
}

func (c *mdConvert) Convert(data *tpl.TemplateData) {
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := c.engine.Convert(data.Content, &buf, parser.WithContext(context)); err != nil {
		log.Println(err)
		return
	}

	data.MdHtml = buf.String()

	metaData := meta.Get(context)
	if value, ok := metaData["Title"]; ok {
		data.Title = fmt.Sprintf("%v", value)
	} else if value, ok := metaData["title"]; ok {
		data.Title = fmt.Sprintf("%v", value)
	}

	if value, ok := metaData["Keywords"]; ok {
		data.Keywords = fmt.Sprintf("%v", value)
	} else if value, ok := metaData["keywords"]; ok {
		data.Keywords = fmt.Sprintf("%v", value)
	} else if value, ok := metaData["Tags"]; ok {
		data.Keywords = fmt.Sprintf("%v", value)
	} else if value, ok := metaData["tags"]; ok {
		data.Keywords = fmt.Sprintf("%v", value)
	}
	if value, ok := metaData["Description"]; ok {
		data.Description = fmt.Sprintf("%v", value)
	} else if value, ok := metaData["description"]; ok {
		data.Description = fmt.Sprintf("%v", value)
	} else if value, ok := metaData["Summary"]; ok {
		data.Description = fmt.Sprintf("%v", value)
	} else if value, ok := metaData["summary"]; ok {
		data.Description = fmt.Sprintf("%v", value)
	}

}

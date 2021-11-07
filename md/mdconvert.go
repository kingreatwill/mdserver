package md

import (
	"bytes"
	"log"

	mermaid "github.com/abhinav/goldmark-mermaid"
	toc "github.com/abhinav/goldmark-toc"
	//mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting"
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
			emoji.Emoji,
			//mathjax.MathJax,
			highlighting.Highlighting,
			&toc.Extender{},
			&mermaid.Extender{},
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

func (c *mdConvert) Convert(input []byte) (output []byte) {
	var buf bytes.Buffer
	if err := c.engine.Convert(input, &buf); err != nil {
		log.Println(err)
		return
	}
	output = buf.Bytes()
	return
}

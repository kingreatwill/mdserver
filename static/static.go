package static

import (
	"embed"
)

// https://zhuanlan.zhihu.com/p/351931501
// https://www.cnblogs.com/apocelipes/p/13907858.html

// //go:embed index.html
// var IndexTemplate []byte

// //go:embed header.html
// var HeaderTemplate []byte

// //go:embed footer.html
// var FooterTemplate []byte

// //go:embed *
// var StaticFiles embed.FS

//go:embed templates
var TemplatesFiles embed.FS

// //go:embed css
// var CssFiles embed.FS

//go:embed icons css
var IconsCssFiles embed.FS

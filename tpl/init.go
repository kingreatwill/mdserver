package tpl

import "mime"

func init() {
	mime.AddExtensionType("", "application/octet-stream")
	mime.AddExtensionType(".py", "text/plain")
	mime.AddExtensionType(".c", "text/plain")
	mime.AddExtensionType(".h", "text/plain")
}

func GetMime(ext string) string {
	mineType := mime.TypeByExtension(ext)
	if mineType == "" {
		return "application/octet-stream"
	}
	return mineType
}

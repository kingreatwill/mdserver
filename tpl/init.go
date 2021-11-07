package tpl

import (
	"log"
	"mime"
	"strings"
	"sync"

	"github.com/kingreatwill/README/static"
)

var extensions_icon sync.Map // map[string][]string; slice values are append-only.
func init() {
	mime.AddExtensionType("", "application/octet-stream")
	mime.AddExtensionType(".py", "text/plain")
	mime.AddExtensionType(".c", "text/plain")
	mime.AddExtensionType(".h", "text/plain")

	Extensions_Icon_Init()
}

func GetMime(ext string) string {
	mineType := mime.TypeByExtension(ext)
	if mineType == "" {
		return "application/octet-stream"
	}
	return mineType
}

/*
def extensions_icon_map_init():
    icon_map = {}
    rel_path = os.path.join(os.path.dirname(__file__), 'icons')
    for entry in os.listdir(rel_path):
        if entry.startswith('file_type_') or entry.startswith('folder_type_'):
            ext = '{}'.format(entry.replace('.svg', ''))
            icon_map[ext] = entry
    return icon_map

*/
func Extensions_Icon_Init() {
	dirs, err := static.Files.ReadDir("icons")
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range dirs {
		if v.IsDir() {
			continue
		}
		key := strings.Replace(v.Name(), ".svg", "", 1)
		if strings.HasPrefix(v.Name(), "file_type_") {
			key = strings.Replace(v.Name(), "file_type_", "", 1)
		} else if strings.HasPrefix(v.Name(), "folder_type_") {
			key = strings.Replace(v.Name(), "folder_type_", "", 1)
		}
		extensions_icon.Store("", "")
	}
}

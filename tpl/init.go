package tpl

import (
	"fmt"
	"log"
	"mime"
	"strings"
	"sync"

	"github.com/kingreatwill/README/static"
)

var (
	extensions_file_icon   sync.Map // map[string][]string; slice values are append-only.
	extensions_folder_icon sync.Map
)

func init() {
	mime.AddExtensionType("", "application/octet-stream")
	mime.AddExtensionType(".py", "text/plain")
	mime.AddExtensionType(".c", "text/plain")
	mime.AddExtensionType(".h", "text/plain")
	mime.AddExtensionType(".go", "text/plain")
	mime.AddExtensionType(".cs", "text/plain")
	mime.AddExtensionType(".mod", "text/plain")
	mime.AddExtensionType(".sum", "text/plain")

	extensions_Icon_Init()
	file_icon_update := map[string][]string{
		"file_type_markdown.svg":   {".md"},
		"file_type_bundler.svg":    {"gemfile"},
		"file_type_zip.svg":        {".gz", ".7z", ".tar", ".tgz", ".bz"},
		"file_type_go_package.svg": {".mod", ".sum"},
		"file_type_docker2.svg":    {"dockerfile"},
		"file_type_image.svg":      {".jpeg", ".jpg", ".gif", ".png", ".bmp", ".tiff", ".ico"},
	}
	folder_icon_update := map[string][]string{
		"folder_type_windows.svg": {"win"},
		"folder_type_tests.svg":   {"test", "integration", "specs", "spec"},
		"folder_type_images.svg":  {"img", "image", "imgs"},
		"folder_type_src.svg":     {"source", "sources"},
		"folder_type_log.svg":     {"logs"},
		"folder_type_locale.svg":  {"lang", "language", "languages", "locales", "internationalization", "i18n", "globalization", "g11n", "localization", "l10n"},
	}
	for icon, v := range file_icon_update {
		for _, key := range v {
			extensions_file_icon.Store(key, icon)
		}
	}
	for icon, v := range folder_icon_update {
		for _, key := range v {
			extensions_folder_icon.Store(key, icon)
		}
	}

}

func GetMime(ext string) string {
	mineType := mime.TypeByExtension(ext)
	if mineType == "" {
		return "application/octet-stream"
	}
	return mineType
}

func GetExtensionsIcon(ext string, isdir bool) string {
	ext = strings.ToLower(ext)
	if isdir {
		if v, ok := extensions_folder_icon.Load(ext); ok {
			return fmt.Sprint(v)
		}
		return "default_folder.svg"
	}
	if v, ok := extensions_file_icon.Load(ext); ok {
		return fmt.Sprint(v)
	}
	return "default_file.svg"
}

func extensions_Icon_Init() {
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
			key = strings.Replace(key, "file_type_", "", 1)
			extensions_file_icon.Store("."+key, v.Name())
			extensions_file_icon.Store(key, v.Name())
		} else if strings.HasPrefix(v.Name(), "folder_type_") {
			key = strings.Replace(key, "folder_type_", "", 1)
			extensions_folder_icon.Store(key, v.Name())
		}
	}
}

package tpl

import (
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

	Extensions_Icon_Init()
	for _, k := range []string{".md"} {
		extensions_file_icon.Store(k, "file_type_markdown.svg")
	}
	for _, k := range []string{"gemfile"} {
		extensions_file_icon.Store(k, "file_type_bundler.svg")
	}
	for _, k := range []string{".gz", ".7z", ".tar", ".tgz", ".bz"} {
		extensions_file_icon.Store(k, "file_type_zip.svg")
	}
	for _, k := range []string{".mod", ".sum", ".tar", ".tgz", ".bz"} {
		extensions_file_icon.Store(k, "file_type_go_package.svg")
	}
	for _, k := range []string{"dockerfile"} {
		extensions_file_icon.Store(k, "file_type_docker2.svg")
	}
}

/*
extensions_icon_map.update({


        'file_type_jpeg': 'file_type_image.svg',
        'file_type_jpg': 'file_type_image.svg',
        'file_type_gif': 'file_type_image.svg',
        'file_type_png': 'file_type_image.svg',
        'file_type_bmp': 'file_type_image.svg',
        'file_type_tiff': 'file_type_image.svg',
        'file_type_ico': 'file_type_image.svg',

        'folder_type_lang': 'folder_type_locale.svg',
        'folder_type_language': 'folder_type_locale.svg',
        'folder_type_languages': 'folder_type_locale.svg',
        'folder_type_locales': 'folder_type_locale.svg',
        'folder_type_internationalization': 'folder_type_locale.svg',
        'folder_type_i18n': 'folder_type_locale.svg',
        'folder_type_globalization': 'folder_type_locale.svg',
        'folder_type_g11n': 'folder_type_locale.svg',
        'folder_type_localization': 'folder_type_locale.svg',
        'folder_type_l10n': 'folder_type_locale.svg',
        'folder_type_logs': 'folder_type_log.svg',
        'folder_type_img': 'folder_type_images.svg',
        'folder_type_image': 'folder_type_images.svg',
        'folder_type_imgs': 'folder_type_images.svg',
        'folder_type_source': 'folder_type_src.svg',
        'folder_type_sources': 'folder_type_src.svg',

        'folder_type_tests': 'folder_type_test.svg',
        'folder_type_integration': 'folder_type_test.svg',
        'folder_type_specs': 'folder_type_test.svg',
        'folder_type_spec': 'folder_type_test.svg',
        'folder_type_win': 'folder_type_windows.svg',
    })
*/

func GetMime(ext string) string {
	mineType := mime.TypeByExtension(ext)
	if mineType == "" {
		return "application/octet-stream"
	}
	return mineType
}

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
			key = strings.Replace(key, "file_type_", "", 1)
			extensions_file_icon.Store("."+key, v.Name())
			extensions_file_icon.Store(key, v.Name())
		} else if strings.HasPrefix(v.Name(), "folder_type_") {
			key = strings.Replace(key, "folder_type_", "", 1)
			extensions_folder_icon.Store(key, v.Name())
		}
	}
}

package tpl

import (
	"io/fs"
	"log"
	"os"
	"path"
)

func Listfile(pathname string) (files []string) {
	for _, fi := range Listdir(pathname) {
		if !fi.IsDir() {
			fullName := path.Join(pathname, fi.Name())
			files = append(files, fullName)
		}
	}
	return
}

func Listdir(pathname string) (fileInfos []fs.FileInfo) {
	dirEntries, err := os.ReadDir(pathname)
	// rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		log.Println(err)
		return nil
	}
	for _, de := range dirEntries {
		fi, _ := de.Info()
		fileInfos = append(fileInfos, fi)
	}
	return
}

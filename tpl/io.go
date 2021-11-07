package tpl

import (
	"io/fs"
	"io/ioutil"
	"log"
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

func Listdir(pathname string) []fs.FileInfo {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		log.Println(err)
		return nil
	}
	return rd
}

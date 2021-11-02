package memdata

import (
	"fmt"
	"os"
)

func check(e error) {

	if e != nil {
		panic(e)
	}

}

type MemPath struct {
	Data []byte
	Path string
}

type PathMap struct {
	Pathmap map[string]*MemPath
}

func NewPathMap() *PathMap {

	pathmap := new(PathMap)
	pathmap.Pathmap = make(map[string]*MemPath)
	return pathmap

}

func (pathmap PathMap) GetMemPath(path string) *MemPath {
	return pathmap.Pathmap[path]
}

func (pathmap PathMap) SetMemPath(path string) *MemPath {
	p := new(MemPath)
	p.Path = path
	pathmap.Pathmap[p.Path] = p
	return p
}

func NewPath(path string, mapper *PathMap) *MemPath {

	newpath := mapper.SetMemPath(path)
	return newpath

}

func (path MemPath) ShowPath() {

	fmt.Println(path.Path)

}

func (path MemPath) Backup() {

	fd, err := os.Create(path.Path)
	check(err)
	defer fd.Close()

	fd.Write(path.Data)

}

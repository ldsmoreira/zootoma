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
	Data Data
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

func NewPath(path string, mapper *PathMap) *MemPath {
	p := new(MemPath)
	p.Path = path
	mapper.Pathmap[p.Path] = p
	return p

}

func (path MemPath) ShowPath() {
	fmt.Println(path.Path)
}

func (path MemPath) Backup() {
	fd, err := os.Create(path.Path)
	check(err)
	defer fd.Close()

	fd.Write(path.Data.Content)
}

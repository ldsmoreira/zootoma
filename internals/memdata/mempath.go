package memdata

import (
	"fmt"
	"os"
)

//check funtion checks if there are errors
func check(e error) {

	if e != nil {
		panic(e)
	}

}

//MemPath is the basic data element of tomazoo
type MemPath struct {
	Data []byte //Data persisted in memory
	Path string //Path string with the path to the persisted data
}

//PathMap is the basic data structure to manage MemPath objects
type PathMap struct {
	Pathmap map[string]*MemPath //MemPath is the Hash Map used to find data in O(1)
}

//NewPathMap is the constructor of the PathMap objects
func NewPathMap() *PathMap {

	pathmap := new(PathMap)
	pathmap.Pathmap = make(map[string]*MemPath)
	return pathmap

}

//GetMemPath is a method of PathMap that return the MemPath object
//associated with the path string
func (pathmap PathMap) GetMemPath(path string) *MemPath {
	return pathmap.Pathmap[path]
}

//SetPathMap is a method that set a key of path inside PathMap
func (pathmap PathMap) SetMemPath(path string) *MemPath {
	p := new(MemPath)
	p.Path = path
	pathmap.Pathmap[p.Path] = p
	return p
}

//NewPath is the constructor of MemPath
func NewPath(path string, mapper *PathMap) *MemPath {

	newpath := mapper.SetMemPath(path)
	return newpath

}

//ShowPath shows the path associated with MemPath object
func (path MemPath) ShowPath() {

	fmt.Println(path.Path)

}

func (path MemPath) Backup() {

	fd, err := os.Create(path.Path)
	check(err)
	defer fd.Close()

	fd.Write(path.Data)

}

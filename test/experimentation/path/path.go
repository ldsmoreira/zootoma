package path

import (
	"fmt"
	"os"
	"tomabase/test/experimentation/datacarrier"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Path struct {
	Data   datacarrier.Data
	Way    string
	Child  []*Path
	Parent *Path
}

type PathMap struct {
	Pathmap map[string]*Path
}

func NewPathMap() *PathMap {

	pathmap := new(PathMap)
	pathmap.Pathmap = make(map[string]*Path)
	return pathmap
}

func NewPath(path string, mapper *PathMap) *Path {
	p := new(Path)
	p.Way = path
	mapper.Pathmap[p.Way] = p
	return p

}

func (path Path) ShowPath() {
	fmt.Println(path.Way)
}

func (path Path) Backup() {
	fd, err := os.Create(path.Way)
	check(err)
	defer fd.Close()

	fd.Write(path.Data.Content)
}

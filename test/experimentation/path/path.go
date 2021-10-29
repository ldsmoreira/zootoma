package path

import (
	"fmt"
	"tomabase/test/experimentation/datacarrier"
)

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

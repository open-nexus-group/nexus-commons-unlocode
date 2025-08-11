//go:generate ../generator/generator.go -output=../unlocode/data.go SubDiv.csv UNLOC-1.csv UNLOC-2.csv UNLOC-3.csv UNLOC-4.csv
package subdivision 

import (
	"log"
	"gitlab.com/open-agent-nexus/nexus-commons-unlocode/pkg/data"
)


type SubDivCollection struct {
	subdivisions []data.SubDivision
}

func (s *SubDivCollection) GetByName() {
	log.Println(s.subdivisions)
}

func NewSubDivisionCollection() *SubDivCollection {
	collection := SubDivCollection{
		subdivisions: data.Subdivisions,
	}
	return &collection
}

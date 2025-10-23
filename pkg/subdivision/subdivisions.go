//go:generate ../../generator/generator.go --type=subdivision --files=2024-2-SubDiv.csv
package subdivision

import (
	"github.com/open-nexus-group/nexus-commons-unlocode/pkg/data"
)

type SubDivCollection struct {
	subdivisions       []data.Subdivision
	subdivisionCodeMap map[string]data.Subdivision
}

func CreateSubdivisionMapByCode(subdivisions []data.Subdivision) map[string]data.Subdivision {
	subdivisionMap := make(map[string]data.Subdivision, len(subdivisions)) // Pre-allocate map capacity
	for _, sd := range subdivisions {
		subdivisionMap[sd.Code] = sd
	}
	return subdivisionMap
}

func (collection *SubDivCollection) GetByCode(code string) data.Subdivision {
	return collection.subdivisionCodeMap[code]
}

func NewSubDivisionCollection() *SubDivCollection {
	collection := SubDivCollection{
		subdivisions:       data.Subdivisions,
		subdivisionCodeMap: CreateSubdivisionMapByCode(data.Subdivisions),
	}
	return &collection
}

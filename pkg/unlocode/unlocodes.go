//go:generate ../generator/generator.go -output=../unlocode/data.go SubDiv.csv UNLOC-1.csv UNLOC-2.csv UNLOC-3.csv UNLOC-4.csv
package unlocode

import (
	"log"
	"gitlab.com/open-agent-nexus/nexus-commons-unlocode/pkg/data"
)

type UnLocodeCollection struct {
	unlocodes []data.Unlocode
}

func (collection *UnLocodeCollection) GetByName() {
	log.Println(collection.unlocodes)
}

func NewUnLocodeCollection() *UnLocodeCollection{
	collection := UnLocodeCollection{
		unlocodes: data.Unlocodes,
	}
	return &collection
}

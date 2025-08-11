//go:generate ../../generator/generator.go --type=unlocode --files=UNLOC-1.csv,UNLOC-2.csv,UNLOC-3.csv,UNLOC-4.csv
package unlocode

import (
	"errors"
	"github.com/hbollon/go-edlib"
	"gitlab.com/open-agent-nexus/nexus-commons-unlocode/pkg/data"
	"maps"
	"slices"
	"strings"
)

var (
	Err_Not_Found     = errors.New("could not find any results")
	Err_During_Lookup = errors.New("lookup failed - see logs")
)

type UnLocodeRepository interface {
	GetByLocode(code string) (result data.Unlocode, err error)
	GetByName(name string) (result data.Unlocode, err error)
	FindByName(name string, amount int) (result []data.Unlocode, err error)
	FindByLocode(code string, amount int) (result []data.Unlocode, err error)
}

type UnLocodeCollection struct {
	unlocodes  []data.Unlocode
	locodeMap  map[string]data.Unlocode
	locodeKeys []string
	nameMap    map[string]data.Unlocode
	nameKeys   []string
}

func CreateUnLocodeMapByLocode(unlocodes []data.Unlocode) (unlocodeMap map[string]data.Unlocode, keys []string) {
	unlocodeMap = make(map[string]data.Unlocode, len(unlocodes)) // Pre-allocate map capacity
	for _, u := range unlocodes {
		unlocodeMap[strings.ToLower(u.LoCode)] = u
	}
	keys = slices.Collect(maps.Keys(unlocodeMap))
	return unlocodeMap, keys
}

func CreateUnLocodeMapByName(unlocodes []data.Unlocode) (nameMap map[string]data.Unlocode, keys []string) {
	nameMap = make(map[string]data.Unlocode, len(unlocodes)) // Pre-allocate map capacity
	for _, u := range unlocodes {
		nameMap[strings.ToLower(strings.TrimSpace(u.NameWoDiacritics))] = u
	}
	keys = slices.Collect(maps.Keys(nameMap))
	return nameMap, keys
}

func NewUnLocodeCollection() UnLocodeRepository {
	unlocodeMap, unlocodeKeys := CreateUnLocodeMapByLocode(data.Unlocodes)
	nameMap, nameKeys := CreateUnLocodeMapByName(data.Unlocodes)
	collection := UnLocodeCollection{
		unlocodes:  data.Unlocodes,
		locodeMap:  unlocodeMap,
		locodeKeys: unlocodeKeys,
		nameMap:    nameMap,
		nameKeys:   nameKeys,
	}
	return &collection
}

func (collection *UnLocodeCollection) GetByLocode(code string) (result data.Unlocode, err error) {
	result = collection.locodeMap[code]
	if (result == data.Unlocode{}) {
		return data.Unlocode{}, Err_Not_Found
	}
	return result, nil
}

func (collection *UnLocodeCollection) GetByName(name string) (result data.Unlocode, err error) {
	result = collection.nameMap[name]
	if (result == data.Unlocode{}) {
		return data.Unlocode{}, Err_Not_Found
	}
	return result, nil
}

func (c *UnLocodeCollection) FindByLocode(code string, amount int) (results []data.Unlocode, err error) {
	res, err := edlib.FuzzySearchSet(code, c.locodeKeys, amount, edlib.Levenshtein)
	if err != nil {
		panic("Can not find by locode")
	}
	for _, v := range res {
		res, _ := c.GetByLocode(v)
		results = append(results, res)
	}
	return results, nil
}

func (c *UnLocodeCollection) FindByName(name string, amount int) (results []data.Unlocode, err error) {
	res, err := edlib.FuzzySearchSet(name, c.nameKeys, amount, edlib.Levenshtein)
	if err != nil {
		panic("Can not find by name")
	}
	for _, v := range res {
		res, _ := c.GetByName(v)
		results = append(results, res)
	}
	return results, nil
}

//go:generate ../../generator/generator.go --type=unlocode --files=UNLOC-1.csv,UNLOC-2.csv,UNLOC-3.csv,UNLOC-4.csv
package unlocode

import (
	"errors"
	"maps"
	"slices"
	"strings"

	"github.com/hbollon/go-edlib"
	"github.com/open-nexus-group/nexus-commons-unlocode/pkg/data"
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
	FindByNameAndProximity(name string, amount int, minSimilarity float32) (result []data.Unlocode, err error)
	FindByNameAndProximityWithFunction(name string, amount int, minSimilarity float32, function data.Function) (result []data.Unlocode, err error)
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

func (c *UnLocodeCollection) FindByNameAndProximity(name string, amount int, minSimilarity float32) (results []data.Unlocode, err error) {
	res, err := edlib.FuzzySearchSetThreshold(name, c.nameKeys, amount, minSimilarity, edlib.Levenshtein)
	if err != nil {
		panic("Can not find by name and proximity")
	}

	sortedNames := sortByProximity(name, res)

	for _, v := range sortedNames {
		for _, unlocode := range c.unlocodes {
			if strings.ToLower(strings.TrimSpace(unlocode.NameWoDiacritics)) == v {
				results = append(results, unlocode)
			}
		}
	}
	return results, nil
}

func (c *UnLocodeCollection) FindByNameAndProximityWithFunction(name string, amount int, minSimilarity float32, function data.Function) (results []data.Unlocode, err error) {
	res, err := edlib.FuzzySearchSetThreshold(name, c.nameKeys, amount, minSimilarity, edlib.Levenshtein)
	if err != nil {
		panic("Can not find by name and proximity with function")
	}

	sortedNames := sortByProximity(name, res)

	for _, v := range sortedNames {
		for _, unlocode := range c.unlocodes {
			if strings.ToLower(strings.TrimSpace(unlocode.NameWoDiacritics)) == v {
				if unlocode.Function != nil {
					for _, f := range *unlocode.Function {
						if f == function {
							results = append(results, unlocode)
							break
						}
					}
				}
			}
		}
	}
	return results, nil
}

/*
Sorts the matches by proximity to the search term.
The higher the score, the more similar the match is to the search term.
*/
func sortByProximity(searchTerm string, matches []string) []string {
	type scoredResult struct {
		name  string
		score float32
	}
	var scored []scoredResult
	for _, v := range matches {
		similarity, _ := edlib.StringsSimilarity(strings.ToLower(searchTerm), v, edlib.Levenshtein)
		scored = append(scored, scoredResult{name: v, score: similarity})
	}

	slices.SortFunc(scored, func(a, b scoredResult) int {
		if a.score > b.score {
			return -1
		}
		if a.score < b.score {
			return 1
		}
		return 0
	})

	result := make([]string, len(scored))
	for i, s := range scored {
		result[i] = s.name
	}
	return result
}

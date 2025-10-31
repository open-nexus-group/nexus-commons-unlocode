package main

import (
	"log"
	"time"

	"github.com/open-nexus-group/nexus-commons-unlocode/pkg/data"
	"github.com/open-nexus-group/nexus-commons-unlocode/pkg/unlocode"
)

func main() {
	locodeCollection := unlocode.NewUnLocodeCollection()
	start := time.Now()
	log.Printf("Starting lookups: %s", start)
	locationByCode, err := locodeCollection.GetByLocode("cnsha")
	if err != nil {
		log.Printf("Could not getByLocode: %v", err)
	}
	locationByName, err := locodeCollection.GetByName("shanghai hongqiao international apt")
	if err != nil {
		log.Printf("Could not getByName: %v", err)
	}
	foundByCode, err := locodeCollection.FindByLocode("cnsh", 15)
	if err != nil {
		log.Printf("Could not findByLocode: %v", err)
	}
	foundByName, err := locodeCollection.FindByName("shanghai", 10)
	if err != nil {
		log.Printf("Could not findByName: %v", err)
	}
	log.Printf("direct by Locode: CNSHA: %v", locationByCode)
	log.Printf("direct by Name: Shanghai Hongqiao International Apt: %v", locationByName)
	log.Printf("found by Locode: CNSH: %v", foundByCode)
	log.Printf("found by Name: shanghai: %v", foundByName)

	foundByNameAndProximity, err := locodeCollection.FindByNameAndProximity("Hamburg", 3, 0.8)
	if err != nil {
		log.Printf("Could not findByNameAndProximity: %v", err)
	}
	log.Printf("found by Name and Proximity: Hamburg: %v", foundByNameAndProximity)

	foundByNameAndProximityWithFunction, err := locodeCollection.FindByNameAndProximityWithFunction("Hambur", 5, 0.5, data.PORT)
	if err != nil {
		log.Printf("Could not findByNameAndProximityWithFunction: %v", err)
	}
	log.Printf("found by Name and Proximity with Function: Hamburg: %v", foundByNameAndProximityWithFunction)
	foundBySearchLikeWithFunction, err := locodeCollection.SearchLikeWithFunction("hamb", 5, data.PORT)
	if err != nil {
		log.Printf("Could not searchLikeWithFunction: %v", err)
	}
	log.Printf("count of found by Search Like with Function: %d", len(foundBySearchLikeWithFunction))
	log.Printf("found by Search Like with Function: Hamburg: %v", foundBySearchLikeWithFunction)
	elapsed := time.Since(start)
	log.Printf("All operations took %s", elapsed)
}

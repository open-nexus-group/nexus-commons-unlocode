package main

import (
	"log"
	"time"

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
	elapsed := time.Since(start)
	log.Printf("All operations took %s", elapsed)
}

package main

import (
	"gitlab.com/open-agent-nexus/nexus-commons-unlocode/pkg/unlocode"
)

func main() {
	SubdivisionCollection := unlocode.NewSubDivisionCollection() 
	SubdivisionCollection.PrintAll()
}

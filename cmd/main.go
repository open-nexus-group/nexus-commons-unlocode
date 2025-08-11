package main
import (
	"log"
	"gitlab.com/open-agent-nexus/nexus-commons-unlocode/pkg/unlocode"
)
func main() {
	log.Println("hello world")
	locodeCollection := unlocode.NewUnLocodeCollection()
	locodeCollection.GetByName()
}

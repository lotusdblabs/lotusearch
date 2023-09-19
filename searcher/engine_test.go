package searcher

import (
	"fmt"
	"os"
	"testing"
)

var (
	database = "testEngine"
	engine = &Engine{
		IndexPath:    fmt.Sprintf("%s%c%s", "/root/lotusearch/", os.PathSeparator, database),
	}
	// todo add test case
)

func TestGetDocumentCount(t *testing.T) {
	engine.GetDocumentCount()
}
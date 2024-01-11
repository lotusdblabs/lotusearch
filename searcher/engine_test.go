package searcher

import (
	"fmt"
	"github.com/sea-team/gofound/searcher/model"
	"os"
	"testing"
)

var (
	database = "testEngine"
	engine   = &Engine{
		IndexPath: fmt.Sprintf("%s%c%s", "/root/lotusearch/", os.PathSeparator, database),
	}
	// todo add test case

	doc = &model.IndexDoc{}

	worker = make(chan *model.IndexDoc, 1)
)

func TestGetDocumentCount(t *testing.T) {
	engine.GetDocumentCount()
}

func TestIndexDocument(t *testing.T) {
	engine.IndexDocument(doc)

	if engine.GetDocumentCount() != 1 {
		t.Error("IndexDocument test has err")
	}
}

func TestGetQueue(t *testing.T) {
	if engine.GetQueue() > 0 {
		t.Error("GetQueue test has err")
	}
}

func TestDocumentWorkerExec(t *testing.T) {
	engine.DocumentWorkerExec(worker)
}

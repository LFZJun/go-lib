package tag

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestStructTag_Parse(t *testing.T) {
	var tag StructTag = `json:"a" db:"n"`
	spew.Dump(tag.Parse())
}

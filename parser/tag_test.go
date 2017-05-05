package parser

import (
	"testing"
	"github.com/davecgh/go-spew/spew"
)

func TestStructTag_Parse(t *testing.T) {
	var tag StructTag = `json:"a" db:"n"`
	spew.Dump(tag.Parse())
}

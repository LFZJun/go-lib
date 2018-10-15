package csvx

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

type Names struct {
	FirstName string `head:"first_name"`
	LastName  string `head:"last_name"`
	UserName  string `head:"username"`
}

func TestNewRead(t *testing.T) {
	ast := assert.New(t)
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	reader, err := NewReader(r)
	ast.Nil(err)

	names := new(Names)
	err = reader.Read(names)
	ast.Nil(err)
	ast.Equal(Names{"Rob", "Pike", "rob"}, *names)
	err = reader.Read(names)
	ast.Nil(err)
	ast.Equal(Names{"Ken", "Thompson", "ken"}, *names)
	err = reader.Read(names)
	ast.Nil(err)
	ast.Equal(Names{"Robert", "Griesemer", "gri"}, *names)
	err = reader.Read(names)
	ast.Equal(io.EOF, err)
}

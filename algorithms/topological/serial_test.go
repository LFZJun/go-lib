package topological

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func root(f func(m *Member)) *Member {
	//      ->foo->foz-
	// root -         ->zz
	//      ->bar->baz-
	//      ->test

	root := NewMember("root", f)

	foo := NewMember("foo", f)
	bar := NewMember("bar", f)

	foz := NewMember("foz", f)
	baz := NewMember("baz", f)

	zz := NewMember("zz", f)

	root.Add(foo)
	root.Add(bar)

	foo.Add(foz)
	bar.Add(baz)

	foz.Add(zz)
	baz.Add(zz)
	return root
}

func TestMember_Do(t *testing.T) {
	var diagram []string

	r := root(func(m *Member) {
		diagram = append(diagram, m.Val.(string))
	})

	r.Do()

	assert.Equal(t, diagram, []string{"zz", "foz", "foo", "baz", "bar", "root"})
}

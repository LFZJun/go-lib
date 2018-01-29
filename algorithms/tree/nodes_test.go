package tree

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestNode_Size(t *testing.T) {
	n := new(Node)
	_ = n.Size()
}

func TestNodes_Add(t *testing.T) {
	nodes := new(Nodes)
	err := nodes.Add("/foo", 1, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	err = nodes.Add("/foz", 2, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	err = nodes.Add("/bar", 4, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	out, err := yaml.Marshal(nodes)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(string(out))
}

func TestNodes_Find(t *testing.T) {
	nodes := new(Nodes)
	err := nodes.Add("/foo", 1, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	err = nodes.Add("/foz", 1, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	err = nodes.Add("/bar", 1, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(nodes.Find("/foo"))
}

func TestNodes_FuzzyFind(t *testing.T) {
	nodes := new(Nodes)
	err := nodes.Add("/foo", 1, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	err = nodes.Add("/foz", 1, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	err = nodes.Add("/bar", 1, true)
	if err != nil {
		t.Fatal(err)
		return
	}
	ns := new(Nodes)
	nodes.FuzzyFind("/fo", ns)
	out, err := yaml.Marshal(ns)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(string(out))
}

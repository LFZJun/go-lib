package main

import (
	"fmt"
)

type Light struct {
	Site string
}

func (l *Light) On() {
	fmt.Println(l.Site + " light turn on")
}

func (l *Light) Off() {
	fmt.Println(l.Site + " light turn off")
}

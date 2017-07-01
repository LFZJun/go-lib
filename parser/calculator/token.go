package calculator

import (
	"fmt"
	"unicode"
)

type tokenType int

const (
	eof = -(iota + 1)
)

const (
	tokenError tokenType = iota
	tokenEOF
	tokenLeftParen
	tokenRightParen
	tokenInteger
	tokenFloat
	tokenPlus
	tokenMinus
	tokenStar
	tokenSlash
)

var tokenTypeNames = []string{
	"Error",
	"EOF",
	"(",
	")",
	"Integer",
	"Float",
	"+",
	"-",
	"*",
	"/",
}

func (tt tokenType) String() string {
	idx := int(tt)
	if idx < len(tokenTypeNames) {
		return tokenTypeNames[idx]
	}
	return "Unknown"
}

type token struct {
	Position
	typ tokenType
	val string
}

func (t token) String() string {
	switch t.typ {
	case tokenEOF:
		return "EOF"
	case tokenError:
		return t.val
	}

	return fmt.Sprintf("%q", t.val)
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isDigit(r rune) bool {
	return unicode.IsNumber(r)
}

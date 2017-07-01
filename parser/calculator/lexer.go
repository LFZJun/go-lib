package calculator

import (
	"bytes"
	"fmt"
	"github.com/pelletier/go-buffruneio"
	"io"
)

type (
	Lexer struct {
		input         *buffruneio.Reader // Textual source
		buffer        bytes.Buffer       // Runes composing the current token
		tokens        chan token
		depth         int
		line          int
		col           int
		endbufferLine int
		endbufferCol  int
	}

	LexStateFn func() LexStateFn
)

func (l *Lexer) read() rune {
	r, _, err := l.input.ReadRune()
	if err != nil {
		panic(err)
	}
	if r == '\n' {
		l.endbufferLine++
		l.endbufferCol = 1
	} else {
		l.endbufferCol++
	}
	return r
}

func (l *Lexer) next() rune {
	r := l.read()

	if r != eof {
		l.buffer.WriteRune(r)
	}
	return r
}

func (l *Lexer) ignore() {
	l.buffer.Reset()
	l.line = l.endbufferLine
	l.col = l.endbufferCol
}

func (l *Lexer) skip() {
	l.next()
	l.ignore()
}

func (l *Lexer) fastForward(n int) {
	for i := 0; i < n; i++ {
		l.next()
	}
}

func (l *Lexer) emitWithValue(t tokenType, value string) {
	l.tokens <- token{
		Position: Position{l.line, l.col},
		typ:      t,
		val:      value,
	}
	l.ignore()
}

func (l *Lexer) emit(t tokenType) {
	l.emitWithValue(t, l.buffer.String())
}

func (l *Lexer) peek() rune {
	r, _, err := l.input.ReadRune()
	if err != nil {
		panic(err)
	}
	l.input.UnreadRune()
	return r
}

func (l *Lexer) follow(next string) bool {
	for _, expectedRune := range next {
		r, _, err := l.input.ReadRune()
		defer l.input.UnreadRune()
		if err != nil {
			panic(err)
		}
		if expectedRune != r {
			return false
		}
	}
	return true
}

func (l *Lexer) errorf(format string, args ...interface{}) LexStateFn {
	l.tokens <- token{
		Position: Position{l.line, l.col},
		typ:      tokenError,
		val:      fmt.Sprintf(format, args...),
	}
	return nil
}

func (l *Lexer) lexVoid() LexStateFn {
	for {
		next := l.peek()
		switch next {
		case '(':
			return l.lexLeftParen
		case ')':
			return l.lexRightParen
		case '+':
			return l.lexPlus
		case '-':
			return l.lexMinus
		case '*':
			return l.lexStar
		case '/':
			return l.lexSlash
		case ' ', '\t', '\r', '\n':
			l.skip()
			continue
		case eof:
			return l.lexEOF
		default:
			return l.lexNumber
		}
	}
}

func (l *Lexer) lexNumber() LexStateFn {
	r := l.peek()
	if r == '+' || r == '-' {
		l.next()
	}
	pointSeen := false
	expSeen := false
	digitSeen := false
	for {
		next := l.peek()
		if next == '.' {
			if pointSeen {
				return l.errorf("cannot have two dots in one float")
			}
			l.next()
			if !isDigit(l.peek()) {
				return l.errorf("float cannot end with a dot")
			}
			pointSeen = true
		} else if next == 'e' || next == 'E' {
			expSeen = true
			l.next()
			r := l.peek()
			if r == '+' || r == '-' {
				l.next()
			}
		} else if isDigit(next) {
			digitSeen = true
			l.next()
		} else if next == '_' {
			l.next()
		} else {
			break
		}
		if pointSeen && !digitSeen {
			return l.errorf("cannot start float with a dot")
		}
	}

	if !digitSeen {
		return l.errorf("no digit in that number")
	}
	if pointSeen || expSeen {
		l.emit(tokenFloat)
	} else {
		l.emit(tokenInteger)
	}
	return l.lexVoid
}

func (l *Lexer) lexLeftParen() LexStateFn {
	l.next()
	l.emit(tokenLeftParen)
	return l.lexVoid
}

func (l *Lexer) lexRightParen() LexStateFn {
	l.next()
	l.emit(tokenRightParen)
	return l.lexVoid
}

func (l *Lexer) lexPlus() LexStateFn {
	l.next()
	l.emit(tokenPlus)
	return l.lexVoid
}

func (l *Lexer) lexMinus() LexStateFn {
	l.next()
	l.emit(tokenMinus)
	return l.lexVoid
}

func (l *Lexer) lexStar() LexStateFn {
	l.next()
	l.emit(tokenStar)
	return l.lexVoid
}

func (l *Lexer) lexSlash() LexStateFn {
	l.next()
	l.emit(tokenSlash)
	return l.lexVoid
}

func (l *Lexer) lexEOF() LexStateFn {
	l.next()
	l.emit(tokenEOF)
	return nil
}

func (l *Lexer) run() {
	for state := l.lexVoid(); state != nil; state = state() {
	}
	close(l.tokens)
}

func lex(input io.Reader) chan token {
	bufferedInput := buffruneio.NewReader(input)
	l := &Lexer{
		input:         bufferedInput,
		tokens:        make(chan token),
		line:          1,
		col:           1,
		endbufferLine: 1,
		endbufferCol:  1,
	}
	go l.run()
	return l.tokens
}

package calculator

import (
	"fmt"
	"github.com/ljun20160606/go-lib/algorithms/stack"
	"math"
	"strconv"
)

type (
	parser struct {
		flow                     chan token
		tokensBuffer             []token
		operationStack, outStack *stack.SimpleStack
	}

	parserStateFn func() parserStateFn
)

var operationPriority = map[tokenType]int{
	tokenPlus:  0,
	tokenMinus: 0,
	tokenStar:  1,
	tokenSlash: 1,
}

func isHigherPriority(op1, op2 token) bool {
	return operationPriority[op1.typ] >= operationPriority[op2.typ]
}

// Formats and panics an error message based on a token
func (p *parser) raiseError(tok *token, msg string, args ...interface{}) {
	panic(tok.Position.String() + ": " + fmt.Sprintf(msg, args...))
}

func (p *parser) peek() *token {
	if len(p.tokensBuffer) != 0 {
		return &(p.tokensBuffer[0])
	}

	tok, ok := <-p.flow
	if !ok {
		return nil
	}
	p.tokensBuffer = append(p.tokensBuffer, tok)
	return &tok
}

func (p *parser) getToken() *token {
	if len(p.tokensBuffer) != 0 {
		tok := p.tokensBuffer[0]
		p.tokensBuffer = p.tokensBuffer[1:]
		return &tok
	}
	tok, ok := <-p.flow
	if !ok {
		return nil
	}
	return &tok
}

func (p *parser) assume(typ tokenType) {
	tok := p.getToken()
	if tok == nil {
		p.raiseError(tok, "was expecting token %s, but token stream is empty", tok)
	}
	if tok.typ != typ {
		p.raiseError(tok, "was expecting token %s, but got %s instead", typ, tok)
	}
}

func (p *parser) parseStart() parserStateFn {
	tok := p.peek()
	// end of stream, parsing is finished
	if tok == nil {
		if p.operationStack.Len != 0 {
			t := p.operationStack.Back().(token)
			p.raiseError(&t, "token %s can't be operated", t)
		}
		return nil
	}
	switch tok.typ {
	case tokenInteger, tokenFloat:
		return p.parseNumber
	case tokenLeftParen:
		return p.parseLeftParen
	case tokenRightParen:
		return p.parseRightParen
	case tokenEOF:
		if p.operationStack.Len != 0 {
			p.applyOperation()
		}
		return nil
	default:
		p.raiseError(tok, "unexpected token")
	}
	return nil
}

func (p *parser) parseNumber() parserStateFn {
	tok := p.getToken()
	p.outStack.Append(tokenToFloat(*tok))
	return p.parseOperation
}

func (p *parser) parseLeftParen() parserStateFn {
	tok := p.getToken()
	next := p.peek()

	if next.typ == tokenEOF {
		p.raiseError(tok, "%s can't be end", tok)
	}

	p.operationStack.Append(*tok)
	return p.parseStart
}

func (p *parser) parseRightParen() parserStateFn {
	tok := p.getToken()
	for p.operationStack.Len > 0 && p.operationStack.Back().(token).typ != tokenLeftParen {
		p.applyOperation()
	}
	defer func() {
		if r := recover(); r != nil {
			p.raiseError(tok, ") can't be start")
		}
	}()
	// pop tokenLeftParen
	p.operationStack.Pop()
	return p.parseStart
}

func (p *parser) parseOperation() parserStateFn {
	tok := p.peek()
	switch tok.typ {
	case tokenLeftParen:
		p.raiseError(tok, "was expecting token %s or %s or %s or %s or %s, but got %s instead", tokenPlus, tokenMinus, tokenStar, tokenSlash, tokenRightParen, tok)
	case tokenRightParen:
		return p.parseRightParen
	}
	p.getToken()
	for p.operationStack.Len > 0 {
		back := p.operationStack.Back().(token)
		if back.typ == tokenLeftParen || !isHigherPriority(back, *tok) {
			break
		}
		p.applyOperation()
	}
	if tok.typ != tokenEOF {
		p.operationStack.Append(*tok)
	}
	return p.parseStart
}

func (p *parser) applyOperation() {
	tok := p.operationStack.Pop().(token)
	switch tok.typ {
	case tokenPlus:
		p.applyPlus()
	case tokenMinus:
		p.applyMinus()
	case tokenStar:
		p.applyMultiply()
	case tokenSlash:
		p.applyDivide()
	default:
		p.raiseError(&tok, "was expecting token %s or %s or %s or %s, but got %s instead", tokenPlus, tokenMinus, tokenStar, tokenSlash, tok)
	}
}

func tokenToFloat(tok token) float64 {
	switch tok.typ {
	case tokenInteger, tokenFloat:
		t, _ := strconv.ParseFloat(tok.val, 64)
		return t
	}
	return math.NaN()
}

func (p *parser) applyPlus() {
	n2, n1 := p.outStack.PopFloat64(), p.outStack.PopFloat64()
	p.outStack.Append(n1 + n2)
}

func (p *parser) applyMinus() {
	n2, n1 := p.outStack.PopFloat64(), p.outStack.PopFloat64()
	p.outStack.Append(n1 - n2)
}

func (p *parser) applyMultiply() {
	n2, n1 := p.outStack.PopFloat64(), p.outStack.PopFloat64()
	p.outStack.Append(n1 * n2)
}

func (p *parser) applyDivide() {
	n2, n1 := p.outStack.PopFloat64(), p.outStack.PopFloat64()
	p.outStack.Append(n1 / n2)
}

func (p *parser) run() {
	for state := p.parseStart(); state != nil; state = state() {
	}
}

func Parse(flow chan token) float64 {
	parser := &parser{
		flow:           flow,
		tokensBuffer:   make([]token, 0),
		operationStack: stack.NewSimpleStack(),
		outStack:       stack.NewSimpleStack(),
	}
	parser.run()
	return parser.outStack.PopFloat64()
}

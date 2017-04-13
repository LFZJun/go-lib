package implement

import "errors"

var NEWNOIMPLEMENT = errors.New("连接池初始化函数未实现")

type ConnectionPool interface {
	Get() interface{}
	Release(conn interface{})
}

type connectionPool struct {
	Size int
	conn chan interface{}
	New  func() (interface{}, error)
}

func (c *connectionPool) size() int {
	if c.Size == 0 {
		return 10
	}
	return c.Size
}

func (c *connectionPool) new() (interface{}, error) {
	if c.New == nil {
		return nil, NEWNOIMPLEMENT
	}
	return c.New()
}

func (c *connectionPool) isEmpty() bool {
	return len(c.conn) == 0
}

func (c *connectionPool) isFull() bool {
	return len(c.conn) == c.size()
}

func (c *connectionPool) Get() interface{} {
	if c.isEmpty() {
		n, err := c.new()
		if err != nil {
			panic(err)
		}
		return n
	}
	return <-c.conn
}

func (c *connectionPool) Release(conn interface{}) {
	switch {
	case c.isFull():
	default:
		c.conn <- conn
	}
}

func (c *connectionPool) init() {
	size := c.size()
	c.conn = make(chan interface{}, size)
	for x := 0; x < size; x++ {
		n, err := c.new()
		if err != nil {
			panic(err)
		}
		c.conn <- n
	}
}

func NewConnectionPool(new func() (interface{}, error), size int) ConnectionPool {
	c := &connectionPool{Size: size, New: new}
	c.init()
	return c
}

package dependencies

import "github.com/uber-go/atomic"

type Operator struct {
	ID      string
	Action  func(o *Operator)
	counter atomic.Int64
}

func NewOperatorMap(queueMax int) *OperatorMap {
	return &OperatorMap{
		inDegree: make(map[*Operator][]*Operator),
		queue:    make(chan *Operator, queueMax),
	}
}

type OperatorMap struct {
	inDegree map[*Operator][]*Operator
	queue    chan *Operator
	counter  atomic.Int64
}

func (om *OperatorMap) AddEdge(in, out *Operator) {
	in.counter.Inc()
	om.inDegree[out] = append(om.inDegree[out], in)
	if _, has := om.inDegree[in]; !has {
		om.inDegree[in] = []*Operator{}
	}
}

func (om *OperatorMap) Run() {
	om.counter.Store(int64(len(om.inDegree)))
	go om.daemon()
	for k, _ := range om.inDegree {
		if k.counter.Load() == 0 {
			om.queue <- k
		}
	}
}

func (om *OperatorMap) daemon() {
	for {
		operator, has := <-om.queue
		if !has {
			return
		}
		go func() {
			operator.Action(operator)
			for _, v := range om.inDegree[operator] {
				dec := v.counter.Dec()
				if dec != 0 {
					continue
				}
				om.queue <- v
			}
			if om.counter.Dec() == 0 {
				close(om.queue)
			}
		}()
	}
}

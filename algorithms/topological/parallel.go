package topological

import "github.com/uber-go/atomic"

type OperatorNode struct {
	// 值
	Val interface{}
	// 触发任务
	Action func(o *OperatorNode)
	// 阈值
	counter atomic.Int64
}

func NewOperatorMap(queueMax int) *OperatorMap {
	return &OperatorMap{
		inDegree: make(map[*OperatorNode][]*OperatorNode),
		queue:    make(chan *OperatorNode, queueMax),
	}
}

type OperatorMap struct {
	// 入度
	inDegree map[*OperatorNode][]*OperatorNode
	// 任务队列
	queue chan *OperatorNode
	// 全部节点数量
	nodeCounter atomic.Int64
}

func (om *OperatorMap) AddEdge(in, out *OperatorNode) {
	in.counter.Inc()
	om.inDegree[out] = append(om.inDegree[out], in)
	if _, has := om.inDegree[in]; !has {
		om.inDegree[in] = []*OperatorNode{}
	}
}

func (om *OperatorMap) Run() {
	om.nodeCounter.Store(int64(len(om.inDegree)))
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
			if om.nodeCounter.Dec() == 0 {
				close(om.queue)
			}
		}()
	}
}

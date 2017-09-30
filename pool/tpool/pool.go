package tpool

import (
	"errors"
	"github.com/uber-go/atomic"
	"sync"
)

// http://www.jianshu.com/p/87bff5cc8d8c

type PoolExecutor struct {
	ctl             atomic.Int32
	corePoolSize    int
	maximumPoolSize int
	largestPoolSize int
	tunnel          chan Command
	workQueue       Queue
	mu              sync.Mutex
	workers         map[interface{}]struct{}
}

func (p *PoolExecutor) Execute(command Command) error {
	if command == nil {
		return errors.New("command can't be nil")
	}
	c := p.ctl.Load()
	if workerCountOf(int(c)) < p.corePoolSize {
		if p.addWork(command, true) {
			return nil
		}
		c = p.ctl.Load()
	}
	if isRunning(int(c)) && p.workQueue.Offer(command) {
		recheck := p.ctl.Load()
		if !isRunning(int(recheck)) && p.remove(command) {
			// reject(command)
		} else if workerCountOf(int(recheck)) == 0 {
			p.addWork(nil, false)
		}
	} else if !p.addWork(command, false) {
		// reject(command)
	}
	return nil
}

func (p *PoolExecutor) tryTerminate() {
	for {
		c := p.ctl.Load()
		if isRunning(int(c)) ||
			runStateAtLeast(int(c), TIDYING) ||
			(runStateOf(int(c)) == SHUTDOWN && !p.workQueue.IsEmpty()) {
			return
		}
		if workerCountOf(int(c)) != 0 {

		}
	}
}

func (p *PoolExecutor) remove(command Command) bool {
	removed := p.workQueue.Remove(command)
	p.tryTerminate()
	return removed
}

func (p *PoolExecutor) addWorkerFailed(worker *Worker) {
	p.mu.Lock()
	if worker != nil {
		delete(p.workers, worker)
		for {
			c := p.ctl.Load()
			if p.ctl.CAS(c, c-1) {
				break
			}
		}
	}
	p.mu.Unlock()
}

func (p *PoolExecutor) addWork(firstTask Command, core bool) bool {
retry:
	for {
		c := p.ctl.Load()
		rs := runStateOf(int(c))
		if rs >= SHUTDOWN &&
			!(rs == SHUTDOWN &&
				firstTask == nil &&
				!p.workQueue.IsEmpty()) {
			return false
		}
		for {
			wc := workerCountOf(int(c))
			if wc >= CAPACITY {
				return false
			}
			var poolSize int
			if core {
				poolSize = p.corePoolSize
			} else {
				poolSize = p.maximumPoolSize
			}
			if wc >= poolSize {
				return false
			}
			// 计数器+1
			if p.ctl.CAS(c, c+1) {
				break retry
			}
			c = p.ctl.Load()
			if runStateOf(int(c)) != rs {
				continue retry
			}
		}
	}
	var workerStarted, workerAdded bool
	p.mu.Lock()
	w := &Worker{tunnel: make(chan Command, 1)}
	w.Put(firstTask)
	rs := runStateOf((int)(p.ctl.Load()))
	if rs < SHUTDOWN || (rs == SHUTDOWN && firstTask == nil) {
		// workers add
		p.workers[w] = struct{}{}
		s := len(p.workers)
		if s > p.largestPoolSize {
			p.largestPoolSize = s
		}
		workerAdded = true
	}
	p.mu.Unlock()
	if workerAdded {
		w.Start()
		workerStarted = true
	}
	if !workerStarted {

	}
	return workerStarted
}

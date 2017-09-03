package timer

import (
	"container/list"
	"sync"
	"time"
)

type (
	Task struct {
		Timeout time.Duration
		Work    func()
		index   *list.Element
		count   uint
		pos     uint
	}
	Slot struct {
		mu    sync.Mutex
		tasks *list.List
	}
	TimingWheel struct {
		// 刻度数量，决定锁粒度、周期时长
		slotsNum uint
		// 刻度，决定最小精确、周期时长
		interval time.Duration
		slots    []*Slot
		quit     chan struct{}
		pos      uint
		ticker   *time.Ticker
	}
)

func (s *Slot) Put(t *Task) {
	s.mu.Lock()
	t.index = s.tasks.PushBack(t)
	s.mu.Unlock()
}

func (s *Slot) Del(t *Task) {
	s.mu.Lock()
	s.tasks.Remove(t.index)
	s.mu.Unlock()
}

func (s *Slot) Walk() {
	s.mu.Lock()
	for cur := s.tasks.Front(); cur != nil; cur = cur.Next() {
		if task := cur.Value.(*Task); task.count > 0 {
			task.count -= 1
		} else {
			task.Work()
			s.tasks.Remove(task.index)
		}
	}
	s.mu.Unlock()
}

func (t *TimingWheel) run() {
	t.slots[0].Walk()
	for {
		select {
		case <-t.ticker.C:
			t.onTick()
		case <-t.quit:
			t.ticker.Stop()
			return
		}
	}

}

func (t *TimingWheel) cal(timeout time.Duration) (pos uint, count uint) {
	nums := uint(timeout / t.interval)
	if nums > t.slotsNum {
		count = nums / t.slotsNum
	}
	pos = (t.pos + nums) % t.slotsNum
	return
}

func (t *TimingWheel) After(task *Task) *Task {
	task.pos, task.count = t.cal(task.Timeout)
	t.slots[task.pos].Put(task)
	return task
}

func (t *TimingWheel) Del(task *Task) {
	t.slots[task.pos].Del(task)
}

func (t *TimingWheel) onTick() {
	t.pos = (t.pos + 1) % t.slotsNum
	t.slots[t.pos].Walk()
}

func (t *TimingWheel) Stop() {
	close(t.quit)
}

func NewTimingWheel(interval time.Duration, slotsNum uint) (t *TimingWheel) {
	if slotsNum == 0 {
		return nil
	}
	t = &TimingWheel{
		slotsNum: slotsNum,
		interval: interval,
		slots:    make([]*Slot, slotsNum),
		quit:     make(chan struct{}),
		ticker:   time.NewTicker(interval),
	}
	for i := range t.slots {
		t.slots[i] = &Slot{tasks: list.New()}
	}
	go t.run()
	return
}

package dependencies

type Member struct {
	Val          interface{}
	Dependencies []*Member
	Action       func(m *Member)
	done         bool
}

func NewMember(val interface{}, f func(m *Member)) *Member {
	return &Member{Val: val, Action: f}
}

func (r *Member) Add(rr *Member) {
	if r == rr {
		panic("不能依赖自己")
	}
	r.Dependencies = append(r.Dependencies, rr)
}

func (r *Member) Do() {
	if r.done {
		return
	}
	for _, d := range r.Dependencies {
		d.Do()
	}
	r.Action(r)
	r.done = true
}

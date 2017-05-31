package ioc

type (
	Plugin interface {
		Value(path string) interface{}
		Prefix() string
		Priority() int
	}
	plugins []Plugin
)

func (pl plugins) Len() int {
	return len(pl)
}

func (pl plugins) Less(i, j int) bool {
	return pl[i].Priority() < pl[j].Priority()
}

func (pl plugins) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

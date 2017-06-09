package slice

const (
	NilPtr errorSlice = iota
	MustPtr
	MustSlice
	MustSameType
)

var errorss = [...]string{
	NilPtr:       "nil pointer passed to StructScan destination",
	MustPtr:      "must pass a pointer, not a value, to StructScan destination",
	MustSlice:    "must pass a slice pointer with src",
	MustSameType: "must pass same type of dest, src",
}

type errorSlice int

func (e errorSlice) String() string {
	return errorss[e]
}

func (e errorSlice) Error() string {
	return errorss[e]
}

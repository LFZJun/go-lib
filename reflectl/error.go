package reflectl

const (
	NilPtr errorSlice = iota
	MustPtr
	MustSlice
)

var errorss = [...]string{
	NilPtr:    "nil pointer passed to StructScan destination",
	MustPtr:   "must pass a pointer, not a value, to StructScan destination",
	MustSlice: "must pass a slice pointer with src",
}

type errorSlice int

func (e errorSlice) String() string {
	return errorss[e]
}

func (e errorSlice) Error() string {
	return errorss[e]
}

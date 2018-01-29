package tpool

const (
	COUNT_BITS = 29
	CAPACITY   = (1 << COUNT_BITS) - 1
)
const (
	RUNNING = (iota - 1) << COUNT_BITS
	SHUTDOWN
	STOP
	TIDYING
	TERMINATED
)

func runStateOf(c int) int {
	return c & ^CAPACITY
}

func workerCountOf(c int) int {
	return c & CAPACITY
}

func ctlOf(rc, wc int) int {
	return rc | wc
}

func isRunning(c int) bool {
	return c < SHUTDOWN
}

func runStateAtLeast(c, s int) bool {
	return c >= s
}

package ioc

type (
	HolderFunc func(holder *Holder)

	HolderSet map[*Holder]struct{}
)

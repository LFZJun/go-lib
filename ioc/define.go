package ioc

type HolderFunc func(holder *Holder)

type HolderSet map[*Holder]struct{}

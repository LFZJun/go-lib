package ioc

type HolderFunc func(holder *holder)

type HolderSet map[*holder]struct{}

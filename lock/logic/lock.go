package logic

type Mutex interface {
	Lock(var1 string) error
	Unlock(var1 string) error
}

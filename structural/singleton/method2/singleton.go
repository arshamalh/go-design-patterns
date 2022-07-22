package method2

import "sync"

type single struct {}

var (
	once sync.Once
	instance *single
)

func GetInstance() *single {
	once.Do(func() {
		instance = &single{}
	})

	return instance
}
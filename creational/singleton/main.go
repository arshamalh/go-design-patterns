package main

import (
	"fmt"
	"singleton/method1"
	// "singleton/method2"
)

func main() {
	for i:=0; i < 100; i ++ {
		instance := method1.GetInstance()
		// instance := method2.GetInstance()
		fmt.Println(instance)
	}
}

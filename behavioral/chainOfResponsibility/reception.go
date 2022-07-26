package main

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registerationDone {
		fmt.Println("Patient registeration is already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registerationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department) {
	r.next = next
}
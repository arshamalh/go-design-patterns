package main

import "fmt"

type File struct {
	name string
}

func (f *File) print(identation string) {
	fmt.Println(identation + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}


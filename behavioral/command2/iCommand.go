package main

type Command interface {
	execute() // could have a error
}
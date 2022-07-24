package main

type IServer interface {
	handleRequest(string, string) (int, string)
}

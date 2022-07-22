package main

import "fmt"

func main() {

    hpPrinter := &Hp{}
    epsonPrinter := &Epson{}

    macComputer := &Mac{}
    winComputer := &Windows{}

	/* Mac combinations */
    macComputer.SetPrinter(hpPrinter)
    macComputer.Print()
    fmt.Println()

    macComputer.SetPrinter(epsonPrinter)
    macComputer.Print()
    fmt.Println()

	/* Windows combinations */
    winComputer.SetPrinter(hpPrinter)
    winComputer.Print()
    fmt.Println()

    winComputer.SetPrinter(epsonPrinter)
    winComputer.Print()
    fmt.Println()
}
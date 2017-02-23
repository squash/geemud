package main

import (
	"bufio"
	"fmt"
	"os"
)

type Iface struct {
	Name    string
	Scanner *bufio.Scanner
}

func CreateIface(name string) *Iface {
	i := new(Iface)
	i.Name = name
	i.Scanner = bufio.NewScanner(os.Stdin)
	return i
}

func (i *Iface) Add(text string) {
	fmt.Print(text)
}

func (i *Iface) Read() []byte {
	i.Scanner.Scan()
	in := i.Scanner.Text()
	return []byte(in)
}

package main

import (
	"io/ioutil"
	"maze/cell"
	"maze/svg"
)

func main() {
	c := cell.NewDividedCell(9, 9)
	f, _ := svg.Format(c)
	ioutil.WriteFile("tmp.svg", f, 0644)
}

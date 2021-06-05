package main

import (
	"encoding/json"
	"fmt"
	"maze/cell"
)

func main() {
	c := cell.NewDividedCell(9, 9)
	j, _ := json.MarshalIndent(c, "", " ")
	fmt.Print(string(j))
}
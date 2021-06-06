package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"maze/cell"
	"maze/svg"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	log.SetOutput(os.Stderr)
	log.Print("logger works!")
	c := cell.NewDividedCell(9, 9)
	f, err := svg.Format(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ioutil.WriteFile("tmp.svg", f, 0644)
	log.Print("done")
}

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

const scale = 50

func main() {
	rand.Seed(time.Now().Unix())
	//rand.Seed(0)
	log.SetOutput(os.Stderr)
	log.Print("logger works!")
	n := 20
	c := cell.NewDividedCell(n, n)
	f, err := svg.Format(c, scale)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ioutil.WriteFile("tmp.svg", f, 0644)
	log.Print("done")
}

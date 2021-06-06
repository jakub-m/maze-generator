package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"maze/cell"
	"maze/svg"
	"os"
	"time"
)

var outfname string
var scale int
var seed int64
var size int
var strokeWidth int

func init() {
	flagScale := flag.Int("scale", 50, "scale")
	flagSeed := flag.Int64("seed", time.Now().Unix(), "random seed")
	flagSize := flag.Int("size", 6, "size")
	flagStroke := flag.Int("stroke", 2, "stroke width")
	flagOutfile := flag.String("out", "-", "output filename or \"-\"")
	flagVerbose := flag.Bool("verbose", false, "verbose mode")
	flag.Parse()

	outfname = *flagOutfile
	scale = *flagScale
	seed = *flagSeed
	size = *flagSize
	strokeWidth = *flagStroke
	if *flagVerbose {
		log.SetOutput(os.Stderr)
	} else {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	log.Println("scale", scale)
	log.Println("seed", seed)
	log.Println("size", size)
	log.Println("stroke", strokeWidth)
	log.Println("out", outfname)
	rand.Seed(seed)
	c := cell.NewDividedCell(size, size)
	f, err := svg.Format(c, scale, strokeWidth)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if outfname == "-" {
		os.Stdout.Write(f)
	} else {
		ioutil.WriteFile(outfname, f, 0644)
	}
	log.Println("done")
}
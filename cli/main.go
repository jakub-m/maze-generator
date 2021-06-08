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

type params struct {
	outfname    string
	scale       int
	seed        int64
	size        int
	strokeWidth int
	verbose     bool
}

func main() {
	p := getParams()
	if p.verbose {
		log.SetOutput(os.Stderr)
	} else {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	log.Printf("params %+v", p)
	rand.Seed(p.seed)
	m := cell.NewMaze(p.size)
	f, err := svg.FormatMaze(m, p.scale, p.strokeWidth)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if p.outfname == "-" {
		os.Stdout.Write(f)
	} else {
		ioutil.WriteFile(p.outfname, f, 0644)
	}
	log.Println("done")
}

func getParams() params {
	scale := flag.Int("scale", 50, "scale")
	seed := flag.Int64("seed", time.Now().Unix(), "random seed")
	size := flag.Int("size", 6, "size")
	strokeWidth := flag.Int("stroke", 2, "stroke width")
	outfname := flag.String("out", "-", "output filename or \"-\"")
	verbose := flag.Bool("verbose", false, "verbose mode")
	flag.Parse()
	return params{
		scale:       *scale,
		seed:        *seed,
		size:        *size,
		strokeWidth: *strokeWidth,
		outfname:    *outfname,
		verbose:     *verbose,
	}
}

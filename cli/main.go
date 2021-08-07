package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"maze/mazelib"
	"maze/svg"
	"os"
	"time"
)

type params struct {
	height      int
	outfname    string
	scale       int
	seed        int64
	strokeWidth int
	verbose     bool
	width       int
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
	m, err := mazelib.NewMaze(p.width, p.height)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ok")
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
	height := flag.Int("h", 6, "height")
	outfname := flag.String("o", "-", "output filename or \"-\"")
	scale := flag.Int("s", 50, "scale")
	seed := flag.Int64("r", time.Now().Unix(), "random seed")
	strokeWidth := flag.Int("t", 2, "stroke width")
	verbose := flag.Bool("v", false, "verbose mode")
	width := flag.Int("w", 6, "width")
	flag.Parse()
	return params{
		height:      *height,
		outfname:    *outfname,
		scale:       *scale,
		seed:        *seed,
		strokeWidth: *strokeWidth,
		verbose:     *verbose,
		width:       *width,
	}
}

package svg

import (
	"bytes"
	"fmt"
	"log"
	"maze/mazelib"
)

const newLine = "\n"

const svgHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`

func FormatMaze(m *mazelib.Maze, scale, strokeWidth int) ([]byte, error) {
	offset := pos{x: strokeWidth, y: strokeWidth}
	var buf bytes.Buffer
	_, err := buf.WriteString(svgHeader + newLine)
	if err != nil {
		return nil, err
	}
	openTag := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, m.Width*scale+offset.x+strokeWidth, m.Height*scale+offset.y+strokeWidth)
	_, err = buf.WriteString(openTag + newLine)
	if err != nil {
		return nil, err
	}
	forEachWall(m, func(w wall) {
		g := fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="%d"/>`,
			w.x0*scale+offset.x,
			w.y0*scale+offset.y,
			w.x1*scale+offset.x,
			w.y1*scale+offset.y,
			strokeWidth)
		buf.WriteString(g + newLine)
	})
	_, err = buf.WriteString(`</svg>`)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func forEachWall(m *mazelib.Maze, fn func(w wall)) {
	for ix := -1; ix < m.Width+1; ix++ {
		for iy := -1; iy < m.Height+1; iy++ {
			cRef := mazelib.Cell{X: ix, Y: iy}
			cEast := mazelib.Cell{X: ix + 1, Y: iy}
			if !m.HasPassage(cRef, cEast) {
				fn(wall{
					x0: cEast.X,
					y0: cEast.Y,
					x1: cEast.X,
					y1: cEast.Y + 1,
				})
			} else {
				log.Printf("has passage: %s -> %s", cRef, cEast)
			}
			cSouth := mazelib.Cell{X: ix, Y: iy + 1}
			if !m.HasPassage(cRef, cSouth) {
				fn(wall{
					x0: cSouth.X,
					y0: cSouth.Y,
					x1: cSouth.X + 1,
					y1: cSouth.Y,
				})
			} else {
				log.Printf("has passage: %s -> %s", cRef, cSouth)
			}
		}
	}
}

type wall struct{ x0, y0, x1, y1 int }
type pos struct{ x, y int }

package svg

import (
	"bytes"
	"fmt"
	"maze/cell"
	"maze/maze"
)

const newLine = "\n"

const svgHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`

func FormatMaze(m maze.Maze, scale, strokeWidth int) ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.WriteString(svgHeader + newLine)
	if err != nil {
		return nil, err
	}
	openTag := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, m.Width*scale, m.Height*scale)
	_, err = buf.WriteString(openTag + newLine)
	if err != nil {
		return nil, err
	}
	formatWalls(m.RootCell, scale, strokeWidth, &buf)
	_, err = buf.WriteString(`</svg>`)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func formatWalls(c cell.Cell, scale, strokeWidth int, buf *bytes.Buffer) {
	walls := cell.InternalWalls(c)
	for _, w := range walls {
		for _, line := range w.Parts {
			g := fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="%d"/>`,
				line.Start.X*scale,
				line.Start.Y*scale,
				line.End.X*scale,
				line.End.Y*scale,
				strokeWidth)
			buf.WriteString(g + newLine)
		}
	}
}

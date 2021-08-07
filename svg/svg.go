package svg

import (
	"bytes"
	"fmt"
	"maze/mazelib"
)

const newLine = "\n"

const svgHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`

func FormatMaze(m *mazelib.Maze, scale, strokeWidth int) ([]byte, error) {
	offset := mazelib.Pos{X: strokeWidth, Y: strokeWidth}
	var buf bytes.Buffer
	_, err := buf.WriteString(svgHeader + newLine)
	if err != nil {
		return nil, err
	}
	openTag := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, m.Width*scale+offset.X+strokeWidth, m.Height*scale+offset.Y+strokeWidth)
	_, err = buf.WriteString(openTag + newLine)
	if err != nil {
		return nil, err
	}
	m.ForEachWall(func(a, b mazelib.Pos) {
		g := fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="%d"/>`,
			a.X*scale+offset.X,
			a.Y*scale+offset.Y,
			b.X*scale+offset.X,
			b.Y*scale+offset.Y,
			strokeWidth)
		buf.WriteString(g + newLine)
	})
	_, err = buf.WriteString(`</svg>`)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

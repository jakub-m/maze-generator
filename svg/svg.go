package svg

import (
	"bytes"
	"fmt"
	"maze/cell"
)

const newLine = "\n"

const svgHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`

func Format(c cell.Cell, scale int) ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.WriteString(svgHeader + newLine)
	if err != nil {
		return nil, err
	}
	openTag := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, c.Dim.Width*scale, c.Dim.Height*scale)
	_, err = buf.WriteString(openTag + newLine)
	if err != nil {
		return nil, err
	}
	formatWalls(c, scale, &buf)
	_, err = buf.WriteString(`</svg>`)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func formatWalls(c cell.Cell, scale int, buf *bytes.Buffer) {
	walls := cell.InternalWalls(c)
	for _, w := range walls {
		for _, line := range w.Parts {
			g := fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black"/>`,
				line.Start.X*scale,
				line.Start.Y*scale,
				line.End.X*scale,
				line.End.Y*scale)
			buf.WriteString(g + newLine)
		}
	}
}

package svg

import (
	"bytes"
	"fmt"
	"log"
	"maze/cell"
)

const scale = 50

const svgHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`

func Format(c cell.Cell) ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.WriteString(svgHeader + "\n")
	if err != nil {
		return nil, err
	}
	openTag := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, c.Dim.Width * scale, c.Dim.Height * scale)
	_, err = buf.WriteString(openTag)
	if err != nil {
		return nil, err
	}
	err = formatCell(c, cell.Pos{0, 0}, &buf)
	if err != nil  {
		return nil, err
	}
	_, err = buf.WriteString(`</svg>`)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func formatCell(c cell.Cell, origin cell.Pos, buf *bytes.Buffer) error {
	log.Printf("formatCell %v", origin)
	if isLeaf(c) {
		return formatLeafCell(c, origin, buf)
	} else {
		for _, s := range c.Subcells {
			newOrigin := origin.Add(s.RelativePos)
			err := formatCell(s.Cell, newOrigin, buf)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func formatLeafCell(c cell.Cell, origin cell.Pos, buf *bytes.Buffer) error {
	log.Printf("formatLeafCell %v", origin)
	x := origin.X * scale
	y := origin.Y * scale
	width := c.Dim.Width * scale
	height := c.Dim.Height * scale
	r := fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" style="fill:rgb(0,0,255);stroke-width:3;stroke:rgb(0,0,0)"/>`, x, y, width, height)
	_, err := buf.WriteString(r)
	return err
}

func isLeaf(c cell.Cell) bool {
	return len(c.Subcells) == 0
}
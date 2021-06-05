package svg

import (
	"bytes"
	"fmt"
	"maze/cell"
)

const svgHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`

func Format(cell cell.Cell) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(svgHeader)
	openTag := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">`, cell.Dim.Width, cell.Dim.Height)
	buf.WriteString(openTag)
	buf.WriteString(`</svg>`)
	return buf.Bytes(), nil
}

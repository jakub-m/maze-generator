package svg

import (
	"bytes"
	"fmt"
	"maze/cell"
)

const newLine = "\n"

const svgHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`

func FormatMaze(m cell.Maze, scale, strokeWidth int) ([]byte, error) {
	offset := cell.Pos{strokeWidth, strokeWidth}
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
	formatWalls(m, offset, scale, strokeWidth, &buf)
	_, err = buf.WriteString(`</svg>`)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func formatWalls(m cell.Maze, off cell.Pos, scale, strokeWidth int, buf *bytes.Buffer) {
	walls := cell.InternalWalls(m.RootCell)
	walls = append(walls, m.OutterWalls...)
	for _, w := range walls {
		for _, line := range w.Parts {
			g := fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="black" stroke-width="%d"/>`,
				line.Start.X*scale+off.X,
				line.Start.Y*scale+off.Y,
				line.End.X*scale+off.X,
				line.End.Y*scale+off.Y,
				strokeWidth)
			buf.WriteString(g + newLine)
		}
	}
}

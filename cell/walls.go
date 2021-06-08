package cell

import "fmt"

type Wall struct {
	Parts []Line
}

type Line struct {
	Start, End Pos
}

func OutterWalls(mazeDim Dim) []Wall {
	w := mazeDim.Width
	h := mazeDim.Height
	h1 := Wall{[]Line{{Pos{0, 0}, Pos{w, 0}}}}
	h2 := Wall{[]Line{{Pos{0, h}, Pos{w, h}}}}
	v1p := randomPassageOffset(h)
	v1 := Wall{[]Line{
		{Pos{0, 0}, Pos{0, v1p}},
		{Pos{0, v1p + 1}, Pos{0, h}},
	}}
	v2p := randomPassageOffset(h)
	v2 := Wall{[]Line{
		{Pos{w, 0}, Pos{w, v2p}},
		{Pos{w, v2p + 1}, Pos{w, h}},
	}}
	return []Wall{h1, h2, v1, v2}
}

// InternalWalls returns non-overlapping set of walls with absolute coordinates.
func InternalWalls(c Cell) []Wall {
	zero := Pos{0, 0}
	return internalWallsRec(zero, c)
}

func internalWallsRec(absOrigin Pos, c Cell) []Wall {
	var walls []Wall
	for _, sub := range c.Subcells {
		if !isFirstSub(sub) {
			if isSplitHor(sub) {
				w := topWall(absOrigin, sub)
				walls = append(walls, w)
			} else if isSplitVer(sub) {
				w := leftWall(absOrigin, sub)
				walls = append(walls, w)
			} else {
				panic(fmt.Sprintf("BUG. Unreachable state %v", sub))
			}
		}
		recWalls := internalWallsRec(absOrigin.Add(sub.RelativePos), sub.Cell)
		walls = append(walls, recWalls...)
	}
	return walls
}

func isFirstSub(sub Subcell) bool {
	return sub.RelativePos.X == 0 && sub.RelativePos.Y == 0
}

// isSplitHor tells if the subcell was created after horizontal split
func isSplitHor(sub Subcell) bool {
	return sub.RelativePos.X == 0 && sub.RelativePos.Y != 0
}

// isSplitVer tells if the subcell was created after vertical split
func isSplitVer(sub Subcell) bool {
	return sub.RelativePos.X != 0 && sub.RelativePos.Y == 0
}

func topWall(absOrigin Pos, sub Subcell) Wall {
	width := sub.Cell.Dim.Width
	p := sub.PassageOffset
	parts := []Line{
		{
			Start: Pos{
				0,
				0},
			End: Pos{
				X: p,
				Y: 0,
			},
		},
		{
			Start: Pos{
				p + 1,
				0},
			End: Pos{
				X: width,
				Y: 0,
			},
		},
	}
	translation := absOrigin.Add(sub.RelativePos)
	translateLines(translation, parts)
	parts = filterInvisibleLines(parts)
	return Wall{parts}
}

func leftWall(absOrigin Pos, sub Subcell) Wall {
	height := sub.Cell.Dim.Height
	p := sub.PassageOffset
	parts := []Line{
		{
			Start: Pos{
				0,
				0},
			End: Pos{
				X: 0,
				Y: p,
			},
		},
		{
			Start: Pos{
				0,
				p + 1},
			End: Pos{
				X: 0,
				Y: height,
			},
		},
	}
	translation := absOrigin.Add(sub.RelativePos)
	translateLines(translation, parts)
	parts = filterInvisibleLines(parts)
	return Wall{parts}
}

func filterInvisibleLines(lines []Line) []Line {
	var filtered []Line
	for _, e := range lines {
		if e.Start.X == e.End.X && e.Start.Y == e.End.Y {
			continue
		} else {
			filtered = append(filtered, e)
		}
	}
	return filtered

}

func translateLines(translation Pos, lines []Line) {
	for i, l := range lines {
		lines[i] = Line{
			Start: l.Start.Add(translation),
			End:   l.End.Add(translation),
		}
	}
}

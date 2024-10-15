package screen

type Point struct {
	X    int
	Y    int
	Char rune
}

func (p *Point) SetOnScreen(s Screen, content [][]rune) [][]rune {
	crossX := p.X < 0 || p.X >= s.W
	crossY := p.Y < 0 || p.Y >= s.H

	if !crossY && !crossX {
		content[p.Y][p.X] = p.Char
	}

	return content
}

func (s *Screen) AddPoint(x int, y int, char rune) {
	s.pts = append(s.pts, Point{X: x, Y: y, Char: char})
}

func (s *Screen) AddPointAtB(x int, y int, char rune) {
	s.pts = append(s.pts, Point{X: x, Y: s.H - 1 - y, Char: char})
}

func (s *Screen) ClearPoints() {
	s.pts = []Point{}
}

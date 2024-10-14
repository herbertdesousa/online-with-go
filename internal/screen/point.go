package screen

type Point struct {
	X    int
	Y    int
	Char rune
}

func (p *Point) SetOnScreen(s screen, content [][]rune) [][]rune {
	crossX := p.X < 0 || p.X >= s.w
	crossY := p.Y < 0 || p.Y >= s.h

	if !crossY && !crossX {
		content[p.Y][p.X] = p.Char
	}

	return content
}

func (s *screen) AddPoint(x int, y int, char rune) {
	s.pts = append(s.pts, Point{X: x, Y: y, Char: char})
}

func (s *screen) AddPointAtB(x int, y int, char rune) {
	s.pts = append(s.pts, Point{X: x, Y: s.h - 1 - y, Char: char})
}

func (s *screen) ClearPoints() {
	s.pts = []Point{}
}

package screen

import "fmt"

type screen struct {
	w   int
	h   int
	pts []Point
}

func NewScreen(w int, h int) *screen {
	return &screen{w: w, h: h}
}

func (s *screen) build() [][]rune {
	content := make([][]rune, s.h)

	for y := 0; y < s.h; y++ {
		row := make([]rune, s.w)

		for x := 0; x < s.w; x++ {
			row[x] = '#'
		}

		content[y] = row
	}

	for i := 0; i < len(s.pts); i++ {
		content = s.pts[i].SetOnScreen(*s, content)
	}

	return content
}

func (s *screen) print(r [][]rune) {
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			fmt.Printf("%c", r[i][j])
		}
		fmt.Println()
	}
}

func (s *screen) Draw() {
	s.print(s.build())
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

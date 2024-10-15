package screen

import "fmt"

type Screen struct {
	W   int
	H   int
	pts []Point
}

func NewScreen(w int, h int) *Screen {
	return &Screen{W: w, H: h}
}

func (s *Screen) build() [][]rune {
	content := make([][]rune, s.H)

	for y := 0; y < s.H; y++ {
		row := make([]rune, s.W)

		for x := 0; x < s.W; x++ {
			row[x] = '#'
		}

		content[y] = row
	}

	for i := 0; i < len(s.pts); i++ {
		content = s.pts[i].SetOnScreen(*s, content)
	}

	return content
}

func (s *Screen) print(r [][]rune) {
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			fmt.Printf("%c", r[i][j])
		}
		fmt.Println()
	}
}

func (s *Screen) Draw() {
	s.print(s.build())
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

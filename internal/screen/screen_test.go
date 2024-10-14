package screen

import (
	"testing"
)

func compareRuneMatrices(matrix1, matrix2 [][]rune) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}

	for i := 0; i < len(matrix1); i++ {
		if len(matrix1[i]) != len(matrix2[i]) {
			return false
		}
	}

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[i]); j++ {
			if matrix1[i][j] != matrix2[i][j] {
				return false
			}
		}
	}

	return true
}

func Test_screen_build(t *testing.T) {
	type fields struct {
		w   int
		h   int
		pts []Point
	}

	s1 := NewScreen(5, 4)
	s1.AddPoint(0, 0, 'a')
	s1.AddPoint(1, 0, 'b')
	s1.AddPoint(0, 3, 'c')
	s1.AddPoint(4, 1, 'd')
	s1.AddPoint(4, 3, 'e')
	s1.AddPoint(0, 0, 'A') // overlap Char 'a'
	s1.AddPointAtB(1, 0, '1')
	// out of bounds
	s1.AddPoint(-1, 0, 'W')
	s1.AddPoint(20, 0, 'W')
	s1.AddPoint(0, -1, 'W')
	s1.AddPoint(0, 20, 'W')

	tests := []struct {
		name   string
		fields screen
		want   [][]rune
	}{
		{
			name:   "render 5x4 screen with points",
			fields: *s1,
			want: [][]rune{
				{'A', 'b', '#', '#', '#'},
				{'#', '#', '#', '#', 'd'},
				{'#', '#', '#', '#', '#'},
				{'c', '1', '#', '#', 'e'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &screen{
				w:   tt.fields.w,
				h:   tt.fields.h,
				pts: tt.fields.pts,
			}

			if !compareRuneMatrices(s.build(), tt.want) {
				t.Error()
			}
		})
	}
}

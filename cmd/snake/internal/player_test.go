package internal

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	p := player{}
	p3c1 := playerCell{x: 2, y: 0, direction: CELL_RIGHT}
	p3c2 := playerCell{x: 1, y: 0, direction: CELL_RIGHT}
	p3c3 := playerCell{x: 0, y: 0, direction: CELL_RIGHT}

	p3c1.previous = &p3c2
	p3c2.next = &p3c1

	p3c2.previous = &p3c3
	p3c3.next = &p3c2

	p.head = &p3c1
	p.tail = &p3c3

	tests := []struct {
		name string
		want *player
	}{
		{name: "min 3 body", want: &p},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlayer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func countCells(p *player) int {
	previous := p.head.previous
	cells := 1

	for {
		if previous != nil {
			cells++
			previous = previous.previous
		} else {
			break
		}
	}

	return cells
}

func Test_player_AddTail(t *testing.T) {
	// add to null tail
	// add to screen boundaries

	p1 := NewPlayer()

	t.Run("check default player size", func(t *testing.T) {
		if cells := countCells(p1); cells != 3 {
			t.Error()
		}
	})

	// =-->
	p1.AddTail()

	t.Run("add 1 cell to the left of tail cell directed to the right", func(t *testing.T) {
		if cells := countCells(p1); cells != 4 {
			t.Error()
		}
		if p1.tail.x == 4 {
			t.Error()
		}
	})

	// <--=

	//  ^
	//  |
	// | | <- add here

	// | | <- add here
	//  |
	//  v

	/*type fields struct {
		cells []playerCell
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "added to =-->",
			fields: fields{
				cells: []playerCell{p1c1, p1c2, p1c3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1 := &player{
				cells: tt.fields.cells,
			}
			p1.AddTail()
		})
	}*/
}

func compareCell(celA playerCell, celB playerCell) error {
	baseErr := fmt.Sprintf("comparsion celA{x: %d, y: %d, dir: %d}, celB{x: %d, y: %d, dir: %d} fail", celA.x, celA.y, celA.direction, celB.x, celB.y, celB.direction)

	if celA.x != celB.x {
		return errors.New(fmt.Sprintf("x conflict, %s", baseErr))
	} else if celA.y != celB.y {
		return errors.New(fmt.Sprintf("y conflict, %s", baseErr))
	} else if celA.direction != celB.direction {
		return errors.New(fmt.Sprintf("direction conflict, %s", baseErr))
	}

	return nil
}

type composableTests struct {
	got  playerCell
	want playerCell
}

func composeTests(t *testing.T, tests []composableTests) {
	for _, tt := range tests {
		if err := compareCell(tt.got, tt.want); err != nil {
			t.Error(fmt.Sprintf("[%s] %s", t.Name(), err))
		}
	}
}

func Test_player_Move(t *testing.T) {

	t.Run("up", func(t *testing.T) {
		//    ^
		// ---┘
		t.Run("from right", func(t *testing.T) {
			p := player{}
			p.addTailAt(4, 0, CELL_RIGHT) // head
			p.addTailAt(3, 0, CELL_RIGHT) // mid
			p.addTailAt(2, 0, CELL_RIGHT) // mid
			p.addTailAt(1, 0, CELL_RIGHT) // tail

			p.Move(MOVE_UP)

			composeTests(
				t,
				[]composableTests{
					{got: *p.head, want: playerCell{x: 4, y: 1, direction: CELL_UP}},
					{got: *p.head.previous, want: playerCell{x: 4, y: 0, direction: CELL_CORNER_RIGHT_UP}},
					{got: *p.head.previous.previous, want: playerCell{x: 3, y: 0, direction: CELL_RIGHT}},
					{got: *p.head.previous.previous.previous, want: playerCell{x: 2, y: 0, direction: CELL_RIGHT}},
				},
			)

			for p.head.previous.previous.previous.previous != nil {
				t.Error("invalid player size")
			}
		})

		// ^
		// └---
		t.Run("from left", func(t *testing.T) {
			p := player{}
			p.addTailAt(7, 0, CELL_LEFT) // head
			p.addTailAt(8, 0, CELL_LEFT) // mid
			p.addTailAt(9, 0, CELL_LEFT) // tail

			p.Move(MOVE_UP)

			composeTests(
				t,
				[]composableTests{
					{got: *p.head, want: playerCell{x: 7, y: 1, direction: CELL_UP}},
					{got: *p.head.previous, want: playerCell{x: 7, y: 0, direction: CELL_CORNER_LEFT_UP}},
					{got: *p.head.previous.previous, want: playerCell{x: 8, y: 0, direction: CELL_LEFT}},
				},
			)

			if p.head.previous.previous.previous != nil {
				t.Error("cell unexisting")
			}
		})

		// ^
		// |
		t.Run("from up", func(t *testing.T) {
			p := player{}
			p.addTailAt(5, 3, CELL_UP) // head
			p.addTailAt(5, 2, CELL_UP) // mid
			p.addTailAt(5, 1, CELL_UP) // tail

			p.Move(MOVE_UP)

			composeTests(
				t,
				[]composableTests{
					{got: *p.head, want: playerCell{x: 5, y: 4, direction: CELL_UP}},
					{got: *p.head.previous, want: playerCell{x: 5, y: 3, direction: CELL_UP}},
					{got: *p.head.previous.previous, want: playerCell{x: 5, y: 2, direction: CELL_UP}},
				},
			)

			if p.head.previous.previous.previous != nil {
				t.Error("cell unexisting")
			}
		})

		// error
		t.Run("from down", func(t *testing.T) {
			p := player{}
			p.addTailAt(0, 1, CELL_DOWN) // head
			p.addTailAt(0, 2, CELL_DOWN) // mid
			p.addTailAt(0, 3, CELL_DOWN) // tail

			p.Move(MOVE_UP)

			composeTests(
				t,
				[]composableTests{
					{got: *p.head, want: playerCell{x: 0, y: 1, direction: CELL_DOWN}},
					{got: *p.head.previous, want: playerCell{x: 0, y: 2, direction: CELL_DOWN}},
					{got: *p.head.previous.previous, want: playerCell{x: 0, y: 3, direction: CELL_DOWN}},
				},
			)

			if p.head.previous.previous.previous != nil {
				t.Error("cell unexisting")
			}
		})
	})

	t.Run("down", func(t *testing.T) {
		// -┐
		//  v
		t.Run("from right", func(t *testing.T) {
		})

		// ┌-
		// v
		t.Run("from left", func(t *testing.T) {
		})

		// error
		t.Run("from up", func(t *testing.T) {
		})

		// |
		// v
		t.Run("from down", func(t *testing.T) {
		})
	})

	t.Run("right", func(t *testing.T) {
		// --->
		t.Run("from right", func(t *testing.T) {
			p := NewPlayer()

			type pos struct {
				x int
				y int
			}

			current := p.head
			var initPositions []pos

			for {
				if current == nil {
					break
				}

				initPositions = append(initPositions, pos{x: current.x, y: current.y})
				current = current.previous
			}

			p.Move(MOVE_RIGHT)

			current = p.head
			i := 0

			for {
				if current == nil {
					break
				}

				expectedPositionX := initPositions[i].x + 1

				if expectedPositionX != current.x {
					t.Errorf("expected %d but got %d", expectedPositionX, current.x)
				}

				i++
				current = current.previous
			}
		})

		// error
		t.Run("from left", func(t *testing.T) {
		})

		// ┌>
		t.Run("from up", func(t *testing.T) {
		})

		// └>
		t.Run("from down", func(t *testing.T) {
		})
	})

	t.Run("left", func(t *testing.T) {
		// error
		t.Run("from right", func(t *testing.T) {
		})

		// <---
		t.Run("from left", func(t *testing.T) {
		})

		// <┘
		t.Run("from up", func(t *testing.T) {
		})

		// <┐
		t.Run("from down", func(t *testing.T) {
		})
	})

	/*t.Run("move from right to up", func(t *testing.T) {
		gotP := NewPlayer()

		wantC1 := playerCell{x: gotP.head.x, y: gotP.head.y + 1, direction: CELL_UP}

		wantC2 := playerCell{x: gotP.head.x, y: gotP.head.y, direction: CELL_CORNER_RIGHT_UP}
		wantC1.previous = &wantC2
		wantC2.next = &wantC1

		wantC3 := playerCell{x: gotP.head.x - 1, y: gotP.head.y, direction: CELL_RIGHT}
		wantC2.previous = &wantC3
		wantC3.next = &wantC2

		wantP := player{
			head: &wantC1,
			tail: &wantC3,
		}

		gotP.Move(MOVE_UP)

		current := gotP.head

		for {
			if current == nil {
				break
			}

			fmt.Println(current)
			current = current.previous
		}

		gotP.Move(MOVE_UP)

		if !reflect.DeepEqual(gotP, wantP) {
			t.Error()
		}
	})*/
}

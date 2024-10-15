package internal

type playerCellDirection int

const (
	CELL_UP playerCellDirection = iota
	CELL_DOWN
	CELL_LEFT
	CELL_RIGHT
	CELL_CORNER_RIGHT_UP
	CELL_CORNER_LEFT_UP
	CELL_CORNER_RIGHT_DOWN
	CELL_CORNER_LEFT_DOWN
)

type playerMoveDirection int

const (
	MOVE_UP playerMoveDirection = iota
	MOVE_DOWN
	MOVE_LEFT
	MOVE_RIGHT
)

type playerCell struct {
	x         int
	y         int
	direction playerCellDirection
	next      *playerCell
	previous  *playerCell
}

type player struct {
	head *playerCell
	tail *playerCell
	//s    *screen.Screen
	//sPts []screen.Point
}

/*
	┌-> -> -> -> -> ->  ┐
   	|                   |
   	^                   v
   	|                   |
   	<- <- <- <- <- <- <-┘
*/

//const SCREEN_POINT_CHAR = '@'
//const SCREEN_POINT_HEAD_CHAR = '@'

func NewPlayer() *player {
	p := player{}
	// -->
	p.addTailAt(2, 1, CELL_RIGHT) // head
	p.addTailAt(1, 1, CELL_RIGHT) // mid
	p.addTailAt(0, 1, CELL_RIGHT) // tail

	return &p
}

func (p *player) AddTail() {
	p.addTailAt(p.tail.x+1, p.tail.y, p.tail.direction)
}

func (p *player) addHeadAt(x int, y int, dir playerCellDirection) {
	nextCell := playerCell{x: x, y: y, direction: dir}

	prevHead := p.head

	if prevHead == nil {
		p.head = &nextCell
		// unsynced head: when tail -> mid -> mid but head is nil
		return
	}

	prevHead.next = &nextCell
	nextCell.previous = prevHead

	p.head = &nextCell
}

func (p *player) addTailAt(x int, y int, dir playerCellDirection) *playerCell {
	nextCell := playerCell{x: x, y: y, direction: dir}

	prevTail := p.tail

	if prevTail == nil {
		p.head = &nextCell
		p.tail = &nextCell
		return &nextCell
	}

	nextCell.next = prevTail
	prevTail.previous = &nextCell

	p.tail = &nextCell

	return &nextCell
}

func (p *player) dropTail() {
	nextAfterTail := p.tail.next

	if nextAfterTail == nil {
		p.tail = nil
		return
	}

	nextAfterTail.previous = nil
	p.tail = nextAfterTail
}

func (p *player) Move(dir playerMoveDirection) {
	if dir == MOVE_UP {
		oldHeadDir := p.head.direction

		if oldHeadDir == CELL_DOWN {
			return
		}

		p.addHeadAt(p.head.x, p.head.y+1, CELL_UP)

		switch oldHeadDir {
		case CELL_RIGHT:
			p.head.previous.direction = CELL_CORNER_RIGHT_UP
		case CELL_LEFT:
			p.head.previous.direction = CELL_CORNER_LEFT_UP
		}

		p.dropTail()
		return
	}

	next := p.head

	for {
		if next == nil {
			break
		}

		next.x += 1
		next = next.previous
	}

	//p.addHeadAt(p.head.x+1, p.head.y, CELL_RIGHT)
}

/*func (p *player) Draw() {
	for _, i := range p.sPts {
		p.s.AddPoint(i.X, i.Y, i.Char)
	}
}*/

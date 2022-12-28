package util

type Pos struct{ X, Y int }

func (p Pos) Add(m Pos) Pos {
	return Pos{p.X + m.X, p.Y + m.Y}
}

func (p *Pos) Move(m Move) {
	p.X += m.X
	p.Y += m.Y
}

type Move Pos

func (m Move) String() string {
	if m.X < 0 {
		return "<"
	} else if m.X > 0 {
		return ">"
	} else if m.Y > 0 {
		return "^"
	} else if m.Y < 0 {
		return "v"
	} else {
		return "o"
	}
}

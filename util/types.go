package util

type Pos2 struct{ X, Y int }
type Pos3 struct{ X, Y, Z int }

func (p Pos3) Add(m Pos3) Pos3 {
	return Pos3{p.X + m.X, p.Y + m.Y, p.Z + m.Z}
}
func MinPos(p1, p2 Pos3) Pos3 {
	return Pos3{Min(p1.X, p2.X), Min(p1.Y, p2.Y), Min(p1.Z, p2.Z)}
}
func MaxPos(p1, p2 Pos3) Pos3 {
	return Pos3{Max(p1.X, p2.X), Max(p1.Y, p2.Y), Max(p1.Z, p2.Z)}
}

func (p Pos2) Add(m Pos2) Pos2 {
	p.X += m.X
	p.Y += m.Y
	return p
}
func (p Pos2) Sub(m Pos2) Pos2 {
	return Pos2{p.X - m.X, p.Y - m.Y}
}

func (p *Pos2) Move(m Move) {
	p.X += m.X
	p.Y += m.Y
}
func (p1 Pos2) Comp(p2 Pos2) int {
	if p1.X < p2.X {
		return -1
	}
	if p1.X > p2.X {
		return 1
	}
	return p1.Y - p2.Y
}

type Move Pos2

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

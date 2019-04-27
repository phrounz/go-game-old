package main

type Pos struct {
	x float64
	y float64
}

func (p *Pos) AddPos(other *Pos) Pos {
	return Pos{x: p.x + other.x, y: p.y + other.y}
}

func (p *Pos) Add(x float64, y float64) Pos {
	return Pos{x: p.x + x, y: p.y + y}
}

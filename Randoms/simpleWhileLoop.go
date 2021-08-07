package main

import (
	"fmt"
)

type Point struct {
	x int64
	y int64
	z int64
}

func (p *Point) move(x, y, z int64) {
	p.x += x
	p.y += y
	p.z += z

}
func (p *Point) display() string {
	return fmt.Sprintf("{x:%d, y:%d, z:%d", p.x, p.y, p.z)

}

func main() {
	p := Point{x: 5, z: 12, y: 6}
	p.move(2, 1, 9)
	println(p.display())
}

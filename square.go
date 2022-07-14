package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (obj Square) End() Point {
	return Point{int(obj.a) + obj.start.x, obj.start.y + int(obj.a)}
}

func (obj Square) Area() uint {
	return obj.a * obj.a
}

func (obj Square) Perimeter() uint {
	return 4 * (obj.a)
}

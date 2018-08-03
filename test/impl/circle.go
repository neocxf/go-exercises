package impl

type Circle struct {
	R int
}

func (c *Circle) Area() int64 {
	//var pi float64 = math.Pi
	return int64(3 * c.R * c.R)
}

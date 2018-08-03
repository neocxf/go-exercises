package impl

type Rectangle struct {
	Length int
	Width  int
}

func (r *Rectangle) Area() int64 {
	return int64(r.Length * r.Width)
}

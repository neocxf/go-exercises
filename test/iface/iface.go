package iface

type Shaper interface {
	Area() int64
}

type SS struct {
}

type Args map[string]string

type Data map[string]string

type Fetcher interface {
	Fetch(args Args) (Data, error)
}

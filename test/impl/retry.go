package impl

import (
	"fmt"
	. "github.com/neocxf/go-exercises/test/iface"
	"time"
)

type Retrier struct {
	RetryCount   int
	WaitInterval time.Duration
	Fetcher      Fetcher
}

func (r *Retrier) Fetch(args Args) (Data, error) {
	for retry := 1; retry < r.RetryCount; retry++ {
		fmt.Printf("Retrier retries to fetch for %d\n", retry)

		if data, err := r.Fetcher.Fetch(args); err != nil {
			fmt.Printf("Retrier fetched for %d\n", retry)
			return data, nil
		} else if retry == r.RetryCount {
			fmt.Errorf("Retrier failed to fetch for %d times\n", retry)
			return Data{}, err
		}

		fmt.Printf("Retrier is waiting after error fetch for %v\n", r.WaitInterval)

		time.Sleep(r.WaitInterval)
	}

	return Data{}, nil
}

package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type M struct {
	M map[string]string
	sync.RWMutex
}

func main() {
	var g errgroup.Group
	mp := M{
		M: make(map[string]string),
	}
	g.Go(func() error {
		mp.Lock()
		defer mp.Unlock()
		mp.M["A"] = "BAD"
		return nil
	})
	g.Go(func() error {
		mp.RLock()
		defer mp.RUnlock()
		mp.M["B"] = "GOOD"
		return nil
	})
	if err := g.Wait(); err != nil {
		fmt.Println("goroutine err:", err)
	}
}

//go run -race main.go
